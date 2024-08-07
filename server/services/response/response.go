package response

import "github.com/gin-gonic/gin"

// Mensagem simples de Sucesso
func ResponseSuccess(c *gin.Context, status int) {
	c.JSON(status, nil)
}

// Mensagem de resposta com Body
func ResponseBody(c *gin.Context, status int, response any) {
	c.JSON(status, gin.H{
		"Response": response,
	})
}

// Mensagem simples de error
func ResponseError(c *gin.Context, status int, message string, err error) {
	c.JSON(status, gin.H{
		"Message": message,
		"Error":   err.Error(),
	})
}
