package handles

import (
	"development/application/fiance/conf"
	"development/application/fiance/library"
	"development/application/fiance/server/services/encryption"
	"development/application/fiance/server/services/response"
	"development/application/fiance/server/util"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
Essa função vai criar um novo registro de usuario

URL: http://localhost:5050/v1/go-fiance/usuario
*/
func CreateUser(c *gin.Context) {
	var user library.CreateUserParams
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "Erro ao converter para JSON", err)
		return
	}
	senha, err := encryption.GenerateHash(user.Password)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "Falhar ao encript de senha", err)
		return
	}
	user.Password = senha
	store := conf.Conn()
	usuario, err := store.Conn.CreateUser(store.Cxt, user)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "Erro ao criar user", err)
		return
	}
	response.ResponseBody(c, http.StatusCreated, usuario)
}

/*
Essa funçao vai retornar um Slice de Usuarios

URL: http://localhost:5050/v1/go-fiance/buscar-usuarios
*/
func GetUsers(c *gin.Context) {
	store := conf.Conn()
	if c.Query("username") == "" {
		response.ResponseError(c, http.StatusNotAcceptable, "Verifique a URL", errors.New("formatacao de url invalida"))
		return
	}
	user, err := store.Conn.GetUser(store.Cxt, fmt.Sprintf("%%%s%%", c.Query("username")))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "Erro ao buscar users", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, user)
}

/*
Essa funçao vai retornar um Usuario

URL: http://localhost:5050/v1/go-fiance/buscar-usuario/id
*/
func GetUserById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter o Id", err)
		return
	}
	store := conf.Conn()
	user, err := store.Conn.GetUserById(store.Cxt, util.ConvertUUID(id))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "Erro ao buscar user", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, user)
}

/*
Essa função vai atualizar um Usuario

Essa url pode receber dois(2) parametros do tipo Query, sendo eles 'pass' e 'data'

URL: http://localhost:5050/v1/go-fiance/usuario/id?type=
*/
func UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter o Id", err)
		return
	}
	switch c.Query("type") {
	case "pass":
		var user library.UpdateUserPassParams
		if err := c.ShouldBindJSON(&user); err != nil {
			response.ResponseError(c, http.StatusBadRequest, "Erro ao converter para JSON", err)
			return
		}
		senha, err := encryption.GenerateHash(user.Password)
		if err != nil {
			response.ResponseError(c, http.StatusBadRequest, "Falhar ao encript de senha", err)
			return
		}
		user.Password = senha
		user.ID = util.ConvertUUID(id)
		store := conf.Conn()
		if err = store.Conn.UpdateUserPass(store.Cxt, user); err != nil {
			response.ResponseError(c, http.StatusBadRequest, "falha ao atualizar usuario", err)
			return
		}
		response.ResponseSuccess(c, http.StatusOK)

	case "data":
		var user library.UpdateUserParams
		if err := c.ShouldBindJSON(&user); err != nil {
			response.ResponseError(c, http.StatusBadRequest, "Erro ao converter para JSON", err)
			return
		}
		user.ID = util.ConvertUUID(id)
		store := conf.Conn()
		if err = store.Conn.UpdateUser(store.Cxt, user); err != nil {
			response.ResponseError(c, http.StatusBadRequest, "falha ao atualizar usuario", err)
			return
		}
		response.ResponseSuccess(c, http.StatusOK)

	default:
		response.ResponseError(c, http.StatusNotFound, "Url não existe", errors.New("falha ao buscar url"))
	}
}

/*
Essa função vai deletar um Usuario

URL: http://localhost:5050/v1/go-fiance/usuario/id
*/
func DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter o Id", err)
		return
	}
	store := conf.Conn()
	if err = store.Conn.DeleteUser(store.Cxt, util.ConvertUUID(id)); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao deletar usuario", err)
		return
	}
	response.ResponseSuccess(c, http.StatusOK)
}
