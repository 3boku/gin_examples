package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "메인페이지",
			"url":   "/login",
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "login",
		})
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	setRouter(router)
	_ = router.Run(":8080")
}