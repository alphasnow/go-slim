package file_upload

import (
	"errors"
	"regexp"
	"strconv"
)

type FormData struct {
}

func (s *FormData) ParseContentRange(str string) ([3]int, error) {
	re := regexp.MustCompile(`(\d+)-(\d+)/(\d+)`)
	match := re.FindStringSubmatch(str)
	sizes := [3]int{}

	if len(match) < 4 {
		return sizes, errors.New("Content-Range is not match")
	}
	for i := 0; i < 3; i++ {
		sizes[i], _ = strconv.Atoi(match[i+1])
	}
	return sizes, nil
}
