package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"todo/v1/routes"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

func main() {
	logger, _ := thoth.Init("log")
	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("NO .env file Found"))
		log.Fatal("No env file found")
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		logger.Log(errors.New("PORT not Found in .env"))
		log.Fatal("PORT not set in .env")
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

}
