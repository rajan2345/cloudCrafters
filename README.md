# cloudCrafters

A RESTful API built with Go to manage and query cloud service mappings.

## Features

-   **List Cloud Services:** Retrieve a list of all available cloud services.
-   **Filter by Provider:** Get services offered by a specific cloud provider.
-   **Detailed Service Information:** Fetch details for a specific service using its code.
-   **Service Mappings:** View mappings between different cloud services.
-   **Health Check:** A simple endpoint to verify the API's operational status.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

-   [Go](https://golang.org/doc/install) (version 1.25.4 or higher)
-   [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
-   A `.env` file (see Configuration section)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/cloudCrafters.git
    cd cloudCrafters
    ```

2.  **Create a `.env` file:**

    Create a `.env` file in the root of the project and add the following environment variables:

    ```env
    # PostgreSQL settings
    POSTGRES_DB=cloudcrafters
    POSTGRES_USER=user
    POSTGRES_PASSWORD=password
    DB_PORT=5432

    # Application settings
    DATABASE_URL="host=localhost user=user password=password dbname=cloudcrafters port=5432 sslmode=disable"
    APP_PORT=8080
    ```

### Running the Application

There are two ways to run the application:

**1. Using Docker Compose (Recommended)**

This is the easiest way to get started, as it handles the database setup for you.

1.  **Start the services:**

    ```bash
    docker-compose up -d
    ```

2.  **Run the Go application:**

    ```bash
    go run cmd/api/main.go
    ```

**2. Running Locally**

If you have a local PostgreSQL instance, you can run the application directly.

1.  **Ensure your PostgreSQL database is running and accessible.**

2.  **Update the `DATABASE_URL` in your `.env` file to point to your local database.**

3.  **Run the application:**

    ```bash
    go run cmd/api/main.go
    ```

### Seeding the Database

To populate the database with initial data, run the application with the `seed` argument:

```bash
go run cmd/api/main.go seed
```

## API Endpoints

Here are the available API endpoints:

| Method | Endpoint                          | Description                               |
| ------ | --------------------------------- | ----------------------------------------- |
| `GET`  | `/health`                         | Health check                              |
| `GET`  | `/services`                       | Get all services                          |
| `GET`  | `/services/{provider}`            | Get services by provider                  |
| `GET`  | `/services/{provider}/{code}`     | Get a specific service by provider & code |
| `GET`  | `/mapping`                        | Get service mappings                      |

## Project Structure

The project follows a standard Go application structure:

```
.
├── cmd/api/main.go         # Application entry point
├── internal/               # Internal application logic
│   ├── config/             # Configuration loading
│   ├── db/                 # Database connection and migration
│   ├── mappings/           # Mappings-related logic
│   ├── router/             # API router and routes
│   ├── seed/               # Database seeding
│   └── services/           # Services-related logic
├── pkg/                    # Shared libraries
├── docker-compose.yml      # Docker Compose configuration
└── go.mod                  # Go module definition
```
