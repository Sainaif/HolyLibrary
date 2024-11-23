# Backend

## Project Setup
Make sure you have [Go](https://golang.org/dl/) installed and properly configured. Clone the repository and navigate to the `backend` directory.

### Install Dependencies
No external dependencies need manual installation. Go modules will handle dependencies automatically.

```bash
go mod tidy
```
## Setting Up the Database

### 1. Install PostgreSQL
If PostgreSQL is not installed, download it from [https://www.postgresql.org/download/](https://www.postgresql.org/download/) and install it.

### 2. Access the PostgreSQL Shell
Log in to the PostgreSQL shell:
```bash
psql -U postgres
```

### 3. Create the Database
Run the following commands to set up the database:
```sql
-- Create a new database
CREATE DATABASE holylibrary;

-- Create a user with a password
CREATE USER <holylibrary_user> WITH PASSWORD '<password>';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE holylibrary TO <holylibrary_user>;
```

### 4. Verify the Setup
Exit the PostgreSQL shell:
```bash
\q
```

Test the connection to the database using the `.env` file values:
```bash
psql -U <holylibrary_user> -d holylibrary
```

If successful, the database is ready.

---
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