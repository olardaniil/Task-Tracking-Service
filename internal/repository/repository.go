package repository

import (
	"github.com/jmoiron/sqlx"
	"task_tracking_service/internal/entity"
)

type User interface {
	CreateUser(user *entity.UserInput) (int, error)
	GetUserBalance(userID int) (int, error)
	GetUserTasksHistoryByUserID(userID int) ([]entity.Task, error)
}

type Quest interface {
	CreateQuest(quest *entity.QuestInput) (int, error)
	GetQuestsAndTasks() ([]entity.Quest, error)
}

type Task interface {
	GetTaskByID(taskID int) (*entity.Task, error)
	GetCountTaskProgress(task *entity.TaskProgress) (int, error)
	GetTaskStatusesByQuestAndUser(questID, userID int) ([]entity.TaskStatus, error)
	TaskCompletion(userID, taskID, taskCost, questID int, isLastTask bool) error
}

type Repository struct {
	User
	Quest
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserRepo(db),
		Quest: NewQuestRepo(db),
		Task:  NewTaskRepo(db),
	}
}
