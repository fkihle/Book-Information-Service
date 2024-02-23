package structs

// ------------------------------------------------------------------------------------------------
// STRUCTS -- RESTCOUNTRIES
// ------------------------------------------------------------------------------------------------

type CountryOutput struct {
	CountryName string `json:"country"`
	CCA2        string `json:"isocode"`
	Books       int    `json:"books"`
	Authors     int    `json:"authors"`
	Popluation  int    `json:"readership"`
}

type Country struct {
	Name         Name         `json:"name"`
	TLD          []string     `json:"tld"`
	CCA2         string       `json:"cca2"`
	CCN3         string       `json:"ccn3"`
	CCA3         string       `json:"cca3"`
	CIOC         string       `json:"cioc"`
	Independent  bool         `json:"independent"`
	Status       string       `json:"status"`
	UNMember     bool         `json:"unMember"`
	Currencies   Currencies   `json:"currencies"`
	IDD          IDD          `json:"idd"`
	Capital      []string     `json:"capital"`
	AltSpellings []string     `json:"altSpellings"`
	Region       string       `json:"region"`
	Subregion    string       `json:"subregion"`
	Languages    Languages    `json:"languages"`
	Translations Translations `json:"translations"`
	Latlng       []float64    `json:"latlng"`
	Landlocked   bool         `json:"landlocked"`
	Borders      []string     `json:"borders"`
	Area         float64      `json:"area"`
	Demonyms     Demonyms     `json:"demonyms"`
	Flag         string       `json:"flag"`
	Maps         Maps         `json:"maps"`
	Population   int          `json:"population"`
	FIFA         string       `json:"fifa"`
	Car          Car          `json:"car"`
	Timezones    []string     `json:"timezones"`
	Continents   []string     `json:"continents"`
	Flags        Flags        `json:"flags"`
	CoatOfArms   CoatOfArms   `json:"coatOfArms"`
	StartOfWeek  string       `json:"startOfWeek"`
	CapitalInfo  CapitalInfo  `json:"capitalInfo"`
	PostalCode   PostalCode   `json:"postalCode"`
}

type Name struct {
	Common     string                `json:"common"`
	Official   string                `json:"official"`
	NativeName map[string]NativeName `json:"nativeName"`
}

type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Currencies struct {
	EUR Currency `json:"EUR"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type IDD struct {
	Root     string   `json:"root"`
	Suffixes []string `json:"suffixes"`
}

type Languages struct {
	Cat string `json:"cat"`
}

type Translations map[string]Translation

type Translation struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Demonyms struct {
	Eng Gender `json:"eng"`
	Fra Gender `json:"fra"`
}

type Gender struct {
	F string `json:"f"`
	M string `json:"m"`
}

type Maps struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}

type Car struct {
	Signs []string `json:"signs"`
	Side  string   `json:"side"`
}

type Flags struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
	Alt string `json:"alt"`
}

type CoatOfArms struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
}

type CapitalInfo struct {
	Latlng []float64 `json:"latlng"`
}

type PostalCode struct {
	Format string `json:"format"`
	Regex  string `json:"regex"`
}
