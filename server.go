package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	ID    string `json;"id"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/user", show)
	e.Logger.Fatal(e.Start(":1300"))
}

func show(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
