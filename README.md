# gin_examples

go-gin을 공부할것입니다.

## 1일차

### gin 시작하기

1. `go mod init name`으로 go.mod파일 생성하기
2. `$go get -u github.com/gin-gonic/gin`로 gin 설치하기

3. ```go
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
   ```

4. `go run 파일명.go`명령어로 서버를 실행시키고 web브라우저에서 `http://localhost:8080/`에 접속하면 서버를 확인할 수 있다.

### HTML랜더링

**우선 필자는 templates폴더에 index.html파일의 코드를 이렇게 작성하였다.**

```html
<!DOCTYPE html>
<body>
    <h1>Hello Wolrd!</h1>
</body>
</html>
```

1. LoadHTMLGlob() 혹은 LoadHTMLFiles()를 사용해서 html 파일을 가져온다.
2. LoadHTMLGlob()과 LoadHTMLFiles()의 차이는 LoadHTMLGlob를 메소드를 보면 매개변수가 string형 1개만 받아오고, LoadHTMLFiles 메소드를 보면 매개변수가 string형 파일을 여러개 받아온다.

3. ```go
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
        })  //"/index" 경로에 대한 GET 요청을 처리하는 핸들러 함수를 등록한다.
        //이 핸들러 함수는 "index.html" html 파일을 렌더링하여 HTML 응답을 반환한다.
        //gin.H 구조체를 사용하여 템플릿에 전달할 데이터를 설정한다.
        router.Run(":8080") //8080포트에서 서버를 실행한다.
        }
   ```

4. 파일을 실행 후 localhost:8080/index에 접속하면 html파일이 잘 로드 된다.
