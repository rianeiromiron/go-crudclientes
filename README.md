# CRUD Clientes

A simple web application to manage clients (customers), built with Go's standard library and MySQL. It provides full CRUD (Create, Read, Update, Delete) operations through server-rendered HTML pages.

## Features

- List all clients
- Create a new client
- Edit an existing client
- Delete a client
- Auto-creates the required MySQL table on startup if it doesn't exist

## Tech Stack

- [Go](https://go.dev/) 1.24 (standard `net/http` router, no external web framework)
- MySQL (via [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql))
- Server-side rendering with `html/template`
- Plain CSS for styling

## Project Structure

```
.
├── main.go                       # Application entry point and route definitions
├── db/
│   └── db.go                     # MySQL connection and schema initialization
├── models/
│   └── cliente.go                # Cliente entity and repository (data access layer)
├── handlers/
│   └── cliente_handlers.go       # HTTP handlers for the clients CRUD
├── templates/
│   ├── clientes_list.html        # Client listing page
│   └── cliente_form.html         # Create/edit form
└── static/
    └── style.css                 # Stylesheet
```

## Prerequisites

- Go 1.24 or later
- A running MySQL server

## Configuration

The database connection is configured through environment variables (defaults shown):

| Variable      | Default         | Description          |
|---------------|-----------------|-----------------------|
| `DB_USER`     | `contabilidad`  | MySQL username        |
| `DB_PASSWORD` | `contadocker`   | MySQL password        |
| `DB_HOST`     | `127.0.0.1`     | MySQL host            |
| `DB_PORT`     | `3306`          | MySQL port             |
| `DB_NAME`     | `crudclientes`  | MySQL database name   |

On startup, the app connects to MySQL and automatically creates the `clientes` table if it does not already exist.

## Running the Application

1. Make sure MySQL is running and reachable with the configured credentials.
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
4. Open your browser at [http://localhost:8085](http://localhost:8085).

## Routes

| Method | Path                     | Description                  |
|--------|--------------------------|-------------------------------|
| GET    | `/clientes`              | List all clients              |
| GET    | `/clientes/nuevo`        | Show the "new client" form    |
| POST   | `/clientes/nuevo`        | Create a new client           |
| GET    | `/clientes/{id}/editar`  | Show the "edit client" form   |
| POST   | `/clientes/{id}/editar`  | Update an existing client     |
| POST   | `/clientes/{id}/eliminar`| Delete a client                |
