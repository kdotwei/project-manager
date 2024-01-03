package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func ProjectPage(context *gin.Context) {
	project_id := context.Param("id")
	context.HTML(http.StatusOK, "project.html", gin.H{
		"projectID": project_id,
	})
}

func EditPage(context *gin.Context) {
	project_id := context.Param("id")
	context.HTML(http.StatusOK, "edit.html", gin.H{
		"projectID": project_id,
	})
}

func EditTaskPage(context *gin.Context) {
	project_id := context.Param("id")
	task_id := context.Param("taskId")

	log.Println("project_id: ", project_id)
	log.Println("task_id: ", task_id)

	context.HTML(http.StatusOK, "edit_task.html", gin.H{
		"projectID": project_id,
		"taskID":    task_id,
	})
}
