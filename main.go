package main

import (
	"assignment-01/constants"
	"assignment-01/handler"
	"log"
	"net/http"
	"os"
	"time"
)

// ------------------------------------------------------------------------------------------------
// MAIN
// ------------------------------------------------------------------------------------------------
func main() {

	// Get the PORT from environment variables
	port := os.Getenv("PORT")

	// Set port to a default ("8080") if port not found automatically
	if port == "" {
		log.Println("$PORT not found. Set to default: 8080")
		port = "8080"
	}

	// Endpoints and their handler functions
	http.HandleFunc(constants.DEFAULT_PATH, handler.DefaultHandler)
	http.HandleFunc(constants.BOOKCOUNT_PATH, handler.BookcountHandler)
	http.HandleFunc(constants.READERSHIP_PATH, handler.ReadershipHandler)
	http.HandleFunc(constants.STATUS_PATH, handler.StatusHandler)

	// Record the start time for use in StatusHandler
	handler.StartTime = time.Now()

	// Start server and listen on port
	log.Println("Service listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
