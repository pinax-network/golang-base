package base_input

type SearchPair struct {
	Attribute string
	Value     string
}

type SearchInput struct {
	SearchPairs []string `form:"search[]" binding:"dive,searchpair"`
}
