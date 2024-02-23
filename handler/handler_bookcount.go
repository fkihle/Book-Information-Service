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
	"sync"
)

// ------------------------------------------------------------------------------------------------
// BookcountHandler(): Allow GET requests, filter out the rest
// ------------------------------------------------------------------------------------------------
func BookcountHandler(w http.ResponseWriter, r *http.Request) {

	// System hardening; Only allow and process GET requests
	switch r.Method {
	case http.MethodGet:
		BookCountGetRequestHandler(w, r)
	default:
		// Output error message for all other request types
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

// ------------------------------------------------------------------------------------------------
// BookCountGetRequestHandler(): Main request processing function
// ------------------------------------------------------------------------------------------------
func BookCountGetRequestHandler(w http.ResponseWriter, r *http.Request) {

	// Grab the language parameter(s) from the query string
	languageQuery := r.URL.Query().Get("language")

	// Check for and remove duplicates, placing the rest in a list.
	langs := uniqueLanguages(languageQuery)

	// BooksOutput struct for collecting the relevant data for output
	var booksOutputData []structs.BooksOutput

	if len(langs) > 0 {
		// Iterate through languages at get book data
		for _, lang := range langs {

			// Get book data for current language query
			bookCount, books, err := fetchBooksData(lang)
			if err != nil {
				log.Printf("Error fetching books for language %s: %v\n", lang, err)
				return
			}

			// Calculate number of unique authors from collected data
			uniqueAuthors := countUniqueAuthors(books)

			// Append data for current query to BooksOutput struct array
			booksOutputData = append(booksOutputData, structs.BooksOutput{
				Language:    lang,
				BookCount:   bookCount,
				AuthorCount: uniqueAuthors,
				Fraction:    float64(bookCount) / float64(totalNumofBooks()),
			})
		}

		// Marshal BooksOutput data into JSON
		jsonOutputData, err := json.Marshal(booksOutputData)
		if err != nil {
			log.Println("Error generating JSON", err)
			return
		}

		// Set the content type to application/json
		w.Header().Set("Content-Type", "application/json")

		// Output the gathered data as JSON to browser
		w.Write(jsonOutputData)

	} else {
		// print instruction to client if no language input is found
		_, err := fmt.Fprint(w, "Please provide a language in the format \"/bookcount?language=en\"\n")
		if err != nil {
			http.Error(w, "Error displaying /bookcount", http.StatusInternalServerError)
			return
		}

	}
}

// ------------------------------------------------------------------------------------------------
// uniqueLanguages(): removes duplicate languages from query
// ------------------------------------------------------------------------------------------------
func uniqueLanguages(languageQuery string) []string {

	// A map to track unique languages
	uniqueLangs := make(map[string]bool)
	var result []string

	// Split query into list
	langs := strings.Split(languageQuery, ",")

	for _, lang := range langs {
		if _, exists := uniqueLangs[lang]; !exists && lang != "" {
			// Using lang as the key to avoid duplicates
			uniqueLangs[lang] = true
			result = append(result, lang)
		}
	}

	return result
}

// ------------------------------------------------------------------------------------------------
// countUniqueAuthors(): returns the count of unique authors
// ------------------------------------------------------------------------------------------------
func countUniqueAuthors(books []structs.Book) int {

	// A map to track unique authors
	authorsMap := make(map[string]bool)

	for _, book := range books {
		for _, author := range book.Authors {
			// Using author's name as the key to avoid duplicates
			authorsMap[author.Name] = true
		}
	}
	return len(authorsMap)
}

// ------------------------------------------------------------------------------------------------
// totalNumofBooks(): gets the total number of books in gutendex from the "main" page
// ------------------------------------------------------------------------------------------------
func totalNumofBooks() int {

	var totalBooks structs.BooksData

	// Make the HTTP GET request to GUTENDEX API mainpage
	resp, err := http.Get(constants.GUTENDEX_API)
	if err != nil {
		log.Println("Error getting total number of books in gutendex", err)
		return 0
	}
	defer resp.Body.Close()

	// Decode recieved JSON data
	if err := json.NewDecoder(resp.Body).Decode(&totalBooks); err != nil {
		log.Println("Error decoding JSON:", err)
		return 0
	}
	// Return the total number of books in GUTENDEX API
	return totalBooks.Count
}

// ------------------------------------------------------------------------------------------------
// fetchPage(): a goroutine for fetching a single page of books data
// ------------------------------------------------------------------------------------------------
func fetchPage(url string, ch chan<- []structs.Book, wg *sync.WaitGroup) {

	// Signal that this goroutine is done after it finishes
	defer wg.Done()

	var booksData structs.BooksData

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching page:", err)
		ch <- nil // Send nil to indicate failure
		return
	}
	defer resp.Body.Close()

	// Decode JSON data from resp.Body into booksData
	if err := json.NewDecoder(resp.Body).Decode(&booksData); err != nil {
		log.Println("Error decoding JSON:", err)
		ch <- nil // Send nil to indicate failure
		return
	}

	// Send the fetched books data back through the channel
	ch <- booksData.Results
}

// ------------------------------------------------------------------------------------------------
// fetchBooksData(): fetches all pages of book data for a given language
// ------------------------------------------------------------------------------------------------
func fetchBooksData(language string) (int, []structs.Book, error) {

	var allBooks []structs.Book
	var booksData structs.BooksData

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Create a channel to receive books data from goroutines
	ch := make(chan []structs.Book)

	url := constants.GUTENDEX_API + "?languages=" + language

	// Make an initial request to get the first page and total book count
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	// Decode JSON data from resp.Body into booksData
	if err := json.NewDecoder(resp.Body).Decode(&booksData); err != nil {
		return 0, nil, err
	}
	// Append the first page of results
	allBooks = append(allBooks, booksData.Results...)

	// Given total number of books, estimate the total number of pages
	var totalPages int
	if len(booksData.Results) > 0 {
		totalPages = (booksData.Count / len(booksData.Results)) + 1
	} else {
		totalPages = 0
	}

	// Start fetching the remaining pages concurrently
	for i := 2; i <= totalPages; i++ {
		// Increment the WaitGroup counter
		wg.Add(1)
		// Start a goroutine for each page
		go fetchPage(url+"&page="+strconv.Itoa(i), ch, &wg)
	}

	// Close the channel once all goroutines are done
	go func() {
		// Wait for all goroutines to finish
		wg.Wait()
		// Close the channel to signal that no more data will be sent
		close(ch)
	}()

	// Collect the results from the channel
	for books := range ch {
		if books != nil {
			// Append the results from each goroutine
			allBooks = append(allBooks, books...)
		}
	}

	// Return the total book count and all fetched books
	return booksData.Count, allBooks, nil
}
