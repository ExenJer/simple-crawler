package crawler

import (
	"errors"
	"net/http"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"simplecrawler/client"
	"simplecrawler/parser"
)

var (
	ErrFindTitleTag = errors.New("cannot find a `title` tag")
)

func ParseTitle(url string, httpClient *http.Client, response chan<- Response) {
	resp, err := client.FetchBody(http.MethodGet, url, httpClient)
	if err != nil {
		response <- ErrorResponse(url, HTTPErrorCode, err.Error())

		return
	}

	title, err := BodyTitle(resp)
	if err != nil {
		response <- ErrorResponse(url, ParsingErrorCode, err.Error())

		return
	}

	response <- SuccessResponse(url, SuccessCode, title)
}

func BodyTitle(resp string) (string, error) {
	ts := parser.TokenizePayload(resp)
	t := parser.TokenSearch(ts, parser.Search{
		StartToken: html.StartTagToken,
		EndToken:   html.TextToken,
		Sought:     atom.Title,
	})

	if t == nil {
		return "", ErrFindTitleTag
	}

	ex := parser.NewTitleTokenExtractor(*t)

	return ex.Extract(), nil
}
