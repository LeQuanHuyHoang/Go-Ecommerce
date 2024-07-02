package controller

import (
	"github.com/LeQuanHuyHoang/Go-Ecommerce/internal/service"
	"github.com/LeQuanHuyHoang/Go-Ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (ctl UserController) GetUserID(c *gin.Context) {
	response.SuccessResponse(c, 20001, "")
}
