package input

type Pagination struct {
	Limit int `form:"limit" binding:"omitempty,gte=1,lte=50" default:"20" minimum:"1" maximum:"50"` // limit results between 1 and 50
	Page  int `form:"page" binding:"omitempty,gte=1" default:"1" minimum:"1"`
}
