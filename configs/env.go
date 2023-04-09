package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvPORT() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading PORT from .env file")
	}
	return os.Getenv("PORT")
}

func EnvSpreadSheetAPI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Spreadsheetapiurl from .env file")
	}
	return os.Getenv("SPREADSHEETAPI")
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Mongourl from .env file")
	}
	return os.Getenv("MONGOURI")
}
