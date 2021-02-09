package parser

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Search struct {
	StartToken, EndToken html.TokenType
	Sought               atom.Atom
}

const maxTokensDepth = 1000

func TokenizePayload(payload string) *html.Tokenizer {
	return html.NewTokenizer(strings.NewReader(payload))
}

func TokenSearch(tokens *html.Tokenizer, s Search) *html.Token {
	var (
		searchValueFound, searchBreak bool
		searchValue                   *html.Token
		depth                         = 0
	)

	for {
		if searchBreak || maxTokensDepth == depth {
			break
		}

		depth++

		n := tokens.Next()

		switch n {
		default:
		case html.ErrorToken:
			searchBreak = true
		case s.StartToken:
			token := tokens.Token()

			if tagMatch(token, s.Sought) {
				searchValueFound = true
			}
		case s.EndToken:
			if searchValueFound {
				token := tokens.Token()

				searchBreak, searchValueFound = true, false
				searchValue = &token

				break
			}
		}
	}

	return searchValue
}

func tagMatch(token html.Token, sought atom.Atom) bool {
	return token.Data == sought.String()
}
