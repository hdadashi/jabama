# JaBaMa — Hotel Reservation Web Application

A server-rendered hotel reservation web application built with Go and PostgreSQL. The project uses Go's standard HTTP server, Chi for routing, and HTML templates to render the user interface.

## Technology Stack

| Technology | Purpose |
|---|---|
| Go 1.19 | Backend application |
| PostgreSQL | Relational database |
| Chi | HTTP routing and middleware |
| SCS | Session management |
| Nosurf | CSRF protection |
| HTML templates | Server-rendered web pages |
| JavaScript, HTML and CSS | Frontend interface |

## Features and Architecture

The project is organized into separate packages to keep HTTP handling, database access, rendering, and application logic distinct.

- **HTTP routing** — Request routing using Chi.
- **Session management** — Server-side session handling using SCS.
- **Database connectivity** — PostgreSQL integration.
- **Template rendering** — Server-rendered HTML using Go templates.
- **Request handling** — Dedicated handlers for application requests.
- **Form protection** — CSRF protection dependency through Nosurf.
- **Data models** — Structures representing reservations, users, rooms, and room restrictions.

## Project Structure

```text
jabama/
├── blades/       # Application templates or view-related files
├── config/       # Application configuration
├── css/          # Stylesheets
├── driver/       # Database connectivity
├── forms/        # Form handling and validation
├── handlers/     # HTTP handlers
├── helpers/      # Shared helper functions
├── images/       # Image assets
├── js/           # JavaScript assets
├── models/       # Application data models
├── render/       # Template rendering and cache
├── repository/   # Data access layer
├── routes/       # HTTP route definitions
├── vendor/       # Vendored dependencies
├── main.go       # Application entry point
├── go.mod        # Go module and dependencies
└── go.sum        # Dependency checksums
```

## Prerequisites

Before running the application, install:

- Go 1.19 or a compatible Go version.
- PostgreSQL.
- Git.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/hdadashi/jabama.git
cd jabama
```

### 2. Configure PostgreSQL

Create a PostgreSQL database for the application:

```sql
CREATE DATABASE bookings;
```

The current application configuration expects a PostgreSQL database named `bookings` on `localhost:5432`.

**Important:** The database connection string currently appears in `main.go`. Update it to match your local PostgreSQL credentials before running the application. Avoid committing real passwords or other secrets to source control.

The database schema and any required seed data must also be initialized before the application can use the database. Consult the repository's database-related code for the required schema.

### 3. Install dependencies

Download the Go module dependencies:

```bash
go mod download
```

### 4. Run the application

Start the application from the repository root:

```bash
go run .
```

The HTTP server is configured to listen on port `8080`.

Open the application in your browser:

[http://localhost:8080](http://localhost:8080)

## Development

### Build the application

```bash
go build -o jabama .
```

Run the compiled executable:

```bash
./jabama
```

On Windows:

```powershell
.\jabama.exe
```

### Update dependencies

```bash
go mod tidy
```

Review any resulting changes to `go.mod` and `go.sum` before committing them.

## Configuration and Security

Before deploying the application:

- Move database credentials out of source code and into environment-based configuration or a secrets manager.
- Enable secure session cookies when serving over HTTPS.
- Configure the application's production mode appropriately.
- Verify CSRF protection and session settings.
- Use database migrations to manage schema changes.
- Configure appropriate logging, error handling, and production HTTP server timeouts.

The current source initializes the application with development-oriented settings. Review these settings before exposing the application publicly.

## Contributing

Contributions are welcome.

1. Fork the repository.
2. Create a branch for your changes.
3. Implement and test your changes.
4. Submit a pull request describing the changes.

## License

A license has not been verified for this repository. Check the repository's licensing terms before redistributing or using the project in other applications.

## Repository

Source code: [github.com/hdadashi/jabama](https://github.com/hdadashi/jabama)
