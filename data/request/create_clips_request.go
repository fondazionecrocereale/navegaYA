package request

type CreateClipsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
	Url      string `validate:"required" json:"url"`

}