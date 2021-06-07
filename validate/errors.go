package validate

import (
	"bytes"
	"strings"
)

type SortValidationError struct {
	Errors []error
}

func (s *SortValidationError) Error() string {
	buff := bytes.NewBufferString("")

	for _, e := range s.Errors {
		buff.WriteString(e.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}
