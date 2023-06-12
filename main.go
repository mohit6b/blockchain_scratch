package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mohit6b/pow-blockchain-go/src/router"
)

func main() {
	godotenv.Load(".env")
	nodeId := os.Getenv("NODE_ID")
	router := router.SetupRouter()

	// Start the server
	log.Println("Server started on http://localhost:" + nodeId)
	// Start the server
	log.Fatal(http.ListenAndServe(":"+nodeId, router))
}
