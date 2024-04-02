package api

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service
}

var (
	ctx = context.Background()
)

func New(serv service) *gin.Engine {
	h := Handler{service: serv}

	r := gin.New()

	folders := r.Group("/")
	{
		folders.GET("/", h.GetAllFolders)
		folders.POST("/", h.AddFolder)
		folders.DELETE("/:id", h.DeleteFolder)
		folders.PUT("/", h.UpdateFolder)

		tasks := r.Group(":id/task")
		{
			tasks.POST("/", h.AddTask)
			tasks.PUT("/", h.UpdateTask)
			tasks.DELETE("/:taskID", h.DeleteTask)
		}
	}

	return r
}
