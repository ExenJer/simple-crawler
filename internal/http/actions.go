package http

import (
	"encoding/json"
	"net/http"

	"simplecrawler/internal/crawler"
	"simplecrawler/internal/http/request"
)

func crawlerAction(w http.ResponseWriter, r *http.Request) {
	var urlRequest request.URLRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&urlRequest); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)

		return
	}

	if err := urlRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)

		return
	}

	c := crawler.New(urlRequest.SliceURL())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(c.ExtractTitles()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
