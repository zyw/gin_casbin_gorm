package gcg

import (
	"gin_casbin_gorm/middleware"
	"gin_casbin_gorm/models"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	var response models.Response
	response.Code = 0
	response.Message = "success"
	response.Data = "pong"
	c.JSON(200, response)
}

func InitRouter() {
	g := gin.Default()
	// 登录认证
	g.POST("/api/v1/login", middleware.AuthJWTMiddleWare().LoginHandler)
	// 这里的接口没有使用权限认证中间件
	version1 := g.Group("/api/v1", middleware.AuthJWTMiddleWare().MiddlewareFunc())
	{
		version1.GET("/ping", ping)
	}
	// 接口使用权限认证中间件
	version2 := g.Group("/api/v2").Use(middleware.AuthJWTMiddleWare().MiddlewareFunc()).Use(middleware.CasbinMiddleWare())
	{
		version2.GET("ping", ping)
	}
	_ = g.Run(":8099")
}
