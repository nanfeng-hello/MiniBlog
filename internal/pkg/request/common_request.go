package request

type ByIdRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
