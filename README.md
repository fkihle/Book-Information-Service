# **Book Information Service**

The Book Information Service is a REST web application in Go that provides the client with information about books available in a given language based on the Gutenberg library (which holds classic books - most of which are now in the public domain - in a wide range of languages). The service further determines the number of potential readers (as the /readership endpoint) presumed to be able to read books in that language.

The web service consists of four root paths:

<pre>
/librarystats/v1/  
/librarystats/v1/bookcount/  
/librarystats/v1/readership/  
/librarystats/v1/status/
</pre>

The REST web services used for the Book Information Service are:

**Gutendex API**  
Endpoint: http://129.241.150.113:8000/books/  
Documentation: http://129.241.150.113:8000/

**Language2Countries API**  
Endpoint: http://129.241.150.113:3000/language2countries/  
Documentation: http://129.241.150.113:3000/  

**REST Countries API**  
Endpoint: http://129.241.150.113:8080/v3.1  
Documentation: http://129.241.150.113:8080/  


The mapping of paths to Handler functions are:
<pre>
DefaultHandler    -> /librarystats/v1/  
BookcountHandler  -> /librarystats/v1/bookcount/  
ReadershipHandler -> /librarystats/v1/readership/  
StatusHandler     -> /librarystats/v1/status/  
</pre>

## **DefaultHandler**

The **`DefaultHandler`** provides information to the user on how to use the service.


## **BookcountHandler**

The **`BookcountHandler`** endpoint focuses returns the count of books for any given language, identified via country 2-letter language ISO codes (ISO 639 Set 1), as well as the number of unique authors. This can be a single as well as multiple languages (comma-separated language codes).

The endpoint expects a URL path of the format  
**`.../bookcount/?language={:two_letter_language_code+}/`**

<pre>
Example requests:
bookcount/?language=no
bookcount/?language=no,es
</pre>

The **`BookcountHandler`** function retrieves data from the GUTENDEX API to create a **`BooksOutput`** struct for each Book that matches the provided country code. The function returns a JSON-encoded response containing the **`BooksOutput`** structs for all matching Books. The output includes the language, total number of books, number of unique authors and a "fraction" that shows the percentage of books for the language compared to the total number of books in the Gutendex library.


<pre>
bookcount/?language=no,es (example output):

[
  {
     "language": "no",
     "books": 21,
     "authors": 14,
     "fraction": 0.0005
  },
  {
     "language": "es",
     "books": 846,
     "authors": 407,
     "fraction": 0.0116
  }
]

</pre>

## **ReadershipHandler**

The **`ReadershipHandler`** endpoint returns the number of potential readers for books in a given language, i.e., the population per country in which that language is official (and hence assuming that the inhabitants can potentially read it). This is reported in addition to the number of books and authors associated with a given language. 

The endpoint expects a URL path of the format  
**`.../readership/{:two_letter_language_code}{?limit={:number}}`**

<pre>
Example requests:
readership/no
readership/no/?limit=5
</pre>

The **`ReadershipHandler`** function retrieves data from the LANGUAGE2COUNTRIES API to find a list of countries that speak the queried language. It then uses the list of countries to retrieve their populations from the RESTCOUNTRIES API and their Books information from the GUTENDEX API. The information from the three API calls are used to create a **`CountryOutput`** struct for each Country that matches the provided language code. The function returns a JSON-encoded response containing the **`CountryOutput`** structs for all matching countries.

The output includes the country name, the country's ISO-code, the total number of books, unique authors and population (potential readers) for the given country.

<pre>
readership/es/?limit=3 (example output):
[
  {
    "country": "Argentina",
    "isocode": "AR",
    "books": 1,
    "authors": 1,
    "readership": 45376763
  },
  {
    "country": "Aruba",
    "isocode": "AW",
    "books": 0,
    "authors": 0,
    "readership": 106766
  },
  {
    "country": "Belize",
    "isocode": "BZ",
    "books": 0,
    "authors": 0,
    "readership": 397621
  }
]
</pre>

## **StatusHandler**

The **`StatusHandler`** endpoint indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.

The endpoint expects a URL path of the format **`.../status/`**.


The **`StatusHandler`** function sends requests to the three REST web services; Gutendex, Language2countries and RESTcountries to check their status codes. The function also grabs the version number from the url and calculates the server uptime using a uptimeHandler() function.

The function returns a JSON-encoded response containing a **`StatusOutput`** struct filled with the HTTP status responses from the three API calls, the Book Information Service version number and the server uptime since the last restart.

<pre>
example ouput:
{
   "gutendexapi": 200,
   "languageapi: 200, 
   "countriesapi": 200,
   "version": "v1",
   "uptime": 1337
}

</pre>

## **Dependencies**

Aggregated list of required dependencies:

- **`assignment-01/constants`**
- **`assignment-01/handler`**
- **`assignment-01/structs`**
- **`encoding/json`**
- **`net/http`**
- **`strconv`**
- **`strings`**
- **`sync`**
- **`time`**
- **`fmt`**
- **`log`**
- **`os`**


## **Usage**

To use the Book Information Service, run an HTTP server that listens for requests to the appropriate endpoints and calls the corresponding handler functions.
