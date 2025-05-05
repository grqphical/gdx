package utils

import (
	"errors"
	"regexp"
)

type FileURI struct {
	Scheme string
	Path   string
}

func ParseFileURI(uri string) (*FileURI, error) {
	var result FileURI
	re := regexp.MustCompile(`(.*):\/\/(\/.*)`)

	matches := re.FindStringSubmatch(uri)

	if len(matches) > 0 {
		if len(matches) > 1 {
			result.Scheme = matches[1]
		}
		if len(matches) > 2 {
			result.Path = matches[2]
		}
	} else {
		return nil, errors.New("invalid URI")
	}

	return &result, nil
}
