package structs

// ------------------------------------------------------------------------------------------------
// STRUCTS -- STATUS OUTPUT
// ------------------------------------------------------------------------------------------------

type StatusOutput struct {
	GutendexAPI  int    `json:"gutendexapi"`
	LanguageAPI  int    `json:"languageapi"`
	CountriesAPI int    `json:"countriesapi"`
	Version      string `json:"version"`
	Uptime       uint   `json:"uptime"`
}
