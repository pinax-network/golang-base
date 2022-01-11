package base_models

type Count struct {
	Count int64
}

type PaginationResult struct {
	Next  int `json:"next_page,omitempty"`
	Prev  int `json:"prev_page,omitempty"`
	Seed  int `json:"seed,omitempty"`
	Pages int `json:"total_pages"`
	Total int `json:"total_docs"`
}
