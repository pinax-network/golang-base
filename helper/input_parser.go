package helper

import (
	"fmt"
	"strings"

	base_global "github.com/pinax-network/golang-base/global"
	base_input "github.com/pinax-network/golang-base/input"
	"github.com/pinax-network/golang-base/validate"
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

func ParseSearchInput(inputSearchPairs []string, allowedSearchFields []string) (searchPairs []base_input.SearchPair, err error) {

	errors := []error{}
	searchPairs = []base_input.SearchPair{}

	for _, p := range inputSearchPairs {
		split := strings.Split(p, ":")
		if len(split) != 2 {
			errors = append(errors, fmt.Errorf("invalid search pair given: '%s', needs to be of format 'field:value'", p))
			continue
		}

		if contains(allowedSearchFields, split[0]) {
			searchPairs = append(searchPairs, base_input.SearchPair{
				Attribute: split[0],
				Value:     split[1],
			})
		} else {
			errors = append(errors, fmt.Errorf("invalid search field given: '%s', needs to be one of %v", p, allowedSearchFields))
			continue
		}
	}

	if len(errors) > 0 {
		err = &validate.SearchValidationError{Errors: errors}
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
