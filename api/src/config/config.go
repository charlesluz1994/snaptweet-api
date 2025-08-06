package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// DOTENV é o pacote para ler variaveis de ambiente. go get github.com/joho/godotenv
var (
	StringConnectionDB = "root:123456@tcp(127.0.0.1:3306)/devbookdb?charset=utf8&parseTime=True&loc=Local"
	Port               = 0
)

// Load the environment variables
func Load() {
	var err error

	// Read the .env file and insert it in the environment.
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT")) //convert to int
	if err != nil {
		Port = 9000 //second alternative in case of error to run in that port.
	}

	StringConnectionDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
