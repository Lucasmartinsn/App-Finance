package handles

import (
	"development/application/fiance/conf"
	"development/application/fiance/library"
	"development/application/fiance/server/services/response"
	"development/application/fiance/server/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
Essa funcao vai criar uma nova conta
*/
func CreateAccount(c *gin.Context) {
	var arg library.CreateAccountsParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter para json", err)
		return
	}
	store := conf.Conn()
	account, err := store.Conn.CreateAccounts(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao criar conta", err)
		return
	}
	response.ResponseBody(c, http.StatusCreated, account)
}

/*
Essa funcao vai buscar as contas
*/
func GetAccountGraph(c *gin.Context) {
	id, err := uuid.Parse(c.Query("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "erro ao converter id", err)
		return
	}
	arg := library.GetAccountGraphParams{
		UserID: util.ConvertUUID(id),
		Type:   c.Query("type"),
	}
	store := conf.Conn()
	acconuntGraph, err := store.Conn.GetAccountGraph(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao buscar dados de Graph", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, acconuntGraph)
}

func GetAccountReports(c *gin.Context) {
	id, err := uuid.Parse(c.Query("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "erro ao converter id", err)
		return
	}
	arg := library.GetAccountReportsParams{
		UserID: util.ConvertUUID(id),
		Type:   c.Query("type"),
	}
	store := conf.Conn()
	acconuntReports, err := store.Conn.GetAccountReports(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao buscar dados de Reports", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, acconuntReports)
}

/*
Buasca accounts
*/
func GetAccounts(c *gin.Context) {
	var arg library.GetAccountsParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "erro ao converter para json", err)
		return
	}
	store := conf.Conn()
	arg.Title = fmt.Sprintf("%%%s%%", arg.Title)
	arg.Description = fmt.Sprintf("%%%s%%", arg.Description)

	accounts, err := store.Conn.GetAccounts(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao buscar Accounts", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, accounts)
}

/*
buasca uma account by Id
*/
func GetAccountsById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter id", err)
		return
	}
	store := conf.Conn()
	account, err := store.Conn.GetAccountsById(store.Cxt, util.ConvertUUID(id))
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao bucar account", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, account)
}

func GetAccountsFull(c *gin.Context) {
	var arg library.GetAccountsFullParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter para json", err)
		return
	}
	store := conf.Conn()
	arg.Title = fmt.Sprintf("%%%s%%", arg.Title)
	arg.Description = fmt.Sprintf("%%%s%%", arg.Description)
	account, err := store.Conn.GetAccountsFull(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao bucar accounts", err)
		return
	}
	response.ResponseBody(c, http.StatusOK, account)
}

/*
Essa função vai atualizar uma account
*/
func UpdateAccount(c *gin.Context) {
	var arg library.UpdateAccountsParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falhar ao converter para json", err)
		return
	}
	store := conf.Conn()
	err := store.Conn.UpdateAccounts(store.Cxt, arg)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, "falha ao bucar accounts", err)
		return
	}
	response.ResponseSuccess(c, http.StatusOK)
}
