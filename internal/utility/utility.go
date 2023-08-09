package utility

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to read .env. Ensure that you rename .env.example to .env")
	}
}

func getEnv(key string) string {
	loadEnv()
	value, status := os.LookupEnv(key)
	if !status {
		log.Fatalf("Missing environment variable %s\n", key)
	}
	return value
}

func Address() string {
	return getEnv("ADDRESS")
}

func Password() string {
	return getEnv("PASSWORD")
}

func Database() int {
	databaseStr := getEnv("DATABASE")
	database, err := strconv.Atoi(databaseStr)
	if err != nil {
		log.Fatalln("Database environment variable must be a number")
	}
	return database
}
