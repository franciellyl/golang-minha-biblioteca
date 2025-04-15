package routes

import (
	"github.com/franciellyl/golang-minha-biblioteca/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/livros", controllers.ListarLivros)
	r.POST("/livros", controllers.AdicionarLivro)
	// r.PUT("/livros/:id/lido", controllers.MarcarComoLido)
	// r.GET("/estatisticas", controllers.Estatisticas)

	return r
}
