package crawler

import (
	"simplecrawler/client"
)

type Crawler struct {
	urls []string
}

func New(urls []string) Crawler {
	return Crawler{urls: urls}
}

func (c Crawler) ExtractTitles() []Response {
	r := make([]Response, 0)
	rc := make(chan Response)

	httpClient := client.Default()

	for _, u := range c.urls {
		go ParseTitle(u, httpClient, rc)
	}

	for i := 0; i < len(c.urls); i++ {
		if p := <-rc; true {
			r = append(r, p)
		}
	}

	return r
}
