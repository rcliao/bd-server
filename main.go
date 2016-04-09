package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type (
	key struct {
		key string
	}
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Get("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Post("/api/token", func(c echo.Context) error {
		key := new(key)
		if err := c.Bind(key); err != nil {
			return err
		}
		data := []byte(key.key)
		return c.JSON(http.StatusOK, fmt.Sprintf("%x", md5.Sum(data)))
	})
	fmt.Printf("Starting Echo server at port 9000\n")
	e.Run(standard.New(":9000"))
}
