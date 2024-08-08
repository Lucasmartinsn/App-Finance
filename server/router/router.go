package router

import (
	handles "development/application/fiance/server/handlers"
	"development/application/fiance/server/middlewares"

	"github.com/gin-gonic/gin"
)

// Essa função vai carregar as Rotas e retornar as rotas da Api
func ConfigRouter(router *gin.Engine) *gin.Engine {
	// Em caso de error esse metodo mas recuperar o 'estado' da aplicação para que ela nao pare
	router.Use(gin.Recovery())

	// Criando crupos de rotas
	main := router.Group("v1/go-fiance", middlewares.CORSMiddleware())
	{
		login := main.Group("")
		{
			login.POST("login", handles.LoginUser)
		}

		user := main.Group("", middlewares.Auth_default())
		{
			user.GET("buscar-usuarios", handles.GetUsers)
			user.GET("buscar-usuario/:id", handles.GetUserById)
			user.POST("usuario", handles.CreateUser)
			user.PUT("usuario/:id", handles.UpdateUser)
			user.DELETE("usuario/:id", handles.DeleteUser)
		}

		category := main.Group("", middlewares.Auth_default())
		{
			category.GET("buscar-category/:id", handles.GetCategoryById)
			category.GET("buscar-categorys", handles.GetCategoriesByUserId)
			category.POST("buscar-categorys", handles.GetCategories)
			category.POST("category", handles.CreateCategory)
			category.PUT("category", handles.UpdateCategory)
		}
		account := main.Group("", middlewares.Auth_default())
		{
			account.GET("buscar-account/:id", handles.GetAccountsById)
			account.GET("buscar-account/reports", handles.GetAccountReports)
			account.GET("buscar-account/graph", handles.GetAccountGraph)
			account.POST("buscar-account/full", handles.GetAccountsFull)
			account.POST("buscar-account", handles.GetAccounts)
			account.POST("account", handles.CreateAccount)
			account.PUT("account", handles.UpdateAccount)
		}
	}
	return router
}
