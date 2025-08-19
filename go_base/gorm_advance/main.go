package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:youda123@@tcp(rm-bp1y051g2ury93z6ufo.mysql.rds.aliyuncs.com:3306)/kt_ruoyi?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	/*
		db.AutoMigrate(&Post{})
		db.AutoMigrate(&User{})
		db.AutoMigrate(&Comment{}) */

	/* 	user := User{Name: "张三", Posts: []Post{Post{Name: "帖子", Comments: []Comment{Comment{Name: "评论"}}}}}
	   	db.Create(&user)
	   	post := Post{UserId: 1, Name: "帖子2", Comments: []Comment{Comment{Name: "评论2"}}}
	   	db.Create(&post) */
	comment := Comment{Model: gorm.Model{ID: 1}, PostID: 1}
	db.Delete(&comment)

	// findUserPost(db)
	// findTopPosts(db)
}

/*
	题目2:关联查询

。基于上述博客系统的模型定义。
。要求:
。编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息，。编写Go代码，使用Gorm查询评论数量最多的文章信息，
*/
func findUserPost(db *gorm.DB) {
	user := User{}
	db.Debug().Model(&user).Where("ID = ?", 1).Preload("Posts.Comments").Find(&user)
	fmt.Println(user)
}

func findTopPosts(db *gorm.DB) {
	post := Post{}
	db.Debug().Model(&post).Select("posts.*,count(posts.id) as comment_count").Joins("left join comments c on posts.id = c.post_id").Group("posts.id").Order("comment_count desc").First(&post)
	fmt.Println(post)
}

/* 题目3:钩子函数
。继续使用博客系统的模型。
。要求:
。为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
。为 Comment 模型添加一个钩子函数，在评论删除时检査文章的评论数量，如果评论数量为0，则更新文章的评论状态为"无评论”。
*/
