# Go web server

Web server application by using Golang.

# Directory structure

```
.
├── server.go
├── router
|   ├── router_test.go                 //Router unit test
│   └── router.go                      //All API endpoints definition
├── model
│   └── models.go                      //payload format struct
├── middleware
|   ├── authenticationHandler.go       //AuthN validation handler for each request
|   ├── authorizationHandler.go        //AuthZ validation handler for each request
|   ├── loggerHandler.go               //Logger handler for each request
|   └── rateLimiterHandler.go          //Rate limiter handler for each request
├── testdata
│   └── newPost.yaml                   //YAML file for post new data unit test
```

# Getting started

## Install Dependencies

From the project root, run:

```
go build ./...
go test ./...
go mod tidy
```

## Quick Start

From the project root, run:

```
go run server.go
```

API endpoints:

- Get health status: http://localhost:8080/v1/health
- Get all payloads: http://localhost:8080/v1/payloads
- Get payloads with query string: http://localhost:8080/v1/payloads?license=Apache-2.0
- Post new payload : http://localhost:8080/v1/payloads

## Testing

From the project root, run:

```
go test ./... -v
```
