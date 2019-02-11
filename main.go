// main.go
package main // import "github.com/nicewook/go_modules_test"

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello justHS!")
	})

	// slack post print
	e.POST("/slack-slash-command", func(c echo.Context) error {
		//fmt.Println(c)
		return c.String(http.StatusOK, "Well received!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal("[MAIN]", e.Start(":"+port))
}
