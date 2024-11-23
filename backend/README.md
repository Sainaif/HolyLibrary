# Backend

## Project Setup
Make sure you have [Go](https://golang.org/dl/) installed and properly configured. Clone the repository and navigate to the `backend` directory.

### Install Dependencies
No external dependencies need manual installation. Go modules will handle dependencies automatically.

```bash
go mod tidy
```

### Configure Environment Variables
Create a `.env` file in the `backend` directory with the following content:
```env
DB_NAME=holylibrary
DB_USER=<holylibrary_user>
DB_PASSWORD=<password>
DB_HOST=localhost
DB_PORT=5432
```

Replace these values with your actual database configuration.

### Run the Development Server
Start the server locally:
```bash
go run main.go
```
The backend will be available at `http://localhost:8080`.

### Build for Production
To compile the backend into a standalone executable:
```bash
go build -o holyLibrary-backend
```
This will generate a binary named `holyLibrary-backend` in the current directory.

### Run the Production Build
Run the compiled executable:
```bash
./holyLibrary-backend
```

### API Endpoints
Here are some of the main endpoints exposed by the backend:

| Endpoint       | Method | Description                       | Auth Required |
|----------------|--------|-----------------------------------|---------------|
| `/api/login`   | POST   | Authenticate a user and return a JWT | No            |
| `/api/register`| POST   | Register a new user               | No            |
| `/api/files`   | GET    | Retrieve a list of CBZ files      | Yes           |
| `/api/files/:id` | GET  | Fetch details for a specific CBZ file | Yes         |
| `/api/files/:id` | PUT  | Update metadata for a specific file | Yes         |

---

### Customize Configuration
See [Go Documentation](https://golang.org/doc/).