package controller

import (
	"fmt"
	"todolist-juveni-martins/model"

	"github.com/labstack/echo"
)

func GetAtividades() interface{} {
	atvs, err := model.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	return atvs
}

// func GetAtividadeById(id string) []model.Atividade {
// 	atividade, err := model.GetById(id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return atividade
// }

func SaveAtividade(c echo.Context) error {
	a := model.Atividade{}
	a.Id = c.FormValue("id")
	a.Data = c.FormValue("data")
	a.Titulo = c.FormValue("titulo")
	a.Descricao = c.FormValue("descricao")

	if a.Id != "" {
		return model.UpdateAtividade(a)
	}
	return model.IsertAtividade(a)
	// err := c.Bind(a)
	// if err != nil {
	// 	return err
	// }
	// err2 := model.UpdateAtividade(a)
	// log.Println(err.Error())
	// return err2
}

func DeleteAtividade(c echo.Context) error {
	id := c.QueryParam("id")
	return model.DeleteAtividade(id)
}
