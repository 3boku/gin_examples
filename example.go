package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // gin 기본 라우터 생성
	router.LoadHTMLGlob("templates/index.html")
	//router.LoadHTMLFiles("폴더명/파일1.html", "폴더명/파일2.html")식으로 사용할 수 있다.
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	}) 	//"/index" 경로에 대한 GET 요청을 처리하는 핸들러 함수를 등록한다. 
		//이 핸들러 함수는 "index.tmpl" 템플릿 파일을 렌더링하여 HTML 응답을 반환한다. 
		//gin.H 구조체를 사용하여 템플릿에 전달할 데이터를 설정한다.
	router.Run(":8080") //8080포트에서 서버를 실행한다.
}