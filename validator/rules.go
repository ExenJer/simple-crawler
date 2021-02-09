package validator

import "net/url"

func IsURL(validationURL string) bool {
	u, err := url.Parse(validationURL)

	return err == nil && u.Scheme != "" && u.Host != ""
}
