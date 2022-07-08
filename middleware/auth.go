package middleware

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//用于接受登录的用户名与密码
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//jwt中payload的数据
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// JWT验证中间件
func AuthJWTMiddleWare() *jwt.GinJWTMiddleware {

	res, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",                                        //标识
		SigningAlgorithm: "HS256",                                            //加密算法 Optional(可选), default is HS256.
		Key:              []byte("Hello World"),                              //密钥
		Timeout:          time.Hour,                                          // jwt令牌有效的持续时间。可选，默认为一小时。
		MaxRefresh:       time.Hour,                                          // 刷新最大延时时间，默认为0表示不可刷新。
		PayloadFunc:      payloadFunc,                                        //登录时将调用的回调函数，这里可以定义返回jwt中的payload数据
		IdentityHandler:  identityHandler,                                    //标识处理程序函数
		Authenticator:    authenticator,                                      //在这里可以写我们的登录验证逻辑
		Authorizator:     authorizator,                                       //当用户通过token请求受限接口时，会经过这段逻辑
		Unauthorized:     unauthorized,                                       //错误时响应
		TokenLookup:      "header: Authorization, query: token, cookie: jwt", // 指定从哪里获取token 其格式为："<source>:<name>" 如有多个，用逗号隔开
		TokenHeadName:    "Bearer",                                           //TokenHeadName is a string in the header. Default value is "Bearer"
		TimeFunc:         time.Now,
	})

	if err != nil {
		fmt.Println("创建GinJWTMiddleware失败！")
		panic(err)
	}

	return res
}

// 1
func payloadFunc(data interface{}) jwt.MapClaims {
	fmt.Println("payloadFunc==这里可以定义返回jwt中的payload数据, 可以添加自定义数据到jwt的payload中")
	// 这里的接收类型，即：data.(*User)，为什么是User类型呢，这里的类型要与Authenticator函数处理完登录后返回的类型保持一致
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{
			"identityKey": v.UserName,
			"roleIdKey":   v.FirstName,
			"lastName":    v.LastName,
		}
	}
	fmt.Println("payloadFunc==没有OK，返回jwt的payload为空")
	return jwt.MapClaims{}
}

// 2 把jwt加密串，还原成保存的数据
func identityHandler(c *gin.Context) interface{} {
	fmt.Println("identityHandler==identity处理")
	claims := jwt.ExtractClaims(c)
	fmt.Println((claims["identityKey"]))
	return &User{
		UserName:  claims["identityKey"].(string),
		FirstName: claims["roleIdKey"].(string),
		LastName:  claims["lastName"].(string),
	}
}

// 1
func authenticator(c *gin.Context) (interface{}, error) {
	fmt.Println("authenticator==用户登录处理")
	var loginParam Login
	if err := c.ShouldBind(&loginParam); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userId := loginParam.Username
	pwd := loginParam.Password
	fmt.Println("authenticator==用户登录处理" + userId + ":" + pwd)
	if (userId == "admin" && pwd == "123456") || (userId == "test" && pwd == "123456") {
		// return map[string]interface{}{
		// 	"UserName":  userId,
		// 	"LastName":  "Bo-Yi",
		// 	"FirstName": "Wu",
		// }, nil
		return &User{
			UserName:  userId,
			LastName:  "Bo-Yi",
			FirstName: "Wu",
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// 2 验证从jwt加密串还原的数据是否有误，或者进行权限认证
func authorizator(data interface{}, c *gin.Context) bool {
	fmt.Println("authorizator==验证用户权限")
	if v, ok := data.(*User); ok {
		// 做一下设置或者预处理工作
		if v.UserName == "admin" {
			return true
		}
	}
	return false
}

// 3 出错后调用的函数
func unauthorized(c *gin.Context, code int, message string) {
	fmt.Println("unauthorized==出现了错误")
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

// 1表示登录时调用的函数
// 2表示验证jwt token时调用的函数
// 3表示出错是调用的函数
