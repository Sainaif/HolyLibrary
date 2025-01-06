# Backend

## Project Setup
Make sure you have [Go](https://golang.org/dl/) installed and properly configured. Clone the repository and navigate to the `backend` directory.

### Install Dependencies
No external dependencies need manual installation. Go modules will handle dependencies automatically.

```bash
go mod tidy
```

## Setting Up Databases

### 1. PostgreSQL

#### Install PostgreSQL
If PostgreSQL is not installed, download it using winget:
```bash
winget install PostgreSQL
```

#### Access the PostgreSQL Shell
Log in to the PostgreSQL shell:
```bash
psql -U postgres
```

#### Create the Database
Run the following commands to set up the database:
```sql
-- Create a new database
CREATE DATABASE holylibrary;

-- Create a user with a password
CREATE USER <holylibrary_user> WITH PASSWORD '<password>';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE holylibrary TO <holylibrary_user>;
```

#### Verify the Setup
Exit the PostgreSQL shell:
```bash
\q
```

Test the connection to the database using the `.env` file values:
```bash
psql -U <holylibrary_user> -d holylibrary
```

If successful, the database is ready.

### 2. MongoDB

#### Install MongoDB
Install MongoDB using winget:
```bash
winget install MongoDB.Server
```

#### Start MongoDB Service
Start the MongoDB service:
```bash
sudo systemctl start mongod
```

#### Verify MongoDB is Running
Check the status:
```bash
sudo systemctl status mongod
```

#### Create Database and User
Access the MongoDB shell:
```bash
mongosh
```

Create a database and user:
```javascript
use holylibrary

db.createUser({
  user: '<mongo_user>',
  pwd: '<password>',
  roles: [{ role: 'readWrite', db: 'holylibrary' }]
})
```

Exit the shell:
```bash
exit
```

### 3. Elasticsearch

#### Install Elasticsearch
Install Elasticsearch using winget:
```bash
winget install Elasticsearch.Elasticsearch
```

#### Start Elasticsearch
Start the service:
```bash
sudo systemctl start elasticsearch
```

Verify it's running:
```bash
curl -X GET "localhost:9200"
```

#### Create an Index
Create an index for storing data:
```bash
curl -X PUT "localhost:9200/holylibrary" -H 'Content-Type: application/json' -d'{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  }
}'
```

Verify index creation:
```bash
curl -X GET "localhost:9200/holylibrary"
```

## Configure Environment Variables

Create a `.env` file in the `backend` directory with the following content:
```env
# PostgreSQL
DB_NAME=holylibrary
DB_USER=<holylibrary_user>
DB_PASSWORD=<password>
DB_HOST=localhost
DB_PORT=5432

# MongoDB
MONGO_URI=mongodb://<mongo_user>:<password>@localhost:27017/holylibrary

# Elasticsearch
ES_HOST=http://localhost:9200
ES_INDEX=holylibrary
```
Replace these values with your actual database configuration.

## Run the Development Server

Start the server locally:
```bash
go run main.go
```
The backend will be available at `http://localhost:8080`.

## Build for Production
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

| Endpoint       | Method | Description                           | Auth Required |
|----------------|--------|---------------------------------------|---------------|
| `/api/login`   | POST   | Authenticate a user and return a JWT  | No            |
| `/api/register`| POST   | Register a new user                   | No            |
| `/api/files`   | GET    | Retrieve a list of CBZ files          | Yes           |
| `/api/files/:id` | GET  | Fetch details for a specific CBZ file | Yes           |
| `/api/files/:id` | PUT  | Update metadata for a specific file   | Yes           |

---

### Customize Configuration
See [Go Documentation](https://golang.org/doc/).

# Backend

## Project Setup
Make sure you have [Go](https://golang.org/dl/) installed and properly configured. Clone the repository and navigate to the `backend` directory.

### Install Dependencies
No external dependencies need manual installation. Go modules will handle dependencies automatically.

```bash
go mod tidy
```

## Setting Up Databases

### 1. PostgreSQL

#### Install PostgreSQL
If PostgreSQL is not installed, download it using winget:
```bash
winget install PostgreSQL
```

#### Access the PostgreSQL Shell
Log in to the PostgreSQL shell:
```bash
psql -U postgres
```

#### Create the Database
Run the following commands to set up the database:
```sql
-- Create a new database
CREATE DATABASE holylibrary;

-- Create a user with a password
CREATE USER <holylibrary_user> WITH PASSWORD '<password>';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE holylibrary TO <holylibrary_user>;
```

#### Verify the Setup
Exit the PostgreSQL shell:
```bash
\q
```

Test the connection to the database using the `.env` file values:
```bash
psql -U <holylibrary_user> -d holylibrary
```

If successful, the database is ready.

### 2. MongoDB

#### Install MongoDB
Install MongoDB using winget:
```bash
winget install MongoDB.Server
```

#### Start MongoDB Service
Start the MongoDB service:
```bash
sudo systemctl start mongod
```

#### Verify MongoDB is Running
Check the status:
```bash
sudo systemctl status mongod
```

#### Create Database and User
Access the MongoDB shell:
```bash
mongosh
```

Create a database and user:
```javascript
use holylibrary

db.createUser({
  user: '<mongo_user>',
  pwd: '<password>',
  roles: [{ role: 'readWrite', db: 'holylibrary' }]
})
```

Exit the shell:
```bash
exit
```

### 3. Elasticsearch

#### Install Elasticsearch
Install Elasticsearch using winget:
```bash
winget install Elasticsearch.Elasticsearch
```

#### Start Elasticsearch
Start the service:
```bash
sudo systemctl start elasticsearch
```

Verify it's running:
```bash
curl -X GET "localhost:9200"
```

#### Create an Index
Create an index for storing data:
```bash
curl -X PUT "localhost:9200/holylibrary" -H 'Content-Type: application/json' -d'{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  }
}'
```

Verify index creation:
```bash
curl -X GET "localhost:9200/holylibrary"
```

## Configure Environment Variables

Create a `.env` file in the `backend` directory with the following content:
```env
# PostgreSQL
DB_NAME=holylibrary
DB_USER=<holylibrary_user>
DB_PASSWORD=<password>
DB_HOST=localhost
DB_PORT=5432

# MongoDB
MONGO_URI=mongodb://<mongo_user>:<password>@localhost:27017/holylibrary

# Elasticsearch
ES_HOST=http://localhost:9200
ES_INDEX=holylibrary
```
Replace these values with your actual database configuration.

## Run the Development Server

Start the server locally:
```bash
go run main.go
```
The backend will be available at `http://localhost:8080`.

## Build for Production
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

| Endpoint       | Method | Description                           | Auth Required |
|----------------|--------|---------------------------------------|---------------|
| `/api/login`   | POST   | Authenticate a user and return a JWT  | No            |
| `/api/register`| POST   | Register a new user                   | No            |
| `/api/files`   | GET    | Retrieve a list of CBZ files          | Yes           |
| `/api/files/:id` | GET  | Fetch details for a specific CBZ file | Yes           |
| `/api/files/:id` | PUT  | Update metadata for a specific file   | Yes           |

---

### Customize Configuration
See [Go Documentation](https://golang.org/doc/).