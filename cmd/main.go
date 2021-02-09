package main

import (
	"log"

	"simplecrawler/internal/http"
)

func main() {
	http.RoutesMap()

	log.Fatal(http.Serve(":8080"))
}
