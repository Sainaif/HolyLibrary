package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"holyLibrary-backend/database"

	"github.com/elastic/go-elasticsearch/v8/esutil"
	"go.mongodb.org/mongo-driver/bson"
)

const uploadDir = "./uploads"

// UploadCBZ handles CBZ file uploads
func UploadCBZ(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file locally
	filePath := filepath.Join(uploadDir, header.Filename)
	os.MkdirAll(uploadDir, os.ModePerm)

	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// Store metadata in MongoDB
	collection := database.MongoClient.Database("holylibrary").Collection("files")
	metadata := bson.M{
		"filename": header.Filename,
		"path":     filePath,
		"size":     header.Size,
		"uploaded": time.Now(),
	}

	_, err = collection.InsertOne(context.Background(), metadata)
	if err != nil {
		http.Error(w, "Failed to save metadata", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", header.Filename)

	indexData := map[string]interface{}{
		"filename": header.Filename,
		"tags":     []string{"cbz", "manga"},
		"uploaded": time.Now(),
	}

	res, err := database.ESClient.Index(
		"cbz-files",
		esutil.NewJSONReader(indexData),
	)
	if err != nil {
		http.Error(w, "Failed to index file", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
}
