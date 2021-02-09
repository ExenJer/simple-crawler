package parser

import "golang.org/x/net/html"

type Extractor interface {
	Extract() string
}

type TitleTokenExtractor struct {
	t html.Token
}

func (tte TitleTokenExtractor) Extract() string {
	return tte.t.Data
}

func NewTitleTokenExtractor(t html.Token) Extractor {
	return TitleTokenExtractor{t: t}
}
