package handler

import (
	"assignment-01/constants"
	"assignment-01/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ------------------------------------------------------------------------------------------------
// StatusHandler(): Provides API, app and uptime diagnostics
// ------------------------------------------------------------------------------------------------
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	// Split the URL path into segments
	pathSegments := strings.Split(r.URL.Path, "/")

	// Store all API URLs for GET requests
	urls := []string{constants.GUTENDEX_API, constants.LANGUAGE2COUNTRIES_API, constants.RESTCOUNTRIES_API}

	// For storing HTTP Status Codes from GET requests
	var statusCodes []int

	// Make GET requests for all APIs
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			// Print error with API url
			fmt.Printf("Error making request to %v: %v\n", url, err)
			return
		}
		defer resp.Body.Close() // Don't forget to close the response body

		// Append the HTTP Status Code from current GET request
		statusCodes = append(statusCodes, resp.StatusCode)
	}

	// Fill StatusOutput struct with data
	statusOutputData := structs.StatusOutput{
		GutendexAPI:  statusCodes[0],
		LanguageAPI:  statusCodes[1],
		CountriesAPI: statusCodes[2],
		Version:      pathSegments[2],
		Uptime:       uptimeHandler(),
	}

	// Marshal StatusOutput data into JSON
	jsonOutputData, err := json.Marshal(statusOutputData)
	if err != nil {
		fmt.Println("Error generating JSON", err)
		return
	}

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Output the gathered data as JSON to browser
	w.Write(jsonOutputData)
}

// ------------------------------------------------------------------------------------------------
// uptimeHandler(): Calculates the server/applications uptime
// ------------------------------------------------------------------------------------------------

// Global variable for storing the start time of the application from main()
var StartTime time.Time

func uptimeHandler() uint {
	// Calculate uptime
	now := time.Now()
	uptime := now.Sub(StartTime)

	// Convert uptime to seconds
	uptimeInSeconds := uptime.Seconds()

	// return the current uptime converted to int from float64
	return uint(uptimeInSeconds)
}
