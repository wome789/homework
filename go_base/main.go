package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type TestRequest struct {
	Param1 string
	Param2 string
	Param3 int
	Param4 string `json:"param4" form:"param4" binding:"required,testValidater"`
}

/* var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
} */

func middleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestParam := TestRequest{}
		if err := c.ShouldBindBodyWithJSON(&requestParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if requestParam.Param1 == "param1" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "param1Error"})
			c.Abort()
			return
		}

	}
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("testValidater", testValidater)
	}

	{
		v1Router := router.Group("v1")
		v1Router.GET("/test", func(c *gin.Context) {
			c.Redirect(http.StatusInternalServerError, "https://www.baidu.com")
		})

	}
	{
		v2Router := router.Group("v2")

		v2Router.POST("/test", middleFunc(), func(c *gin.Context) {
			/* user := c.MustGet(gin.AuthUserKey).(string)
			if secret, ok := secrets[user]; ok {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
			} */
			requestParam := TestRequest{}
			if err := c.ShouldBindBodyWithJSON(&requestParam); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			fmt.Println(requestParam)

			c.JSON(http.StatusOK, gin.H{
				"key": "value",
			})
		})
	}

	router.Run(":8080")
}

var testValidater validator.Func = func(fl validator.FieldLevel) bool {
	param := fl.Field().Interface()
	if param == "1" {
		return false
	}
	return true
}
