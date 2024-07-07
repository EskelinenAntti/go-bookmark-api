# Go API with Gin and GORM

A simple Go backend API with Gin framework and GORM orm.

The application uses sqlite database, which gets created on startup.

## Try it out

### Install Go

https://go.dev/doc/install

### Start the backend

`go run .`

This will install dependencies listed in `go.mod` file, compile the project, and start the backend.

### Call the endpoint

`GET`

```bash
curl http://localhost:8080/currentPage
```

`POST`

```bash
curl -d '{"page":8}' http://localhost:8080/currentPage
```