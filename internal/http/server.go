package http

import "net/http"

func Serve(port string) error {
	return http.ListenAndServe(port, nil)
}
