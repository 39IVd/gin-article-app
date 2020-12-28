package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// showIndexPage : index page를 위한 handler
func showIndexPage(c *gin.Context) {
	// articleList를 fetch해 저장
	articles := getAllArticles()

	// Call the HTML method of the Context
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		// (response를 render 할 방법을 정의 : HTML, JSON 등)
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		// index.html에게 전달할 value (title)
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
