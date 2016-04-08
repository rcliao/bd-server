package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.Get("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	fmt.Printf("Starting Echo server at port 9000")
	e.Run(standard.New(":9000"))
}
