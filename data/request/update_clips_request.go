package request

type UpdateClipsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Url      string `validate:"required" json:"url"`
}