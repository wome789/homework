package main

import (
	"fmt"

	"gorm.io/gorm"
)

/* 题目1:模型定义
。假设你要开发一个博客系统，有以下几个实体:User(用户)、 Post(文章)、Comment(评论)。
。要求:。使用Gorm定义User、Post和Comment模型，其中User与Post,是一对多关系(一个用户可以发布多篇文章)，Post与Comment也是一对多关系(一篇文章可以有多个评论)。编写Go代码，使用Gorm创建这些模型对应的数据库表。 */

/* db.AutoMigrate(&Post{})
db.AutoMigrate(&User{})
db.AutoMigrate(&Comment{})
user := User{Name: "张三", Posts: []Post{Post{Name: "帖子1", Comments: []Comment{Comment{Name: "评论1"}}}}}
db.Create(&user) */

type User struct {
	gorm.Model
	Name     string
	PostSize int
	Posts    []Post
}

type Post struct {
	gorm.Model
	Name         string
	Comments     []Comment
	CommentState int `gorm:"default:1"`
	UserId       uint
}

type Comment struct {
	gorm.Model
	Name   string
	PostID uint
}

/*
*
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
*
*/
func (post *Post) AfterCreate(db *gorm.DB) (err error) {
	var count int
	db.Debug().Model(post).Select("count(1)").Where("user_id = ?", post.UserId).Find(&count)
	fmt.Println("update hook", count)
	db.Debug().Model(&User{}).Where("id = ?", post.UserId).Update("post_size", count)
	return
}

/*
*
为 Comment 模型添加一个钩子函数，在评论删除时检査文章的评论数量，如果评论数量为0，则更新文章的评论状态为"无评论”。
*
*/
func (comment *Comment) AfterDelete(db *gorm.DB) (err error) {
	var count int
	db.Debug().Model(comment).Select("count(1)").Where("post_id = ?", comment.PostID).Find(&count)
	fmt.Println("delete hook", count)
	if count <= 0 {
		db.Debug().Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_state", 0)
	}
	return
}
