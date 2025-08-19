package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
* 创建帖子
**/
func addPosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post := Post{}
		if err := c.ShouldBindBodyWithJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, failWitMessage(err.Error()))
			c.Abort()
			return
		}

		if post.Title == "" || post.Content == "" {
			c.JSON(http.StatusBadRequest, failWitMessage("post.Title 和 post.Content 参数不允许为空"))
			c.Abort()
			return
		}

		user := GinUser{Model: gorm.Model{ID: post.GinUserID}}
		if err := db.Debug().Model(&GinUser{}).First(&user); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("用户不存在 请检查后重试"))
			c.Abort()
			return
		}

		if err := db.Create(&post); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("创建帖子失败,errorInfo="+err.Error.Error()))
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, Ok())
	}
}

/** 删除帖子 **/
func deletePosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post := Post{}
		if err := c.ShouldBindBodyWithJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, failWitMessage(err.Error()))
			c.Abort()
			return
		}
		if post.ID == 0 {
			c.JSON(http.StatusBadRequest, failWitMessage("帖子id字段为空"))
			c.Abort()
			return
		}

		if post.GinUserID == 0 {
			c.JSON(http.StatusBadRequest, failWitMessage("用户id不允许为空"))
			c.Abort()
			return
		}

		existPost := Post{Model: gorm.Model{ID: post.ID}}
		if err := db.Debug().Model(&post).First(&existPost); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("帖子不存在，请检查后重试"))
			c.Abort()
			return
		}

		if existPost.GinUserID != post.GinUserID {
			c.JSON(http.StatusBadRequest, failWitMessage("非拥有者 不允许进行修改"))
			c.Abort()
			return
		}

		if err := db.Delete(&post); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("删除帖子失败,errorInfo="+err.Error.Error()))
		}

		c.JSON(http.StatusCreated, Ok())
	}
}

/** 修改帖子**/
func updatePosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post := Post{}
		if err := c.ShouldBindBodyWithJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, failWitMessage(err.Error()))
			c.Abort()
			return
		}

		if post.ID == 0 {
			c.JSON(http.StatusBadRequest, failWitMessage("帖子id字段为空"))
			c.Abort()
			return
		}

		existPost := Post{Model: gorm.Model{ID: post.ID}}
		if err := db.Model(&post).First(&existPost); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("帖子不存在，请检查后重试"))
			c.Abort()
			return
		}

		if existPost.GinUserID != post.GinUserID {
			c.JSON(http.StatusBadRequest, failWitMessage("非拥有者 不允许进行修改"))
			c.Abort()
			return
		}
		existPost.Content = post.Content
		existPost.Title = post.Title

		if err := db.Save(&existPost); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("更新帖子失败,errorInfo="+err.Error.Error()))
		}

		c.JSON(http.StatusCreated, Ok())
	}
}

func selectPosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryPost := QueryPost{}
		if err := c.ShouldBindBodyWithJSON(&queryPost); err != nil {
			c.JSON(http.StatusBadRequest, failWitMessage(err.Error()))
			c.Abort()
			return
		}
		findPost := []Post{}

		prepareSql := db.Debug().Model(&Post{})
		if queryPost.PostId != 0 {
			prepareSql.Where("id = ?", queryPost.PostId)
		}
		if queryPost.Title != "" {
			prepareSql.Where("title like %?%", queryPost.Title)
		}

		if err := prepareSql.Find(&findPost); err.Error != nil {
			c.JSON(http.StatusBadRequest, failWitMessage("帖子不存在，请检查后重试"))
			c.Abort()
			return
		}

		resultInfo := make(map[string]interface{})
		resultInfo["data"] = findPost
		resultInfo["length"] = len(findPost)

		c.JSON(http.StatusCreated, OkWithData(resultInfo))
	}
}
