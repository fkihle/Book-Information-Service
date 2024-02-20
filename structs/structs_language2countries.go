package structs

// ------------------------------------------------------------------------------------------------
// STRUCTS -- LANGUAGE2COUNTRIES
// ------------------------------------------------------------------------------------------------

type Lang2CountryData struct {
	ISO3166_1_Alpha_3 string `json:"ISO3166_1_Alpha_3"`
	ISO3166_1_Alpha_2 string `json:"ISO3166_1_Alpha_2"`
	Official_Name     string `json:"Official_Name"`
	Region_Name       string `json:"Region_Name"`
	Sub_Region_Name   string `json:"Sub_Region_Name"`
	Language          string `json:"Language"`
}
