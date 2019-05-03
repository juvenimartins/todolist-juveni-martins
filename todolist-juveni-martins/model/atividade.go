package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Atividade struct {
	Id        string `db:"Id" json:"id" form:"id" query:"id"`
	Data      string `db:"Data" json:"data" form:"data" query:"data"`
	Titulo    string `db:"Titulo" json:"titulo" form:"titulo" query:"titulo"`
	Descricao string `db:"Descricao" json:"descricao" form:"descricao" query:"descricao"`
}

type Atividades struct {
	Atividades []Atividade `json:"atividades"`
}

func GetAll() ([]Atividade, error) {
	atividades := []Atividade{}
	db = ConectDb()
	defer db.Close()

	err := db.Select(&atividades, "SELECT Id, Data, Titulo, Descricao FROM atividade ORDER BY Data")

	return atividades, err
}

func GetById(id string) ([]Atividade, error) {
	atividade := []Atividade{}
	db = ConectDb()
	defer db.Close()

	err := db.Select(&atividade, "SELECT Id, Data, Titulo, Descricao FROM atividade WHERE Id = ?", id)
	return atividade, err
}

func IsertAtividade(a Atividade) error {
	db = ConectDb()
	defer db.Close()

	_, err := db.Exec("INSERT INTO atividade (Data, Titulo, Descricao) VALUES (?, ?, ?)",
		a.Data, a.Titulo, a.Descricao)
	return err
}

func UpdateAtividade(a Atividade) error {
	db = ConectDb()
	defer db.Close()

	_, err := db.Exec("UPDATE atividade SET Data = ?, Titulo = ?, Descricao = ? WHERE Id = ?",
		a.Data, a.Titulo, a.Descricao, a.Id)

	return err
}

func DeleteAtividade(id string) error {
	db = ConectDb()
	defer db.Close()

	_, err := db.Exec("DELETE FROM atividade WHERE Id = ?", id)

	return err
}
