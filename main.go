package main

import (
	initializers "github.com/aviral/go-crud/Initializers"
	"github.com/aviral/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("api/ping", controllers.CheckConnection)
	r.POST("api/post", controllers.PostsCreate)
	r.GET("api/posts", controllers.PostIndex)
	r.GET("api/posts/:id", controllers.PostShow)
	r.POST("api/post/update/:id", controllers.PostUpdate)
	r.GET("api/post/delete/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
