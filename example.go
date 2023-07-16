package main

import "github.com/gin-gonic/gin"


func main() {
    r := gin.Default() // 기본 Gin 라우터를 생성한다.
	r.GET("/ping", func(c *gin.Context) { // "/ping" 경로로 들어오는 GET 요청을 처리한다.
    c.JSON(200, gin.H{
	   "message": "pong", // "message" 필드가 "pong"인 JSON 응답을 반환한다. 
		})
	})
	r.Run() // 서버가 실행 되고 8080포트에서 요청을 기다린다.
}