package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Suggestions struct {
	Query   string
	Results []string
}

func getSearchHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	query := r.Form.Get("q")

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf("<h2>Searching for %s...</h2>", query)))
}

func getSuggestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	query := r.Form.Get("q")

	results := Suggestions{
		Query: query,
		Results: []string{
			query + "0",
			query + "1",
			query + "2",
			query + "3",
		},
	}

	// results in [ "query", [ "suggestion 0", "suggestion 1" ... ] ]
	jsoned, err := json.Marshal([]interface{}{
		results.Query,
		results.Results,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsoned)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/suggest", getSuggestHandler)
	r.Get("/search", getSearchHandler)

	http.ListenAndServe(":8080", r)
}
