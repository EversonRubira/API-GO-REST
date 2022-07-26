package models

type Personalidade struct {
	Id int `json: "id"`
	Nome string `jason:"nome"`
	Historia string	`jason:"historia"`
}

var Personalidades []Personalidade
