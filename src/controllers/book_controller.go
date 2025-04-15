package controllers

import (
	"net/http"
	"strconv"

	"github.com/franciellyl/golang-minha-biblioteca/src/models"
	"github.com/franciellyl/golang-minha-biblioteca/src/services"
	"github.com/gin-gonic/gin"
)

func ListarLivros(c *gin.Context) {
	c.JSON(http.StatusOK, services.ListarLivros())
}

func AdicionarLivro(c *gin.Context) {
	var novoLivro models.Book
	if err := c.ShouldBindJSON(&novoLivro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	livroAdicionado := services.AdicionarLivro(novoLivro)
	c.JSON(http.StatusCreated, livroAdicionado)
}

func MarcarComoLido(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if services.MarcarComoLido(id) {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Livro marcado como lido"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Livro não encontrado"})
	}
}
