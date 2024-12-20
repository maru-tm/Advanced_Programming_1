# Fitness Hub

This project is a RESTful server application for managing users. It allows for the creation, retrieval, updating, and deletion of user records via HTTP requests. The server is built using Golang and PostgreSQL.

## Features

- Create a new user
- Retrieve a user by ID
- Update user details
- Delete a user by ID
- Retrieve a list of all users
- Interactive HTML page for testing and interacting with the server

## Project Structure

```
task2/
├── cmd/
│   ├── main.go           # Entry point for the application
├── handlers/
│   ├── handlers.go       # Contains HTTP handler functions
│   ├── service.go        # Business logic for handling user operations
├── models/
│   ├── user.go           # Contains user data model
├── storage/
│   ├── database.go       # Contains database initialization and utility functions
├── template/
│   ├── index.html        # HTML template for testing the API
├── utils/
│   ├── response.go       # Utility functions for HTTP responses
├── go.mod
├── go.sum
```

## Prerequisites

1. Install [Go](https://golang.org/doc/install).
2. Install [PostgreSQL](https://www.postgresql.org/download/).

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/maru-tm/Advanced_Programming_1.git
   cd Advanced_Programming_1
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up the PostgreSQL database:

   - Log in to your PostgreSQL server:
     ```sh
     psql -U postgres
     ```
   - Create a new database:
     ```sql
     CREATE DATABASE mydb;
     ```
   - (Optional) Create a user with a password for the database:
     ```sql
     CREATE USER myuser WITH PASSWORD 'mypassword';
     GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
     ```

4. Update the database connection string in `storage/database.go`:

   ```go
   dsn := "user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
   ```

## Running the Application

1. Start the server:

   ```sh
   go run cmd/main.go
   ```

2. Access the interactive HTML page at: [http://localhost:8080](http://localhost:8080)

## API Endpoints

### Create User

- **POST** `/user`
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

### Get User by ID

- **GET** `/user/{id}`
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

### Update User

- **PUT** `/user/{id}`
- **Request Body:**
  ```json
  {
    "name": "Jane Doe",
    "email": "jane.doe@example.com"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "Jane Doe",
    "email": "jane.doe@example.com"
  }
  ```

### Delete User

- **DELETE** `/user/{id}`
- **Response:** Status code `204 No Content`

### Get All Users

- **GET** `/users`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com"
    }
  ]
  ```

## Testing the API with HTML Page

- Open the provided HTML page at [http://localhost:8080](http://localhost:8080).
- Use the forms to test creating, retrieving, updating, and deleting users.
- Click "Get All Users" to view the list of users in a table.

## Important Notes

- The database schema is automatically migrated at the start of the application.
- If the `users` table already exists, it will be dropped and recreated during server startup.
- Ensure the database connection string matches your PostgreSQL configuration.


