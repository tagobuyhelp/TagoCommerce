package config

import (
	"os"
)

var (
	MongoURI   string
	JWTSecret  string
)

func LoadEnv() {
	MongoURI = os.Getenv("MONGO_URI")
	JWTSecret = os.Getenv("JWT_SECRET")
}