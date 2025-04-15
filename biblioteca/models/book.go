package models

type Book struct {
	ID       int    `json:"id"`
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Lido     boll   `json:"lido"`
	Anotacao string `json:"anotacao"`
}
