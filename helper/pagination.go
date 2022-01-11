package helper

import (
	base_models "github.com/eosnationftw/eosn-base-api/models"
	"math"
)

func CreatePaginationMeta(total, limit, page, seed int) *base_models.PaginationResult {

	if total == 0 {
		return nil
	}

	next := 0
	if (limit * page) < total {
		next = page + 1
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	prev := 0
	if page > 1 {
		prev = page - 1
	}

	return &base_models.PaginationResult{
		Next:  next,
		Prev:  prev,
		Seed:  seed,
		Total: total,
		Pages: totalPages,
	}
}
