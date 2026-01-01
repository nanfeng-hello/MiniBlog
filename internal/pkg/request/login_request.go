package request

// LoginRequest 用户登入请求体
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=1,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
