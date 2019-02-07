// main.go
package main // import "github.com/nicewook/go_modules_test"

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Modules! Hello, Travis-CI! Hello, Heroku!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Errorf("$PORT not set")
	}
	e.Logger.Fatal(e.Start(port))
}
