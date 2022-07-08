package middleware

import (
	"fmt"
	"gin_casbin_gorm/initialize"

	"github.com/gin-gonic/gin"
)

// casbin middleware 权限认证中间件
func CasbinMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userName string = ctx.GetHeader("userName")
		if userName == "" {
			fmt.Println("headers invalid")
			ctx.JSON(200, gin.H{
				"code":    401,
				"message": "Unauthorized",
				"data":    "userName: " + userName,
			})
			ctx.Abort()
			return
		}
		// 请求的path
		p := ctx.Request.URL.Path
		// 请求的方法
		m := ctx.Request.Method
		// 这里认证
		res, err := initialize.Enforcer().Enforce(userName, p, m)
		if err != nil {
			fmt.Println("no permission ")
			fmt.Println(err)
			ctx.JSON(200, gin.H{
				"code":    401,
				"message": "Unauthorized",
				"data":    "",
			})
			ctx.Abort()
			return
		}
		if !res {
			fmt.Println("permission check failed")
			ctx.JSON(200, gin.H{
				"code":    401,
				"message": "Unauthorized",
				"data":    "",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
