package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func p_mian(){
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	
	r.GET("/", func (c*gin.Context)  {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "main website",
		})
	})

	r.Run(":8080")
}