package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type (
	key struct {
		Key string `json:"key"`
	}
	token struct {
		Token string `json:"token"`
	}
)

var randomData []float64

func init() {
	rand.Seed(454)
	for i := 0; i < 10000; i++ {
		randomData = append(randomData, rand.NormFloat64()*100)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Get("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Get("/api/random", func(c echo.Context) error {
		return c.JSON(http.StatusOK, randomData)
	})
	e.Post("/api/token", func(c echo.Context) error {
		key := new(key)
		if err := c.Bind(key); err != nil {
			return err
		}
		data := []byte(key.Key)
		token := &token{
			Token: fmt.Sprintf("%x", md5.Sum(data)),
		}
		return c.JSON(http.StatusOK, token)
	})
	e.Run(standard.New(":9000"))
}
