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

4. `go run 파일명.go`명령어로 서버를 실행시키고 web브라우저에서 `http://localhost:8080/ping`에 접속하면 서버를 확인할 수 있다.

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

4. 파일을 실행 후 `localhost:8080/index`에 접속하면 html파일이 잘 로드 된다.

## 2일차

### 각 파일마다 라우팅 설정

**필자는 login.html파일과 post.html파일을 만들었다 그 코드들은 다음과 같다**

```html
<!DOCTYPE html>
<html lang="ko">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>login</title>
    </head>
    <body>
        아이디 입력:
        <input type="id" />
        <br />
        비밀번호 입력:
        <input type="pw" />
    </body>
</html>
```

> login.html

```html
<!DOCTYPE html>
<html lang="ko">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>포스트</title>
    </head>
    <body>
        <div style="width: 500px; height: 500px; border: 1px solid black">
            <p style="font-size: 60px">이것은 이미지임</p>
        </div>
        <h4>이것은 제목임</h4>
        <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Ad,
            asperiores aut obcaecati illo molestias consectetur ratione dolor
            provident odio amet, cupiditate voluptatem quas, totam quos fugiat
            rem pariatur minus voluptates.
        </p>
    </body>
</html>
```

> post.html

1. 여러개의 html파일을 로드할려면 `LoadHTMLFiles`를 사용해야한다.

2. ```go
   package main

   import (
       "net/http"

       "github.com/gin-gonic/gin"
   )


   func main() {
       router := gin.Default()
       router.LoadHTMLFiles("templates/index.html", "templates/login.html", "templates/post.html")
       //templates/index.html templates/login.html, templates/post.html파일들을 로드해온다.

       router.GET("/", func(c *gin.Context) {
           c.HTML(http.StatusOK, "index.html", gin.H{
           "title": "Main website",
           })
       })// 기본라우터에 index.html파일을 로드해온다.

       router.GET("/login", func(c *gin.Context) {
           c.HTML(http.StatusOK, "login.html", gin.H{
           "title": "login website",
           })
       })// /login라우터에 login.html파일을 로드해온다.

       router.GET("/post", func(c *gin.Context) {
           c.HTML(http.StatusOK, "post.html", gin.H{
           "title": "login website",
           })
       })// /post라우터에 post.html파일을 로드해온다.

       router.Run(":8080") //8080포트에서 서버를 실행한다.
       }
   ```

    이렇게 코드를 작성했다.

3. `localhost:8080/post`, `localhost:8080/login`, `localhost:8080/`에 접속해보면 html파일이 잘 로드된다.

## 3일차

### 라우터 html파일 이동하기

> > 우선 html 코드들을 이렇게 수정했다.

```html
<!DOCTYPE html>
<html lang="ko">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Mainpage</title>
    </head>
    <body>
        <h1>메인페이지</h1>
        <a href="{{.url}}">로그인페이지로 이동</a>
    </body>
</html>
```

> > index.html

1. setUouter 함수를 선언해서 라우터들을 만든다.

```go
func setRouter(router *gin.Engine) {
 router.GET("/", func(c *gin.Context) {// "/"에서 index.html파일을 로드해온다. 그리고 json파일을 보낸다.
  c.HTML(http.StatusOK, "index.html", gin.H{
   "title": "메인페이지",
   "url":   "/index",
  })
 })

 router.GET("/login", func(c *gin.Context) {// "/login"에서 login.html파일을 로드해온다. 그리고 json파일을 보낸다.
  c.HTML(http.StatusOK, "login.html", gin.H{
   "content": "login",
  })
 })
}
```

2. 메인함수를 만들어서 실행시킬 수 있게 한다.

```go
func main() {
 router := gin.Default() //gin 기본 라우터 생성
 router.LoadHTMLGlob("templates/*.html")//templates폴더에 있는 html파일들 가져오기
 setRouter(router)//위에 있는 setRouter함수 실행
 _ = router.Run(":8080")// 8080포트에서 서버 실행
}
```
