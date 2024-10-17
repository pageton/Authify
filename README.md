
# Authify Service

This is an authentication service built using Go, Fiber, and SQLite. It provides functionalities for user registration and login, along with JWT-based authentication.

## Project Structure

```
└── 📁authify
    └── 📁cmd
        └── main.go
    └── 📁config
        └── config.go
    └── 📁db
        └── 📁database
            └── data.db
        └── 📁migrations
            └── queries.sql
            └── schema.sql
            └── sqlc.yaml
        └── 📁model
            └── db.go
            └── models.go
            └── queries.sql.go
    └── 📁handler
        └── login_handler.go
        └── register_handler.go
    └── 📁middleware
        └── auth_middleware.go
        └── cors_middleware.go
        └── error_handling.go
        └── rate_limiting.go
        └── request_logging.go
    └── 📁models
        └── user.go
    └── 📁services
        └── jwt_service.go
    └── .env
    └── .env.example
    └── .gitignore
    └── go.mod
    └── go.sum
    └── README.md
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/pageton/authify.git
   cd authify
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

## Database Setup

To apply the schema for the SQLite database, run the following command:

```bash
sqlite3 ./db/database/data.db < ./db/migrations/schema.sql
```

This will create the necessary tables and indexes for your application.

## Using sqlc

To generate the Go code from your SQL queries, use `sqlc`. First, ensure that you have the correct `sqlc.yaml` configuration file.

### Run the following command to generate Go code from SQL:

```bash
sqlc generate
```

This will generate the Go code based on your schema and queries defined in `queries.sql`.

## Dependencies

This project uses the following dependencies:

- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver for Go.
- [gofiber](https://github.com/gofiber/fiber) - A fast web framework for Go.
- [sqlc](https://sqlc.dev/) - A code generation tool for SQL queries.

## Project Version

This project is currently at version **v0.0.1**.

## JWT Authentication

The service uses JWT (JSON Web Tokens) for user authentication. Ensure you validate the JWT type during authentication.

## Running the Project

### Development Mode

To run the project in development mode, use the following command:

```bash
go run ./cmd
```

This will execute the `main.go` file located in the `cmd` folder.

### Build and Run

To build and run the project in production mode, use the following commands:

1. Build the project:

   ```bash
   go build -o auth ./cmd
   ```

2. Run the built binary:

   ```bash
   ./auth
   ```

## API Endpoints

### Start Point

- **POST** `/register`
  - Body: 
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```

### End Point

- **POST** `/login`
  - Body: 
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```


## Example Request

To register a new user, send a POST request to the `/register` endpoint with the following JSON body:

```json
{
  "username": "john_doe",
  "password": "securepassword123"
}
```

To log in, send a POST request to the `/login` endpoint with the following JSON body:

```json
{
  "username": "john_doe",
  "password": "securepassword123"
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
