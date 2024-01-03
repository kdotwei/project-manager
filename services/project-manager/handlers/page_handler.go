package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func EditPage(context *gin.Context) {
	project_id := context.Param("id")
	context.HTML(http.StatusOK, "edit.html", gin.H{
		"projectID": project_id,
	})
}
