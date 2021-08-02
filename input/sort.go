package base_input

type Direction string

const (
	Ascending  Direction = "asc"
	Descending Direction = "desc"
)

type SortPair struct {
	Attribute string
	Direction Direction
}

type SortInput struct {
	SortPairs []string `form:"sort[]" binding:"dive,sortpair"`
}
