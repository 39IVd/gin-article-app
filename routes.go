// routes.go
// routes definitions을 모아놓은 file
package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

}
