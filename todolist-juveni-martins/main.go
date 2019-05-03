package main

import (
	"net/http"
	"todolist-juveni-martins/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func delete(c echo.Context) error {
	if err := controller.DeleteAtividade(c); err != nil {
		return c.String(http.StatusCreated, err.Error())
	}
	return c.String(http.StatusCreated, "status:ok")
}

func save(c echo.Context) error {
	if err := controller.SaveAtividade(c); err != nil {
		return c.String(http.StatusCreated, err.Error())
	}
	return c.String(http.StatusCreated, "status:ok")
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	// Lista Atividades
	e.GET("/atividade", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, controller.GetAtividades())
	})
	// update se id != '', insert se id == ''
	e.POST("/insert", save)
	e.GET("/delete", delete)

	e.Logger.Fatal(e.Start(":1323"))
}
