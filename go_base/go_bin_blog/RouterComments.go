package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
* 创建评论
**/
func addComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		comment := Comment{}
		if err := c.ShouldBindBodyWithJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, failWitMessage(err.Error()))
			c.Abort()
			return
		}

		if comment.Content == "" || comment.GinUserID == 0 || comment.PostID == 0 {
			c.JSON(http.StatusBadRequest, failWitMessage("comment.Content ||comment.GinUserID ||  comment.PostID  参数不允许为空"))
			c.Abort()
			return
		}

		user := GinUser{Model: gorm.Model{ID: comment.GinUserID}}
		if err := db.Debug().Model(&GinUser{}).First(&user); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("用户不存在 请检查后重试"))
			c.Abort()
			return
		}

		post := Post{Model: gorm.Model{ID: comment.PostID}}
		if err := db.Debug().Model(&Post{}).First(&post); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("帖子不存在 请检查后重试"))
			c.Abort()
			return
		}
		comment.PostID = post.ID
		comment.GinUserID = user.ID

		if err := db.Create(&comment); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("创建评论失败,errorInfo="+err.Error.Error()))
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, Ok())
	}
}

func selectComments(c *gin.Context) {
	queryComment := QueryComment{}
	if err := c.ShouldBindBodyWithJSON(&queryComment); err != nil {
		c.JSON(http.StatusBadRequest, failWitMessage(err.Error()))
		c.Abort()
		return
	}

	prepareSql := db.Debug().Model(&Comment{})
	if queryComment.PostId != 0 {
		prepareSql.Where("post_id = ?", queryComment.PostId)
	}

	findComment := []Comment{}
	if err := prepareSql.Find(&findComment); err.Error != nil {
		c.JSON(http.StatusBadRequest, failWitMessage("评论不存在，请检查后重试"))
		c.Abort()
		return
	}

	resultInfo := make(map[string]interface{})
	resultInfo["data"] = findComment
	resultInfo["length"] = len(findComment)

	c.JSON(http.StatusCreated, OkWithData(resultInfo))
}
