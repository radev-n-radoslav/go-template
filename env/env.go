package env

import _ "github.com/joho/godotenv/autoload"

var ProjectName = getOptionalEnv("PROJECT_NAME", "go-api")
var Env = getOptionalEnv("ENV", "development")
var AppUrl = getOptionalEnv("APP_URL", "127.0.0.1")
var Port = getOptionalEnv("PORT", "8080")

var LogLevel = getOptionalEnv("LOG_LEVEL", "debug")

var DatabaseConnection = getOptionalEnv("DATABASE_CONNECTION", "mysql")
var DatabaseHost = getOptionalEnv("DATABASE_HOST", "localhost")
var DatabasePort = getOptionalEnv("DATABASE_PORT", "3306")
var DatabaseName = getOptionalEnv("DATABASE_NAME", "")
var DatabaseUser = getOptionalEnv("DATABASE_USER", "root")
var DatabasePassword = getOptionalEnv("DATABASE_PASSWORD", "")

var RedisHost = getOptionalEnv("REDIS_HOST", "")
var RedisPort = getOptionalEnv("REDIS_PORT", "")
var RedisPassword = getOptionalEnv("REDIS_PASSWORD", "")
