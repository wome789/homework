package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/**
* 注册用户
**/
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user GinUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, failWitMessage("Failed to hash password"))
			return
		}
		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, failWitMessage("Failed to create user"))
			return
		}

		c.JSON(http.StatusCreated, Ok())
	}

}

/**
* 登录
**/
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user GinUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var storedUser GinUser
		if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// 验证密码
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// 生成 JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       storedUser.ID,
			"username": storedUser.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte("your_secret_key"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		resultMap := make(map[string]string)
		resultMap["token"] = tokenString
		c.JSON(http.StatusOK, OkWithData(resultMap))
	}
}
