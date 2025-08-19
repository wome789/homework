package main

import "github.com/gin-gonic/gin"

var db = createDb()

func main() {

	router := gin.Default()

	userRouter := router.Group("/user")
	userRouter.POST("/register", Register(db))
	userRouter.POST("/login", Login(db))

	checkLoginRouter := router.Group("/")
	// 登录拦截
	checkLoginRouter.Use(AuthLogin())
	postRouter := checkLoginRouter.Group("/post")
	postRouter.POST("/create", addPosts(db))
	postRouter.POST("/delete", deletePosts(db))
	postRouter.POST("/update", updatePosts(db))
	postRouter.POST("/select", selectPosts(db))

	commentRouter := checkLoginRouter.Group("/comment")
	commentRouter.POST("/create", addComments(db))
	commentRouter.POST("/select", selectComments)

	router.Run(":8080")
}
