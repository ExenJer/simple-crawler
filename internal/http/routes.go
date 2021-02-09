package http

import (
	"net/http"
)

func RoutesMap() {
	http.HandleFunc("/crawler", crawlerAction)
}
