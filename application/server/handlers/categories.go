package handles

import (
	"development/application/fiance/conf"
	"development/application/fiance/library"
	"development/application/fiance/server/services/response"
	"development/application/fiance/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
Essa funcao vai creiar uma nova categoria
*/
func CreateCategory(c *gin.Context) {
	var arg library.CreateCategoryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao converter o JSON", err)
	}
	store := conf.Conn()
	category, err := store.Conn.CreateCategory(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao criar categoria", err)
	}
	response.ResponseBody(c, http.StatusCreated, category)
}

/*
Essa funcao vai retornar um Slice de Categories
*/
func GetCategories(c *gin.Context) {
	var arg library.GetCategoryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar converter para json", err)
		return
	}
	store := conf.Conn()
	categories, err := store.Conn.GetCategory(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao busca categories", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, categories)
}

/*
Essa funcao vai retorna uma categoria
*/
func GetCategoryById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar converter id", err)
		return
	}
	store := conf.Conn()
	category, err := store.Conn.GetCategoryById(store.Cxt, util.ConvertUUID(id))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao busca categoria", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, category)
}

/*
Essa funcao vai retorna um Slice com as categorias de um Usuario
*/
func GetCategoriesByUserId(c *gin.Context) {
	id, err := uuid.Parse(c.Query("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar converter id", err)
		return
	}
	arg := library.GetCategoryByUserIdParams{
		UserID: util.ConvertUUID(id),
		Type:   c.Query("type"),
	}
	store := conf.Conn()
	categories, err := store.Conn.GetCategoryByUserId(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao buscar categorias", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, categories)
}

/*
Essa funcao vai atualizar uma categoria
*/
func UpdateCategory(c *gin.Context) {
	var arg library.UpdateCategoryParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		return
	}
	store := conf.Conn()
	err := store.Conn.UpdateCategory(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao deletar categoria", err)
		return
	}
	response.ResponseSuccess(c, http.StatusOK)
}
