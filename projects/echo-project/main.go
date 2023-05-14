package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pvhb/golang/projects/echo-project/cmd/api/handlers"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.IndexPage)
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.GetPostsHandler)
	e.GET("/posts/:id", handlers.GetPostByIDHandler)

	e.Logger.Fatal(e.Start(":6543"))
}
