package initialize

import (
	"fmt"
	c "github.com/LeQuanHuyHoang/Go-Ecommerce/internal/controller"
	"github.com/LeQuanHuyHoang/Go-Ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), AA())
	v1 := r.Group("v1/2024")
	{
		v1.GET("/ping", c.NewUserController().GetUserID)
	}
	return r
}

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before")
		c.Next()
		fmt.Println("After")
	}
}
