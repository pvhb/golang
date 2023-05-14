package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pvhb/golang/projects/e2e-project/authentication"
)

func main() {
	e := echo.New()

	e.GET("/", authentication.Index)

	fmt.Println("starting server on port 6543...")
	e.Start(":6543")
}
