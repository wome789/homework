package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GinUser struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	GinUserID uint
	GinUser   GinUser
}

type QueryPost struct {
	PostId uint
	Title  string
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	GinUserID uint
	GinUser   GinUser
	PostID    uint
	Post      Post
}

type QueryComment struct {
	PostId uint
}

func createDb() *gorm.DB {
	dsn := "root:youda123@@tcp(rm-bp1y051g2ury93z6ufo.mysql.rds.aliyuncs.com:3306)/kt_learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

/* func main() {
	dsn := "root:youda123@@tcp(rm-bp1y051g2ury93z6ufo.mysql.rds.aliyuncs.com:3306)/kt_learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	db.AutoMigrate(&GinUser{}, &Post{}, &Comment{})
}
*/
