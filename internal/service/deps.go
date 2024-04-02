package service

import (
	"context"
	"todo/internal/model"
)

type repository interface {
	GetAllFolders(ctx context.Context) ([]model.FolderWithTasks, error)
	AddFolder(ctx context.Context, data model.NewFolder) error
	DeleteFolder(ctx context.Context, id int) error
	UpdateFolder(ctx context.Context, data model.Folder) error
	AddTask(ctx context.Context, data model.NewTask) error
	DeleteTask(ctx context.Context, folderID int, taskID int) error
	UpdateTask(ctx context.Context, data model.Task) error
}
