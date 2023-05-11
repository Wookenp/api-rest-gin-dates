## Project Overview

This project is a Go/Gin API for managing dates, with endpoints for creating, retrieving, updating, and deleting dates. It provides a solid foundation for building more features and functionality upon.

## Project Details

The API is built with Go and Gin, and it allows users to perform CRUD (Create, Read, Update, and Delete) operations on dates. It provides endpoints for creating new dates, retrieving a single date or a list of dates, updating a date, and deleting a date. The project is organized in a modular and easy-to-follow structure, making it easy to add more endpoints or functionality as needed.

This API can be used as a starting point for building a more complex date management system or integrated with other applications. With its simple and well-organized codebase, developers can easily extend its functionality to meet their specific needs.

## Project structure

This is a project structure for a Go/Gin API for managing dates. Here is a simplified breakdown of the project structure:

```bash
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   ├── common
│   │   ├── config
│   │   │   └── config.go
│   │   ├── db
│   │   │   └── db.go
│   │   ├── envs
│   │   └── models
│   │       └── date.go
│   └── dates
│       ├── controller.go
│       ├── create_date.go
│       ├── delete_date.go
│       ├── list_dates.go
│       ├── retrieve_date.go
│       └── update_date.go
└── README.md
```

- `cmd`: This folder contains the main file of the application that starts the server.

- `go.mod` and `go.sum`: These files define the project's dependencies and their versions.

- `Makefile`: This file contains a set of rules to build, test, and run the application.

- `pkg`: This folder contains the reusable packages that make up the application:

- `common`: This package contains common functionality used throughout the application, such as configuration, database access, environment variables, and data models.

- `dates`: This package contains the code for managing dates, including a controller and six sub-packages for creating, retrieving, updating, and deleting dates.

Overall, the project is organized in a modular and easy-to-follow structure. The Makefile provides a convenient way to build and run the application, while the pkg folder separates the code into reusable packages.

## How to run the application

To run the application, you need to have Go installed on your machine. If you don't have it installed, you can download and install it from the official website.

After installing Go, you can clone the repository and start the application by running the following commands:

```bash
git clone https://github.com/Wookenp/date-manager-simple-api-gin.git
cd date-manager-simple-api-gin
go mod download
make run
```

The `make run` command will build the application and start the server on port 8080. You can then access the API at http://localhost:8000.

If every thing go well you will se the following output in the terminal:

```bash
Running date-manager-api...
go run cmd/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/dates/create      --> github.com/Wookenp/date-manager-api-golang/pkg/dates.CreateDate (3 handlers)
[GIN-debug] GET    /api/v1/dates             --> github.com/Wookenp/date-manager-api-golang/pkg/dates.ListDates (3 handlers)
[GIN-debug] GET    /api/v1/dates/:id         --> github.com/Wookenp/date-manager-api-golang/pkg/dates.Retrievedate (3 handlers)
[GIN-debug] PUT    /api/v1/dates/:id         --> github.com/Wookenp/date-manager-api-golang/pkg/dates.Updatedate (3 handlers)
[GIN-debug] DELETE /api/v1/dates/:id         --> github.com/Wookenp/date-manager-api-golang/pkg/dates.Deletedate (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

## Implemented Programming Fundamentals

This project implements several core Go programming concepts, such as variables, functions, modules, packages, services, interfaces, structs, pointers, slices, maps, and more.
