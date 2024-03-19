package service

import (
	"task_tracking_service/internal/entity"
	"task_tracking_service/internal/repository"
)

type QuestService struct {
	questRepo repository.Quest
}

func NewQuestService(questRepo repository.Quest) *QuestService {
	return &QuestService{questRepo: questRepo}
}

func (s *QuestService) CreateQuest(quest *entity.QuestInput) (int, error) {
	return s.questRepo.CreateQuest(quest)
}

func (s *QuestService) GetQuestsAndTasks() ([]entity.Quest, error) {
	return s.questRepo.GetQuestsAndTasks()
}
