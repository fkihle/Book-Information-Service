package handler

import (
	"assignment-01/constants"
	"assignment-01/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ------------------------------------------------------------------------------------------------
// ReadershipHandler(): Allow GET requests, filter out the rest
// ------------------------------------------------------------------------------------------------
func ReadershipHandler(w http.ResponseWriter, r *http.Request) {

	// System hardening; Only allow and process GET requests
	switch r.Method {
	case http.MethodGet:
		ReadershipGetRequestHandler(w, r)
	default:
		// Output error message for all other request types
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// ------------------------------------------------------------------------------------------------
// ReadershipGetRequestHandler(): Main request processing function
// ------------------------------------------------------------------------------------------------
func ReadershipGetRequestHandler(w http.ResponseWriter, r *http.Request) {

	// Split the URL path into segments
	pathSegments := strings.Split(r.URL.Path, "/")

	if pathSegments[4] != "" {
		// Get countries that speak the queried language
		countryCodes, countryNames, err := language2country(pathSegments[4])
		if err != nil {
			log.Println("Error in language2country(): ", err)
			http.Error(w, "Unable to get information about: "+pathSegments[[4]], http.StatusInternalServerError)
			return
		}

		// Grab the limit parameter from the query string
		limitStr, err := r.URL.Query().Get("limit")
		if err != nil {
			log.Println("Unable to parse limit query: ", err)
			http.Error(w, "\nPlease provide your desired limit in the format \"/readership/no?limit=2\"\n", http.StatusBadRequest)
		}

		// Convert "limitStr" to an integer if present
		var limit int
		if limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				// Handle the error if conversion fails
				log.Println("Error converting limit input to integer: ", err)
				http.Error(w, "Invalid limit value input.\nPlease provide your limit query in the format \"/readership/no?limit=2\"\n", http.StatusBadRequest)
				return
			}
			// Check for negative values
			if limit < 0 {
				// Default to showing all
				limit = len(countryCodes)
			}
		} else {
			// If limit not present, set limit to total number of countries
			limit = len(countryCodes)
		}

		var readershipOutputData []structs.CountryOutput

		// Get data for the "limit" first countries in array
		for i, country := range countryCodes[:limit] {

			// Get population count
			countryPopulation, err := fetchCountryPopulation(country)
			if err != nil {
				log.Println("Error in fetchCountryPopulation(): ", err)
				http.Error(w, "Unable to fetch population for: "+country, http.StatusInternalServerError)
				return
			}

			// Get Book data for countUniqueAuthors() and get total number of books
			bookNum, books, err := fetchBooksData(country)
			if err != nil {
				log.Println("Error in fetchBooksData(): ", err)
				http.Error(w, "Unable to fetch books data for: "+country, http.StatusInternalServerError)
				return
			}

			// Append data for current query to CountryOutput struct array
			readershipOutputData = append(readershipOutputData, structs.CountryOutput{
				CountryName: countryNames[i],
				CCA2:        country,
				Books:       bookNum,
				Authors:     countUniqueAuthors(books),
				Popluation:  countryPopulation,
			})
		}

		// Marshal CountryOutput data into JSON
		jsonOutputData, err := json.Marshal(readershipOutputData)
		if err != nil {
			log.Println("Error generating JSON", err)
			return
		}

		// Set the content type to application/json
		w.Header().Set("Content-Type", "application/json")

		// Output the gathered data as JSON to browser
		w.Write(jsonOutputData)

	} else {
		// Print instruction to client if language segment not present
		_, err := fmt.Fprint(w, "Please provide a language in the format \"/readership/no\" or \"/readership/no?limit=2\"\n")
		if err != nil {
			http.Error(w, "Error displaying instructions.", http.StatusInternalServerError)
			return
		}
	}
}

// ------------------------------------------------------------------------------------------------
// language2country(): return country codes and names that speak a given language
// ------------------------------------------------------------------------------------------------
func language2country(twoLetterCode string) ([]string, []string, error) {

	var lang2countryData []structs.Lang2CountryData
	var countryCodes []string
	var countryNames []string

	url := constants.LANGUAGE2COUNTRIES_API + twoLetterCode

	// Get data for current country from LANGUAGE2COUNTRIES API
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error getting data from LANGUAGE2COUNTRIES API for: "+twoLetterCode, err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Decode recieved JSON data
	if err := json.NewDecoder(resp.Body).Decode(&lang2countryData); err != nil {
		log.Println("Error decoding JSON: ",err)
		return nil, nil, err
	}

	// Extract the 2-letter CCA2/ISO3166_1_Alpha_2 code and country's Official Name
	for _, lang := range lang2countryData {
		countryCodes = append(countryCodes, lang.ISO3166_1_Alpha_2)
		countryNames = append(countryNames, lang.Official_Name)
	}

	return countryCodes, countryNames, err
}

// ------------------------------------------------------------------------------------------------
// fetchCountryPopulation(): returns population count for given country
// ------------------------------------------------------------------------------------------------
func fetchCountryPopulation(country string) (int, error) {

	var singleCountry []structs.Country
	var popu int

	url := constants.RESTCOUNTRIES_API + "alpha/" + country

	// Get data for current country from RESTCOUNTRIES API
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching RESTCOUNTRIES_API", err)
		return 0, err
	}
	defer resp.Body.Close()

	// Decode recieved JSON data
	if err := json.NewDecoder(resp.Body).Decode(&singleCountry); err != nil {
		log.Println("Error decoding JSON from RESTCOUNTRIES_API", err)
	}

	// Extract the 2-letter CCA2 code and population count
	for _, country := range singleCountry {
		popu = country.Population
	}

	return popu, nil
}
