package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"todo/internal/model"
	response "todo/internal/pkg/http"
)

func (h *Handler) AddTask(c *gin.Context) {
	const op = "handlers.AddTask"

	log := slog.With(
		slog.String("op", op),
	)

	var task model.NewTask
	err := c.BindJSON(&task)
	if errors.Is(err, io.EOF) {
		log.Info("request body is empty")
		response.WriteErrorResponse(c, err)
		return
	}
	if err != nil {
		log.Info("failed to decode request body")
		response.WriteErrorResponse(c, err)
		return
	}

	err = h.service.AddTask(ctx, task)
	if err != nil {
		log.Info("failed to added task")
		response.WriteErrorResponse(c, err)
		return
	}

	log.Info("task saved")
	c.Status(http.StatusCreated)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	const op = "handlers.UpdateTask"

	log := slog.With(
		slog.String("op", op),
	)

	var task model.Task

	err := c.BindJSON(&task)
	if errors.Is(err, io.EOF) {
		log.Info("request body is empty")
		response.WriteErrorResponse(c, err)
		return
	}
	if err != nil {
		log.Info("failed to decode request body")
		response.WriteErrorResponse(c, err)
		return
	}

	err = h.service.UpdateTask(ctx, task)
	if err != nil {
		log.Info("failed to update task")
		response.WriteErrorResponse(c, err)
		return
	}

	log.Info("task updated")
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	const op = "handler.DeleteTask"

	log := slog.With(
		slog.String("op", op),
	)

	folderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("invalid param")
		response.WriteErrorResponse(c, err)
		return
	}

	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		log.Info("invalid param")
		response.WriteErrorResponse(c, err)
		return
	}

	err = h.service.DeleteTask(ctx, folderID, taskID)
	if err != nil {
		log.Info("failed to delete task")
		response.WriteErrorResponse(c, err)
		return
	}

	log.Info("task deleted")
	c.Status(http.StatusNoContent)
}
