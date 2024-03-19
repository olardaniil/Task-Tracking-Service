package service

import (
	"fmt"
	"task_tracking_service/internal/entity"
	"task_tracking_service/internal/repository"
)

type TaskService struct {
	taskRepo repository.Task
}

func NewTaskService(taskRepo repository.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) TaskCompletion(taskProgress *entity.TaskProgress) error {
	// Проверка на существование задания
	taskInfo, err := s.taskRepo.GetTaskByID(taskProgress.TaskID)
	if err != nil {
		return err
	}
	if taskInfo.ID == 0 {
		return fmt.Errorf("Задание не найдено")
	}
	//	Есть ли уже записи о выполнении задания
	countTaskProgress, err := s.taskRepo.GetCountTaskProgress(taskProgress)
	if err != nil {
		return err
	}
	if countTaskProgress > 0 && !taskInfo.IsReusable {
		return fmt.Errorf("Вы уже выполнили это задание")
	}
	// Проверка на завершение квеста
	taskStatuses, err := s.taskRepo.GetTaskStatusesByQuestAndUser(taskInfo.QuestID, taskProgress.UserID)
	if err != nil {
		return err
	}

	isLastTask := checkAllCompleted(taskStatuses, taskInfo.ID)
	// Транзакция
	err = s.taskRepo.TaskCompletion(taskProgress.UserID, taskProgress.TaskID, taskInfo.Cost, taskInfo.QuestID, isLastTask)
	if err != nil {
		return err
	}

	return nil
}

func checkAllCompleted(tasks []entity.TaskStatus, taskID int) bool {
	for _, task := range tasks {
		if task.TaskID == taskID && task.IsCompleted == true {
			return false
		}
		if task.TaskID == taskID {
			continue
		}
		if !task.IsCompleted {
			return false
		}
	}
	return true
}
