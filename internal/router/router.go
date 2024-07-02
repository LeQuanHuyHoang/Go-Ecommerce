package router

import (
	c "github.com/LeQuanHuyHoang/Go-Ecommerce/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1/2024")
	{
		v1.GET("/ping", c.NewUserController().GetUserID)
	}
	return r
}
