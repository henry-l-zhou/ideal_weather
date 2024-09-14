package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/henry-l-zhou/ideal_weather/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := server.NewServer()

	fmt.Printf("running server on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server))

}
