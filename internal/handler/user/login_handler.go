package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nanfeng/mini-blog/internal/pkg/request"
	"github.com/nanfeng/mini-blog/internal/pkg/response"
	"github.com/nanfeng/mini-blog/internal/pkg/xerr"
	"github.com/nanfeng/mini-blog/internal/service"
)

// LoginHandler
type LoginHandler struct {
	svc *service.UserService
}

// LoginRouter
func (h *LoginHandler) Register(r *gin.RouterGroup) {
	routers := r.Group("")
	{
		routers.POST("/login", h.Login)
	}
}

// NewLoginHandler 新建 LoginHandler 实例
func NewLoginHandler(svc *service.UserService) *LoginHandler {
	return &LoginHandler{
		svc: svc,
	}
}

// Login 用户登入
// @Param
// @Asses json
// @Produce json
// @Router /user/login [post]
func (h *LoginHandler) Login(c *gin.Context) {
	// 1.从请求中获取参数
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 2.调用 service 层
	token, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		response.Fail(c, xerr.ErrBadRequest.Code, err.Error())
		return
	}
	// 3.返回token给前端
	response.Success(c, token)
}
