package utils

import (
	"errors"
	"net/url"
)

func ValidateURL(raw string) error {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return errors.New("Invalid URL: must be a well-formed HTTP/HTTPS URL")
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("Invalid URL scheme: must be http or https")
	}
	if parsed.Host == "" {
		return errors.New("Invalid URL: missing host")
	}
	return nil
}
