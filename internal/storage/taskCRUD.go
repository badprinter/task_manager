package storage

import (
	"time"

	"github.com/badprinter/task_manager/internal/models"
)

// Возвращает true и заполняет task до актуального состояния в случае успеха
// Возвращает false в случае неудачи
func (s *Storage) CreateTask(task *models.Task) error {
	timeNow := time.Now() // чтобы create_at и update_at было гарантировано одинакого
	_, err := s.db.Exec(
		"INSERT INTO task (title, author, isdeleted, create_at, update_at) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.Author, false, timeNow, timeNow)
	if err != nil {
		return err
	}

	// Заполняем task до актуального состояния
	task.IsDeleted = false
	task.Create_at = timeNow
	task.Update_at = timeNow

	return nil

}

func (s *Storage) GetAllTasks() ([]models.Task, error) {
	rows, err := s.db.Query("SELECT id, title, author FROM task WHERE isdeleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var daties []models.Task
	for rows.Next() {
		var data models.Task
		err := rows.Scan(&data.Id, &data.Title, &data.Author)
		if err != nil {
			return nil, err
		}
		daties = append(daties, data)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return daties, nil

}

func (s *Storage) GetTaskById(id int) (*models.Task, error) {
	var data models.Task
	err := s.db.QueryRow("SELECT title, author FROM task WHERE id = $1 AND isdeleted = false", id).Scan(
		&data.Title, &data.Author)

	if err != nil {
		return nil, err
	}
	data.Id = id
	return &data, nil
}

func (s *Storage) DeleteTaskById(id int) error {
	_, err := s.db.Exec("UPDATE task SET isdeleted = true WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateTask(data *models.Task) error {
	_, err := s.db.Exec("UPDATE task SET title = $1, author = $2, update_at = $3 WHERE id = $4",
		data.Title, data.Author, time.Now(), data.Id)
	if err != nil {
		return err
	}
	return nil
}
