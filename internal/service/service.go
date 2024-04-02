package service

import (
	"context"
	"todo/internal/model"
)

type Service struct {
	repo repository
}

func New(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllFolders(ctx context.Context) ([]model.FolderWithTasks, error) {
	return s.repo.GetAllFolders(ctx)
}

func (s *Service) AddFolder(ctx context.Context, data model.NewFolder) error {
	return s.repo.AddFolder(ctx, data)
}

func (s *Service) DeleteFolder(ctx context.Context, id int) error {
	return s.repo.DeleteFolder(ctx, id)
}

func (s *Service) UpdateFolder(ctx context.Context, data model.Folder) error {
	return s.repo.UpdateFolder(ctx, data)
}

func (s *Service) AddTask(ctx context.Context, data model.NewTask) error {
	return s.repo.AddTask(ctx, data)
}

func (s *Service) DeleteTask(ctx context.Context, folderID int, taskID int) error {
	return s.repo.DeleteTask(ctx, folderID, taskID)
}

func (s *Service) UpdateTask(ctx context.Context, data model.Task) error {
	return s.repo.UpdateTask(ctx, data)
}
