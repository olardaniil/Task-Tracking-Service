package entity

import (
	"fmt"
)

type Task struct {
	ID         int    `json:"id,omitempty" db:"id"`
	QuestID    int    `json:"quest_id,omitempty" db:"quest_id"`
	Name       string `json:"name,omitempty" db:"name"`
	IsReusable bool   `json:"is_reusable,omitempty" db:"is_reusable"`
	Cost       int    `json:"cost,omitempty" db:"cost"`
}

type TaskInput struct {
	Name       string `json:"name,omitempty"`
	IsReusable bool   `json:"is_reusable,omitempty"`
	Cost       int    `json:"cost,omitempty"`
}

func (t *TaskInput) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("отсутствует название задания")
	}
	if len(t.Name) < 1 {
		return fmt.Errorf("Название задания слишком короткое")
	}
	if len(t.Name) > 150 {
		return fmt.Errorf("Название задания слишком длинное")
	}
	if t.Cost < 0 {
		return fmt.Errorf("Стоимость задания не может быть отрицательной")
	}
	return nil

}

type TaskProgress struct {
	UserID int `json:"user_id,omitempty" db:"user_id"`
	TaskID int `json:"task_id,omitempty" db:"task_id"`
}

func (t *TaskProgress) Validate() error {
	if t.UserID == 0 {
		return fmt.Errorf("Отсутствует ID пользователя")
	}
	if t.TaskID == 0 {
		return fmt.Errorf("Отсутствует ID задания")
	}
	return nil
}

type TaskStatus struct {
	TaskID      int  `json:"task_id,omitempty" db:"task_id"`
	IsCompleted bool `json:"is_completed,omitempty" db:"is_completed"`
}
