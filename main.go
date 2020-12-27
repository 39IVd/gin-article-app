package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// 1. create router
	router = gin.Default()

	// 2. Load all the templates
	// /templates/ 안의 모든 파일을 load
	// 한번 로드되면 다시 재로드하지 않음 -> 성능 향상
	router.LoadHTMLGlob("templates/*")

	// 3. Define the route(/) & route handler(func)
	// create a route for index page
	// NOTE: gin.Context : request에 대한 모든 정보를 담고 있음
	// (headers, cookies, etc)
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			// (response를 render 할 방법을 정의 : HTML, JSON 등)
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			// index.html에게 전달할 value (title)
			gin.H{"title": "Home Page"},
		)
	})
	// 4. Start serving the application
	router.Run()

	// 5. execute the application
	// terminal에서 application build -> create an executable named app :
	// go build -o app
	// app 실행 :
	// ./app

}
