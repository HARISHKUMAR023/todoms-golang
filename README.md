# Todo REST API Application

This is a simple Todo REST API application built with Golang and the Gin framework. It allows users to create, read, update, and delete todo items.

## Project Structure

### Directories

- **config**: Contains configuration files.
- **controllers**: Contains the handler functions that define the logic for each API endpoint.
- **models**: Contains the data models.
- **routes**: Contains the route definitions that map HTTP requests to the corresponding controller functions.


### Files

- **main.go**: The entry point of the application.
- **Dockerfile**: Contains the instructions to build a Docker image for the application.
- **go.mod**: The Go module file.
- **go.sum**: The Go checksum file.
- **.env**: Environment variables file.
- **.gitignore**: Specifies files and directories that should be ignored by Git.
- **.air.toml**: Configuration file for the Air live reloading tool.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Docker (optional, for containerization)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/HARISHKUMAR023/todoms-golang
    cd todoms-golang
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

### Running with Docker

1. Build the Docker image:
    ```sh
    docker build -t todo-app .
    ```

2. Run the Docker container:
    ```sh
    docker run -p 8080:8080 todo-app
    ```
### Example .env File
Create a .env file in the root directory of your project with the following content:
``` env
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```
## API Endpoints

- **GET /api/v1/todos**: Retrieve all todos.
- **POST /api/v1/todos**: Create a new todo.
- **GET /api/v1/todos/:id**: Retrieve a specific todo by ID.
- **PUT /api/v1/todos/:id**: Update a specific todo by ID.
- **DELETE /api/v1/todos/:id**: Delete a specific todo by ID.

## Example Requests

### Get All Todos

```sh
curl -X GET http://localhost:8080/api/v1/todos
```
### Creating todos
``` sh
curl -X POST http://localhost:8080/api/v1/todos -H "Content-Type: application/json" -d '{"title": "New Todo", "status": "pending"}'
```
### Geting todos by id 
``` sh
curl -X GET http://localhost:8080/api/v1/todos/1
```
### Updating todo by using id 
``` sh
curl -X PUT http://localhost:8080/api/v1/todos/1 -H "Content-Type: application/json" -d '{"title": "Updated Todo", "status": "completed"}'
```
### Delete todo by id 
``` sh
curl -X DELETE http://localhost:8080/api/v1/todos/1
```
### Contributing
 Contributions are welcome! Please open an issue or submit a pull request.
### License
This project is licensed under the MIT License.

