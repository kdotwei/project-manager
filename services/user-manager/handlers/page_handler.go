// handlers/page_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func UserPage(context *gin.Context) {
	userID := context.Param("id")
	context.HTML(http.StatusOK, "user.html", gin.H{
		"userID": userID,
	})
}
