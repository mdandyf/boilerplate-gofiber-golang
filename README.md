# go-boilerplate

This is a boilerplate project for a Go application using the Gofiber framework. It follows the standard layout and includes additional packages for logging, Redis integration, and JWT authentication.

## Project Structure

```
go-boilerplate
├── cmd
│   └── app
│       └── main.go
├── internal
│   └── app
│       ├── app.go
│       └── app_test.go
├── pkg
│   ├── utils
│   │   └── utils.go
│   ├── logger
│   │   └── logger.go
│   ├── redis
│   │   └── redis.go
│   └── jwt
│       └── jwt.go
├── test
│   └── app_test.go
├── deployment
│   ├── Dockerfile
│   ├── deployment.yaml
│   └── service.yaml
├── go.mod
├── go.sum
└── README.md
```

## Project Description

- `cmd/app/main.go`: Entry point of the application. Initializes and starts the application.
- `internal/app/app.go`: Contains the main logic of the application. Handles routes and business logic.
- `internal/app/app_test.go`: Contains tests for the application logic.
- `pkg/utils/utils.go`: Contains utility functions used across the application.
- `pkg/logger/logger.go`: Implements a logger package for logging application events.
- `pkg/redis/redis.go`: Implements a Redis package for interacting with a Redis database.
- `pkg/jwt/jwt.go`: Implements a JWT package for handling JWT authentication and authorization.
- `test/app_test.go`: Contains tests for the application.
- `deployment/Dockerfile`: Dockerfile for building the application into a Docker image.
- `deployment/deployment.yaml`: Kubernetes deployment configuration for deploying the application.
- `deployment/service.yaml`: Kubernetes service configuration for exposing the application.
- `go.mod`: Go module file that defines the project's dependencies.
- `go.sum`: Contains cryptographic checksums of specific module versions.
- `README.md`: Documentation for the project.

Feel free to modify and extend this boilerplate project according to your specific needs. Happy coding!