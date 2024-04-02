package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo/internal/model"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetAllFolders(ctx context.Context) ([]model.FolderWithTasks, error) {
	const op = "repository.GetAllFolders"

	stmt, err := r.db.Prepare(`
	SELECT * FROM folders
	INNER JOIN tasks ON folders.folder_id = tasks.folder_id
`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	folders := make(map[int]model.FolderWithTasks)
	for rows.Next() {
		var folder model.Folder
		var task model.Task
		err := rows.Scan(&folder.ID, &folder.Name, &folder.Date, &task.ID, &task.Name, &task.Text, &task.Date, &task.Status, &task.FolderID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		folderWithTasks, ok := folders[folder.ID]
		if !ok {
			folderWithTasks = model.FolderWithTasks{Folder: folder, Tasks: []model.Task{}}
			folders[folder.ID] = folderWithTasks
		}

		folderWithTasks.Tasks = append(folderWithTasks.Tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	result := make([]model.FolderWithTasks, 0, len(folders))
	for _, folder := range folders {
		result = append(result, folder)
	}

	return result, nil
}

func (r *Repo) AddFolder(ctx context.Context, data model.NewFolder) error {
	const op = "repository.AddFolder"

	stmt, err := r.db.Prepare(`
	INSERT INTO folders(folde_name, folder_date) VAlUES (?, ?)
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, data.Name, data.Date)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteFolder(ctx context.Context, id int) error {
	const op = "repository.DeleteFolder"

	stmt, err := r.db.Prepare(`
	DELETE FROM folders WHERE folder_id = (?)
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) UpdateFolder(ctx context.Context, data model.Folder) error {
	const op = "repository.UpdateFolder"

	stmt, err := r.db.Prepare(`
	UPDATE folders
	SET folder_name = ?
	WHERE folder_id = ?	
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, data.Name, data.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) AddTask(ctx context.Context, data model.NewTask) error {
	const op = "repository.AddTask"

	stmt, err := r.db.Prepare(`
	INSERT INTO tasks(task_name, task_text, task_date, status, folder_id) VALUES (?, ?, ?, ?, ?) 
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, data.Name, data.Text, data.Date, data.Status, data.FolderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteTask(ctx context.Context, folderID int, taskID int) error {
	const op = "repository.DeleteTask"

	stmt, err := r.db.Prepare(`
		DELETE FROM tasks
		WHERE task_id = ? AND folder_id = ?
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, taskID, folderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) UpdateTask(ctx context.Context, data model.Task) error {
	const op = "repository.UpdateTask"

	stmt, err := r.db.Prepare(`
	UPDATE tasks
	SET task_name = ?,
		task_text = ?,
		task_date = ?,
		status = ?
	WHERE folder_id = ? AND task_id = ?
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, data.Name, data.Text, data.Date, data.Status, data.FolderID, data.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
