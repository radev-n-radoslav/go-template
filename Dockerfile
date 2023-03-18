FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o bin/myapp .

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/bin/myapp .

EXPOSE 8080

ENV ENVIRONMENT "production"

ENTRYPOINT ["/myapp"]