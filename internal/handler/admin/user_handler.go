package admin

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		users.PUT("", h.Update)
		users.GET(":id", h.GetById)
		users.GET("/list", h.GetUserList)
		users.POST("/page", h.PageQuery)
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
	id, _ := uuid.Parse(req.ID)
	// 2.调用 service 层
	if err := h.svc.Delete(c, id); err != nil {
		response.Fail(c, xerr.ErrInternal.Code, err.Error())
		return
	}

	// 3.返回成功信息
	response.Success(c, nil)
}

func (h *UserHandler) Update(c *gin.Context) {
	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, xerr.ErrInvalidParams.Code, err.Error())
		return
	}

	if err := h.svc.Update(c.Request.Context(), &req); err != nil {
		response.Fail(c, xerr.ErrUserNotFount.Code, err.Error())
		return
	}

	response.Success(c, nil)
}

// GetById
func (h *UserHandler) GetById(c *gin.Context) {
	var req request.ByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	id, _ := uuid.Parse(req.ID)
	user, err := h.svc.GetById(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, user)
}

// GetUserList
func (h *UserHandler) GetUserList(c *gin.Context) {
	users, err := h.svc.GetUserList(c.Request.Context())
	if err != nil {
		response.Fail(c, xerr.ErrInternal.Code, err.Error())
	}

	response.Success(c, &users)
}

// PageQuery
func (h *UserHandler) PageQuery(c *gin.Context) {
	var page_query request.UserPageQuery
	if err := c.ShouldBindJSON(&page_query); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	page, err := h.svc.PageQuery(c.Request.Context(), &page_query)
	if err != nil {
		response.Fail(c, xerr.ErrInternal.Code, err.Error())
		return
	}

	response.Success(c, page)
}
