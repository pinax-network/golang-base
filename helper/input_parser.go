package helper

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/global"
	"github.com/eosnationftw/eosn-base-api/input"
	"github.com/eosnationftw/eosn-base-api/validate"
	"strings"
)

func ParsePaginationInput(input *base_input.Pagination) {
	if input.Limit < 1 {
		input.Limit = base_global.DEFAULT_LIMIT
	}
	if input.Page < 1 {
		input.Page = 1
	}
}

func ParseSortInput(inputSortPairs []string, allowedSortFields []string, defaultField string, defaultDir base_input.Direction) (sortPairs []base_input.SortPair, err error) {

	errors := []error{}
	sortPairs = []base_input.SortPair{}

	for _, p := range inputSortPairs {
		split := strings.Split(p, ":")
		dir := defaultDir

		if len(split) == 2 {
			if split[1] == string(base_input.Ascending) {
				dir = base_input.Ascending
			} else if split[1] == string(base_input.Descending) {
				dir = base_input.Descending
			} else {
				errors = append(errors, fmt.Errorf("invalid sort direction given: '%s', needs to be one of ['asc', 'desc']", p))
				continue
			}
		} else if len(split) != 1 {
			errors = append(errors, fmt.Errorf("invalid sort pair given: '%s', needs to be of format 'field:direction' or 'field'", p))
			continue
		}

		if contains(allowedSortFields, split[0]) {
			sortPairs = append(sortPairs, base_input.SortPair{
				Attribute: split[0],
				Direction: dir,
			})
		} else {
			errors = append(errors, fmt.Errorf("invalid sort field given: '%s', needs to be one of %v", p, allowedSortFields))
			continue
		}
	}

	if len(errors) > 0 {
		err = &validate.SortValidationError{Errors: errors}
	}

	if len(sortPairs) == 0 {
		sortPairs = append(sortPairs, base_input.SortPair{
			Attribute: defaultField,
			Direction: defaultDir,
		})
	}

	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
