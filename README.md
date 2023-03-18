# go-template

Go template is as the name suggests a template for go web applications. It models Laravel's file structure and uses the Gin framework and GORM

## Installation

To install the dependencies just run:

```BASH
go mod download
go mod vendor
go mod verify
```

## Config

Most configuration is done trough env variables. So create them on your system or copy the `.env.example` file and rename it to `.env`

## Testing

To test the application in its initial state you can run:

```BASH
go test ./tests/unit/
```

## Building

A Dockerfile is provided to easily get you going. To build the project just run:

```BASH
docker build -t <the name of your image> .
```
