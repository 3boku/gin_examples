package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/public", "./public")

	router.GET("/", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title" : "mainpage",
			"url" : "/",
		})
	})
	_ = router.Run(":8080")
}