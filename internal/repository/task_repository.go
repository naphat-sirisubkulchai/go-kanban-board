package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"gorm.io/gorm"
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

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepo) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepo) Delete(taskID string) error {
	return r.db.Delete(&models.Task{}, "id = ?", taskID).Error
}

func (r *taskRepo) GetByColumnID(columnID string) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("column_id = ?", columnID).Order("position ASC").Find(&tasks).Error
	return tasks, err
}

func (r *taskRepo) UpdatePosition(taskID string, position int) error {
	return r.db.Model(&models.Task{}).Where("id = ?", taskID).Update("position", position).Error
}

func (r *taskRepo) AddTag(taskID, tagID string) error {
	task := models.Task{ID: taskID}
	tag := models.Tag{ID: tagID}
	return r.db.Model(&task).Association("Tags").Append(&tag)
}

func (r *taskRepo) RemoveTag(taskID, tagID string) error {
	task := models.Task{ID: taskID}
	tag := models.Tag{ID: tagID}
	return r.db.Model(&task).Association("Tags").Delete(&tag)
}

func (r *taskRepo) AddAssignee(taskID, userID string) error {
	task := models.Task{ID: taskID}
	user := models.User{ID: userID}
	return r.db.Model(&task).Association("Assignees").Append(&user)
}
