package service

import (
	"task_tracking_service/internal/entity"
	"task_tracking_service/internal/repository"
)

type User interface {
	CreateUser(user *entity.UserInput) (int, error)
	GetBalanceAndHistoryTasks(userID int) (int, []entity.Task, error)
}
type Quest interface {
	CreateQuest(quest *entity.QuestInput) (int, error)
	GetQuestsAndTasks() ([]entity.Quest, error)
}

type Task interface {
	TaskCompletion(taskProgress *entity.TaskProgress) error
}

type Service struct {
	User
	Quest
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:  NewUserService(repos.User),
		Quest: NewQuestService(repos.Quest),
		Task:  NewTaskService(repos.Task),
	}
}
