package usecase

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
)

type BoardUsecase interface {
	CreateBoard(userID, title string) error
	UpdateBoardTitle(boardID, title string) error
	DeleteBoard(boardID string) error
	InviteMember(boardID, invitedUserID string) error
	GetBoardByName(name string) (*models.Board, error)
}

type boardUsecase struct {
	repo repository.BoardRepository
	notiRepo     repository.NotificationRepository
}

func NewBoardUsecase(r repository.BoardRepository, n repository.NotificationRepository) BoardUsecase {
	return &boardUsecase{
		repo:     r,
		notiRepo: n,
	}
}

func (u *boardUsecase) CreateBoard(userID, title string) error {
	board := &models.Board{
		Name:   title,
		OwnerID: userID,
	}
	return u.repo.CreateBoard(board)
}

func (u *boardUsecase) UpdateBoardTitle(boardID, title string) error {
	return u.repo.UpdateBoardTitle(boardID, title)
}

func (u *boardUsecase) DeleteBoard(boardID string) error {
	return u.repo.DeleteBoard(boardID)
}
func (u *boardUsecase) InviteMember(boardID, invitedUserID string) error {
	if err := u.repo.AddMember(boardID, invitedUserID); err != nil {
		return err
	}

	noti := &models.Notification{
		UserID:  invitedUserID,
		Message: "You have been invited to a board",
	}
	return u.notiRepo.CreateNotification(noti)
}
func (u *boardUsecase) GetBoardByName(name string) (*models.Board, error) {
	return u.repo.GetBoardByName(name)
}