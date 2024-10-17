
# Authify Service

Authify is an authentication service built using Go, Fiber, and SQLite. It provides functionalities for user registration, login, and JWT-based authentication.

## Features

- **Fast and efficient** authentication service
- **Supports user registration** and **login** functionalities
- **JWT-based** secure authentication
- **SQLite database support** with both on-disk and in-memory options
- **Synchronous API** ensuring better performance and concurrency handling
- **Rate limiting** for request control
- **CORS support** for secure API access from different origins
- **Middleware handling** for error logging, request logging, and authentication checks

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
            └── 📁db_migrations
                └── setup.go
            └── queries.sql
            └── schema.sql
            └── sqlc.yaml
        └── 📁model
            └── db.go
            └── models.go
            └── queries.sql.go
    └── 📁handler
        └── login_handler.go
        └── logout_handler.go
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
    └── docker-compose.yml
    └── Dockerfile
    └── go.mod
    └── go.sum
    └── LICENSE
    └── Makefile
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

To generate the Go code from your SQL queries, use `sqlc`. Ensure that you have the correct `sqlc.yaml` configuration file.

### Run the following command to generate Go code from SQL:

```bash
make sqlc-generate
```

## .env Configuration

The `.env` file contains configuration variables that the project uses. Here are the key variables:

- `SECRET_KEY`: The secret key used for JWT encryption. You can generate a 256-bit key using OpenSSL:

  ```bash
  openssl rand -base64 32
  ```

- `DATABASE_PATH`: The path to the SQLite database file.

- `PORT`: The port on which the server runs. For Docker or public deployment, set it as `0.0.0.0:3000`.

- `LIMIT`: The maximum number of requests allowed per second (rate limiting).

### Example `.env` file:

```env
SECRET_KEY=your_generated_secret_key
DATABASE_PATH=./db/database/data.db
PORT=:3000
LIMIT=5
```

## Running the Project

### Development Mode

To run the project in development mode:

```bash
go run ./cmd
```

This will execute the `main.go` file located in the `cmd` folder.

### Build and Run

To build and run the project in production mode:

1. Build the project:

   ```bash
   go build -o auth ./cmd
   ```

2. Run the built binary:

   ```bash
   ./auth
   ```

## Docker Setup

### Build the Docker Image

```bash
make docker-build
```

### Run the Docker Container

```bash
make docker-run
```

For public deployment, ensure that the port is set to `0.0.0.0:3000` in the `.env` file.

## API Endpoints

### Registration Endpoint

- **POST** `/register`
  - Body:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```

### Login Endpoint

- **POST** `/login`
  - Body:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```

### Logout Endpoint

- **POST** `/logout`
  - Body:
    ```json
    {
      "user_id": "user_id_to_logout"
    }
    ```

### Protected Endpoint

- **GET** `/protected`
  - Headers: Must include a valid JWT token in the authorization header.

  ```bash
  Authorization: Bearer <your_jwt_token>
  ```

## Makefile Commands

1. **run**: 
   - Runs the Go project by executing the main file located in `cmd/main.go`.

2. **migrate db-up**: 
   - Runs database migrations using the setup file in `db/migrations/db_migrations/setup.go`.

3. **build**: 
   - Compiles the Go project into a binary named `authfiy`.

4. **clean**: 
   - Cleans up the project by removing the compiled binary.

5. **rebuild**: 
   - Cleans and rebuilds the project from scratch.

6. **all**: 
   - Builds and then immediately runs the project.

7. **docker-build**: 
   - Builds the Docker image for the project using the Dockerfile.

8. **docker-run**: 
   - Runs the Docker container and logs the output. The container will run on the host network using the specified port from `.env`.

9. **docker-clean**: 
   - Stops and removes the Docker container if it's running.

10. **docker-restart**: 
    - Stops, removes, and then restarts the Docker container.

11. **docker-compose-up**: 
    - Runs the project using `docker-compose`.

12. **sqlc-generate**: 
    - Generates Go code from SQL queries based on the configuration file located at `db/migrations/sqlc.yaml`.

13. **help**: 
    - Displays all available `Makefile` commands with a brief description.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
