package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var MongoURI string

func init() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    MongoURI = os.Getenv("MONGO_URI")
    if MongoURI == "" {
        log.Fatal("MONGO_URI is not set")
    }
}
