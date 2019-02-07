// main.go
package main // import "github.com/nicewook/go_modules_test"

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello justHS!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("[MAIN] Server is running at 0.0.0.0:%v\n", port)
	log.Println("[MAIN]", e.Start(":"+port))
}
