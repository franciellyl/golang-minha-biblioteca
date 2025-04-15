package controllers

import (
	"net/http"

	"github.com/franciellyl/golang-minha-biblioteca/biblioteca/models"
	"github.com/franciellyl/golang-minha-biblioteca/biblioteca/services"
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
