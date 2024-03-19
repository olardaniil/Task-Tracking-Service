package service

import (
	"task_tracking_service/internal/entity"
	"task_tracking_service/internal/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *entity.UserInput) (int, error) {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetBalanceAndHistoryTasks(userID int) (int, []entity.Task, error) {

	balance, err := s.userRepo.GetUserBalance(userID)
	if err != nil {
		return 0, nil, err
	}

	tasks, err := s.userRepo.GetUserTasksHistoryByUserID(userID)
	if err != nil {
		return 0, nil, err
	}

	return balance, tasks, nil
}
