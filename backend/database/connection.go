package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *sql.DB
var MongoClient *mongo.Client
var ESClient *elasticsearch.Client

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("PostgreSQL connection error:", err)
	}

	log.Println("PostgreSQL connection established")
}

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@mongo:27017/",
			os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD")))

	var err error
	MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err = MongoClient.Ping(context.Background(), nil); err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	log.Println("MongoDB connection established")
}

func ConnectElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_URL"),
		},
	}

	var err error
	ESClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("Failed to connect to Elasticsearch:", err)
	}

	log.Println("Elasticsearch connection established")
}
