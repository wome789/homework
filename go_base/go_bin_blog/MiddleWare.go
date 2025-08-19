package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* 授权验证
**/
func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, failWithCodeAndMessage(302, "未登录请重新登录"))
			c.Abort()
			return
		}
	}
}
