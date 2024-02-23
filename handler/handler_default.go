package handler

import (
	"assignment-01/constants"
	"fmt"
	"net/http"
)

// ------------------------------------------------------------------------------------------------
// DefaultHandler(): Provides info about services and available endpoints
// ------------------------------------------------------------------------------------------------
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure interpretation as HTML by client
	w.Header().Set("content-type", "text/html")

	// Information about possible services and their endpoints
	output := "This service does not provide any functionality on root path level.<br><br>"

	output += "Available services:<br>"
	output += "Go to BOOKCOUNT endpoint <a href=\"" + constants.BOOKCOUNT_PATH + "\">" + constants.BOOKCOUNT_PATH + "</a><br>"
	output += "Go to READERSHIP endpoint <a href=\"" + constants.READERSHIP_PATH + "\">" + constants.READERSHIP_PATH + "</a><br>"
	output += "Go to STATUS endpoint <a href=\"" + constants.STATUS_PATH + "\">" + constants.STATUS_PATH + "</a><br>"

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		// Output error
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}
