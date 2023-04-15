package main

import (
	"github.com/FlyingDuck/library/dal"
	"github.com/FlyingDuck/library/handler"
	"github.com/FlyingDuck/library/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.AccessLog())
	r.Use(middleware.Cors())
	// route
	r.GET("/ping", handler.PingHandler)
	r.Static("/assets", "./assets")
	r.GET("/book/search", handler.SearchHandler)
	r.GET("/book/get/:id", handler.GetHandler)
	r.POST("/book/del/:id", handler.DeleteHandler)
	r.POST("/book/add", handler.AddHandler)
	r.POST("/upload", handler.UploadHandler)
	r.GET("/isbn/:isbn", handler.ISBNHandler)

	// init
	dal.Init()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
