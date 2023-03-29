// app.go

package main

import (

	// tom: for Initialize

	"log"
	"strings"

	// tom: for route handlers
	"encoding/json"
	"net/http"

	// tom: go get required
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	stringy "github.com/gobeam/stringy"
	"github.com/samber/lo"
)

type App struct {
	Router *mux.Router
}

// func (a *App) Run(addr string) { }
func (a *App) Run(addr string) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	println("Server is listening on http://localhost:8010/search/Yoshua%20Bengio")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// http get
// has {{Short description|
// get {{Short description ->
type SearchResponse struct {
	Continue struct {
		Rvcontinue string `json:"rvcontinue"`
		Continue   string `json:"continue"`
	} `json:"continue"`
	Warnings struct {
		Main struct {
			Warnings string `json:"warnings"`
		} `json:"main"`
		Revisions struct {
			Warnings string `json:"warnings"`
		} `json:"revisions"`
	} `json:"warnings"`
	Query struct {
		Normalized []struct {
			Fromencoded bool   `json:"fromencoded"`
			From        string `json:"from"`
			To          string `json:"to"`
		} `json:"normalized"`
		Pages []struct {
			Pageid    int    `json:"pageid"`
			Ns        int    `json:"ns"`
			Title     string `json:"title"`
			Revisions []struct {
				Contentformat string `json:"contentformat"`
				Contentmodel  string `json:"contentmodel"`
				Content       string `json:"content"`
			} `json:"revisions"`
		} `json:"pages"`
	} `json:"query"`
}

type ShortResponse struct {
	results string
	exists  bool
}

// has {{Short description|
// get {{Short description ->
func getShortDescription(wikitext string) (string, error) {
	var target = "{{Short description"
	var hasShortDescription = strings.Contains(wikitext, target)
	var shortDescription string
	if hasShortDescription {
		var withoutShortDes = strings.Split(wikitext, target)
		short := strings.Split(withoutShortDes[1], "}")[0]
		shortDescription = strings.Split(short, "|")[1]
	}
	return shortDescription, nil
}

func searchWiki(text string) string {

	str := stringy.New(text)
	searchText := str.SnakeCase("?", "").Get()
	resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=query&prop=revisions&titles=" + searchText + "&rvlimit=1&formatversion=2&format=json&rvprop=content")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var response SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	var content = response.Query.Pages[0].Revisions[0].Content
	shortDes, _ := getShortDescription(content)
	return shortDes
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, nil)
	}).Methods("GET")

	a.Router.HandleFunc("/search/{search}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		search := vars["search"]
		if lo.IsEmpty(search) {
			respondWithError(w, http.StatusBadRequest, "Search text is required")
			return
		}

		println("SearchText", search)
		wikiSearchResponse := searchWiki(search)
		println("wikiSearchResponse", wikiSearchResponse)

		if lo.IsEmpty(wikiSearchResponse) {
			respondWithJSON(w, http.StatusOK, "No short description found")
		}

		respondWithJSON(w, http.StatusOK, wikiSearchResponse)

	}).Methods("GET")

}
