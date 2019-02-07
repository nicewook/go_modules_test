// main.go
package main // import "github.com/nicewook/go_modules_test"

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Modules!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
