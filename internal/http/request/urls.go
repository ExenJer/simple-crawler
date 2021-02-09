package request

import (
	"errors"
	"fmt"

	"simplecrawler/validator"
)

type plainURL string

type URLRequest struct {
	URLs []plainURL `json:"urls"`
}

const maxUrls = 20

func (ur URLRequest) Validate() error {
	if len(ur.URLs) > maxUrls {
		return errors.New("error validation: `urls` more then 20")
	}

	for i, url := range ur.URLs {
		if err := url.Validate(); err != nil {
			return fmt.Errorf("error validation: index[%d]: %s", i, err.Error())
		}
	}

	return nil
}

func (pu plainURL) Validate() error {
	if !validator.IsURL(pu.String()) {
		return fmt.Errorf("invalid url %q", pu.String())
	}

	return nil
}

func (pu plainURL) String() string {
	return string(pu)
}

func (ur URLRequest) SliceURL() []string {
	urls := make([]string, len(ur.URLs))

	for i, u := range ur.URLs {
		urls[i] = u.String()
	}

	return urls
}
