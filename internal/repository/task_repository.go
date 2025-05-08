package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
)

type TaskRepository interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(taskID string) error
	GetByColumnID(columnID string) ([]models.Task, error)
	UpdatePosition(taskID string, position int) error
	AddTag(taskID, tagID string) error
	RemoveTag(taskID, tagID string) error
	AddAssignee(taskID, userID string) error

}

type taskRepo struct{}

func NewTaskRepository() TaskRepository {
	return &taskRepo{}
}

func (r *taskRepo) Create(task *models.Task) error {
	return config.DB.Create(task).Error
}

func (r *taskRepo) Update(task *models.Task) error {
	return config.DB.Save(task).Error
}

func (r *taskRepo) Delete(taskID string) error {
	return config.DB.Delete(&models.Task{}, "id = ?", taskID).Error
}

func (r *taskRepo) GetByColumnID(columnID string) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("column_id = ?", columnID).Order("position ASC").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepo) UpdatePosition(taskID string, position int) error {
	return config.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("position", position).Error
}

func (r *taskRepo) AddTag(taskID, tagID string) error {
	task := models.Task{ID: taskID}
	tag := models.Tag{ID: tagID}
	return config.DB.Model(&task).Association("Tags").Append(&tag)
}

func (r *taskRepo) RemoveTag(taskID, tagID string) error {
	task := models.Task{ID: taskID}
	tag := models.Tag{ID: tagID}
	return config.DB.Model(&task).Association("Tags").Delete(&tag)
}

func (r *taskRepo) AddAssignee(taskID, userID string) error {
	task := models.Task{ID: taskID}
	user := models.User{ID: userID}
	return config.DB.Model(&task).Association("Assignees").Append(&user)
}