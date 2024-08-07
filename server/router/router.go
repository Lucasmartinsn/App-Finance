package router

import (
	"development/application/fiance/server/handles"

	"github.com/gin-gonic/gin"
)

// Essa função vai carregar as Rotas e retornar as rotas da Api
func ConfigRouter(router *gin.Engine) *gin.Engine {
	// Em caso de error esse metodo mas recuperar o 'estado' da aplicação para que ela nao pare
	router.Use(gin.Recovery())

	// Criando crupos de rotas
	main := router.Group("v1")
	{
		user := main.Group("go-fiance")
		{
			user.GET("buscar-usuarios", handles.GetUsers)
			user.GET("buscar-usuario/:id", handles.GetUserById)
			user.POST("usuario", handles.CreateUser)
			user.PUT("usuario/:id", handles.UpdateUser)
			user.DELETE("usuario/:id", handles.DeleteUser)
		}
	}
	return router
}
