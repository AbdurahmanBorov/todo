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

func (h *Handler) GetAllFolders(c *gin.Context) {
	const op = "handlers.GetAllFolders"

	log := slog.With(
		slog.String("op", op),
	)

	all, err := h.service.GetAllFolders(ctx)
	if errors.Is(err, io.EOF) {
		log.Info("tasks not found")
		response.WriteErrorResponse(c, err)
		return
	}
	if err != nil {
		log.Info("failed to get folders")
		response.WriteErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, all)
}

func (h *Handler) AddFolder(c *gin.Context) {
	const op = "handlers.AddFolder"

	log := slog.With(
		slog.String("op", op),
	)

	var folders model.NewFolder
	err := c.BindJSON(&folders)
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

	err = h.service.AddFolder(ctx, folders)
	if err != nil {
		log.Info("failed to added folder")
		response.WriteErrorResponse(c, err)
		return
	}

	log.Info("folder saved")
	c.Status(http.StatusCreated)
}

func (h *Handler) UpdateFolder(c *gin.Context) {
	const op = "handlers.UpdateFolder"

	log := slog.With(
		slog.String("op", op),
	)

	var folder model.Folder

	err := c.BindJSON(&folder)
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

	err = h.service.UpdateFolder(ctx, folder)
	if err != nil {
		log.Info("failed to update folder")
		response.WriteErrorResponse(c, err)
		return
	}

	log.Info("folder updated")
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteFolder(c *gin.Context) {
	const op = "handler.DeleteFolder"

	log := slog.With(
		slog.String("op", op),
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("invalid param")
		response.WriteErrorResponse(c, err)
		return
	}

	err = h.service.DeleteFolder(ctx, id)
	if err != nil {
		log.Info("failed to delete folder")
		response.WriteErrorResponse(c, err)
		return
	}

	log.Info("folder deleted")
	c.Status(http.StatusNoContent)
}
