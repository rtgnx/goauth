package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	e                       = echo.New()
	port                    = os.Getenv("PORT")
	host                    = os.Getenv("HOST")
	authBackend AuthBackend = PAM{Service: "passwd"}
)

func main() {

	e.Use(
		middleware.BasicAuth(
			func(username string, password string, c echo.Context) (bool, error) {
				return authBackend.Authenticate(Credentials{Login: username, Password: password})
			},
		),
	)

	e.GET("/auth", getAuthenticate)
	e.Logger.Debug(e.Start(fmt.Sprintf("%s:%s", host, port)))
}

func getAuthenticate(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}