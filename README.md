# Go CRUD API with Gin & PostgreSQL

This project is a simple CRUD API built using **Golang**, **Gin framework**, and **GORM** with a **local PostgreSQL** database.

## Features

- Create, Read, Update, Delete blog posts
- Environment-based config (.env)
- GORM auto-migration
- RESTful routes

## Setup

1. Clone the repository

```bash
git clone https://github.com/YourUsername/go-crud-api.git
cd go-crud-api
```

2. Create `.env` file

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=go_crud
DB_PORT=5432
SSL_MODE=disable
```

3. Run PostgreSQL and create the `go_crud` database.

4. Install dependencies

```bash
go mod tidy
```

5. Run the app

```bash
go run main.go
```

Or using **CompileDaemon** (auto-restart on file changes):

```bash
CompileDaemon -build="go build -o app" -command="./app" -include="\.go$" -polling=true
```

## API Endpoints

| Method | Endpoint      | Description       |
|--------|---------------|-------------------|
| GET    | /posts        | List all posts    |
| GET    | /posts/:id    | Get post by ID    |
| POST   | /posts        | Create new post   |
| PUT    | /posts/:id    | Update a post     |
| DELETE | /posts/:id    | Delete a post     |

## License

MIT
