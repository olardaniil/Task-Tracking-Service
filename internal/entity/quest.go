package entity

import "fmt"

type Quest struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Cost  int    `json:"cost,omitempty"`
	Tasks []Task `json:"tasks,omitempty"`
}

type QuestInput struct {
	Name  string      `json:"name,omitempty"`
	Cost  int         `json:"cost,omitempty"`
	Tasks []TaskInput `json:"tasks,omitempty"`
}

func (q *QuestInput) Validate() error {
	if q.Name == "" {
		return fmt.Errorf("Отсутствует название квеста")
	}
	if len(q.Name) < 1 {
		return fmt.Errorf("Название квеста слишком короткое")
	}
	if len(q.Name) > 150 {
		return fmt.Errorf("Название квеста слишком длинное")
	}
	if q.Cost < 0 {
		return fmt.Errorf("Стоимость квеста не может быть отрицательной")
	}
	if len(q.Tasks) == 0 {
		return fmt.Errorf("Отсутствуют задания квеста")
	}
	for _, task := range q.Tasks {
		if err := task.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type QuestProgress struct {
	UserID  int `json:"user_id,omitempty"`
	QuestID int `json:"quest_id,omitempty"`
}
