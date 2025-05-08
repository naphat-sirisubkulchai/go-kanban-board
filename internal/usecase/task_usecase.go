package usecase

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
)

type TaskUsecase interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(taskID string) error
	GetByColumnID(columnID string) ([]models.Task, error)
	Reorder(taskID string, position int) error
	AddTag(taskID, tagID string) error
	RemoveTag(taskID, tagID string) error
	AssignUser(taskID, userID string) error

}

type taskUsecase struct {
	repo repository.TaskRepository
	notiRepo    repository.NotificationRepository

}

func NewTaskUsecase(r repository.TaskRepository, n repository.NotificationRepository) TaskUsecase {
	return &taskUsecase{
		repo:     r,
		notiRepo: n,
	}
}

func (u *taskUsecase) Create(task *models.Task) error {
	return u.repo.Create(task)
}

func (u *taskUsecase) Update(task *models.Task) error {
	return u.repo.Update(task)
}

func (u *taskUsecase) Delete(taskID string) error {
	return u.repo.Delete(taskID)
}

func (u *taskUsecase) GetByColumnID(columnID string) ([]models.Task, error) {
	return u.repo.GetByColumnID(columnID)
}

func (u *taskUsecase) Reorder(taskID string, position int) error {
	return u.repo.UpdatePosition(taskID, position)
}

func (u *taskUsecase) AddTag(taskID, tagID string) error {
	return u.repo.AddTag(taskID, tagID)
}

func (u *taskUsecase) RemoveTag(taskID, tagID string) error {
	return u.repo.RemoveTag(taskID, tagID)
}
func (u *taskUsecase) AssignUser(taskID, userID string) error {
	err := u.repo.AddAssignee(taskID, userID)
	if err != nil {
		return err
	}

	noti := &models.Notification{
		UserID:  userID,
		Message: "You have been assigned to a task",
	}
	return u.notiRepo.CreateNotification(noti)
}