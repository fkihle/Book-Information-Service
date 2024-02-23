package structs

// ------------------------------------------------------------------------------------------------
// STRUCTS -- GUTENDEX
// ------------------------------------------------------------------------------------------------

type BooksData struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"` // using *string to handle null values
	Results  []Book  `json:"results"`
}

type Book struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
	// Translators   []string `json:"translators"`
	Subjects      []string `json:"subjects"`
	Bookshelves   []string `json:"bookshelves"`
	Languages     []string `json:"languages"`
	Copyright     bool     `json:"copyright"`
	MediaType     string   `json:"media_type"`
	Formats       Formats  `json:"formats"`
	DownloadCount int      `json:"download_count"`
}

type Author struct {
	Name      string `json:"name"`
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
}

type Formats map[string]string

// Output from /books
type BooksOutput struct {
	Language    string  `json:"language"`
	BookCount   int     `json:"books"`
	AuthorCount int     `json:"authors"`
	Fraction    float64 `json:"fraction"`
}
