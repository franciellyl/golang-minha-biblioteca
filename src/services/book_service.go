package services

import (
	"github.com/franciellyl/golang-minha-biblioteca/src/data"
	"github.com/franciellyl/golang-minha-biblioteca/src/models"
)

func ListarLivros() []models.Book {
	return data.Books
}

func AdicionarLivro(book models.Book) models.Book {
	data.UltimoID++
	book.ID = data.UltimoID
	book.Lido = false
	data.Books = append(data.Books, book)
	return book
}

func MarcarComoLido(id int) bool {
	for i, l := range data.Books {
		if l.ID == id {
			data.Books[i].Lido = true
			return true
		}
	}
	return false
}
