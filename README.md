
# Go Gin REST API with MySQL

This project is a basic RESTful API built using the **Gin** web framework and **MySQL** for the database. The application demonstrates how to structure a Go project, set up a MySQL database connection, and perform basic CRUD operations (Create, Read, Update, Delete).

## Features

- RESTful API with **Gin**
- **MySQL** integration for data persistence
- Environment variable management with `.env`
- CRUD operations for managing users
- Middleware support for logging and CORS
- Follows best practices for folder structure and modular code
- Example routes for creating and fetching users

## Technologies

- **Go**: A statically typed, compiled programming language.
- **Gin**: A fast HTTP web framework for Go.
- **MySQL**: Relational database management system.
- **Golang MySQL driver**: `go-sql-driver/mysql` for connecting Go with MySQL.

## Project Structure

```
├── config/             # Database configuration
│   └── database.go     # MySQL connection setup
├── controllers/        # Route handler functions
│   └── user.go         # User-related handlers
├── models/             # Database models
│   └── user.go         # User model definition
├── routes/             # API route definitions
│   └── routes.go       # Route group setup
├── services/           # Business logic, database interactions
│   └── user.go         # User service for CRUD operations
├── .env                # Environment variables file
├── .gitignore          # Git ignore rules
├── go.mod              # Go module file
├── main.go             # Application entry point
└── README.md           # Project documentation
```

## Getting Started

### Prerequisites

Ensure you have the following installed:

- **Go** (v1.16 or later)
- **MySQL** (v5.7 or later)
- **Git**

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/go-gin-mysql-api.git
   cd go-gin-mysql-api
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Create the `.env` file**:

   Create a `.env` file in the root of the project to configure your MySQL connection.

   ```bash
   DB_USER=root
   DB_PASSWORD=your_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=my_database
   ```

4. **Set up the MySQL database**:

   - Create a new MySQL database:

     ```sql
     CREATE DATABASE my_database;
     ```

   - Run migrations (manually or using a migration tool) to set up your tables. Example SQL migration:

     ```sql
     CREATE TABLE users (
       id INT AUTO_INCREMENT PRIMARY KEY,
       username VARCHAR(50) NOT NULL,
       email VARCHAR(100) NOT NULL
     );
     ```

5. **Run the application**:

   ```bash
   go run main.go
   ```

   The application will be running on `http://localhost:8080`.

### API Endpoints

Here are the available API endpoints:

| Method | Endpoint          | Description                |
|--------|-------------------|----------------------------|
| GET    | `/api/v1/users`    | Get all users              |
| POST   | `/api/v1/users`    | Create a new user          |

### Example API Usage

- **Create a new user**:

  POST request to `http://localhost:8080/api/v1/users` with JSON payload:

  ```json
  {
    "username": "john_doe",
    "email": "john.doe@example.com"
  }
  ```

- **Get all users**:

  GET request to `http://localhost:8080/api/v1/users`

### Environment Variables

The following environment variables need to be set:

- `DB_USER`: Your MySQL database username
- `DB_PASSWORD`: Your MySQL database password
- `DB_HOST`: MySQL host (usually `localhost`)
- `DB_PORT`: MySQL port (default is `3306`)
- `DB_NAME`: MySQL database name

### Folder Structure Breakdown

- **`config/database.go`**: Contains the logic to connect to MySQL using the `go-sql-driver/mysql` package.
- **`models/`**: Defines the data models used in the application.
- **`services/`**: Handles the database operations and business logic.
- **`controllers/`**: Contains route handlers that handle incoming API requests.
- **`routes/`**: Groups and registers API routes with Gin.

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a pull request

## License

This project is licensed under the MIT License.

## Acknowledgements

- [Gin Web Framework](https://gin-gonic.com/)
- [MySQL Driver for Go](https://github.com/go-sql-driver/mysql)
