package router

import (
	"github.com/gin-gonic/gin"
)

// Essa função vai carregar as Rotas e retornar as rotas da Api
func ConfigRouter(router *gin.Engine) *gin.Engine {
	// Em caso de error esse metodo mas recuperar o 'estado' da aplicação para que ela nao pare
	router.Use(gin.Recovery())

	// Criando crupos de rotas
	main := router.Group("v1")
	{
		store := main.Group("go-fiance")
		{
			store.GET("buscar-usuario")
		}
	}
	return router
}
