package admin

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/mini-blog/internal/pkg/request"
	"github.com/nanfeng/mini-blog/internal/pkg/response"
	"github.com/nanfeng/mini-blog/internal/pkg/xerr"
	"github.com/nanfeng/mini-blog/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func (h *UserHandler) Register(r *gin.RouterGroup) {
	users := r.Group("/admin/users")
	{
		users.POST("", h.Create)
		users.DELETE(":id", h.Delete)
	}
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	// 1.从请求中获取参数
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, xerr.ErrInvalidParams.Code, xerr.ErrInternal.Msg)
	}

	// 2.调用 service 层
	id, err := h.svc.Create(c, &req)
	if err != nil {
		if errors.Is(err, xerr.ErrUsernameTaken) {
			response.Fail(c,
				xerr.ErrUsernameTaken.Code,
				xerr.ErrUsernameTaken.Msg)
		}
		if errors.Is(err, xerr.ErrEmailTaken) {
			response.Fail(c,
				xerr.ErrEmailTaken.Code,
				xerr.ErrEmailTaken.Msg)
		}
		response.Fail(c,
			xerr.ErrInternal.Code,
			err.Error())
		return
	}
	response.Success(c, id)
}

// Delete
func (h *UserHandler) Delete(c *gin.Context) {
	// 1.获取路径参数
	var req request.ByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.Fail(c, xerr.ErrInvalidParams.Code, err.Error())
		return
	}

	// 2.调用 service 层
	if err := h.svc.Delete(c, req.ID); err != nil {
		response.Fail(c, xerr.ErrInternal.Code, err.Error())
		return
	}

	// 3.返回成功信息
	response.Success(c, nil)
}
