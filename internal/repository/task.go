package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"task_tracking_service/internal/entity"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) GetTaskByID(taskID int) (*entity.Task, error) {
	var task entity.Task
	taskQuery := `SELECT * FROM tasks WHERE id = $1`
	err := r.db.Get(&task, taskQuery, taskID)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepo) GetCountTaskProgress(task *entity.TaskProgress) (int, error) {
	var countTaskProgress int
	countTaskProgressQuery := fmt.Sprintf("SELECT count(*) FROM tasks_progress WHERE user_id = $1 AND task_id = $2")
	err := r.db.Get(&countTaskProgress, countTaskProgressQuery, task.UserID, task.TaskID)
	if err != nil {
		return 0, err
	}
	return countTaskProgress, nil
}

func (r *TaskRepo) GetTaskStatusesByQuestAndUser(questID, userID int) ([]entity.TaskStatus, error) {
	query := `
        SELECT t.id, CASE WHEN tp.task_id IS NULL THEN false ELSE true END AS is_completed
        FROM tasks t
        LEFT JOIN tasks_progress tp ON t.id = tp.task_id AND tp.user_id = $1
        WHERE t.quest_id = $2
    `
	rows, err := r.db.Query(query, userID, questID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taskStatuses []entity.TaskStatus
	for rows.Next() {
		var taskStatus entity.TaskStatus
		if err := rows.Scan(&taskStatus.TaskID, &taskStatus.IsCompleted); err != nil {
			return nil, err
		}
		taskStatuses = append(taskStatuses, taskStatus)
	}
	return taskStatuses, nil
}

func (r *TaskRepo) TaskCompletion(userID, taskID, taskCost, questID int, isLastTask bool) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	// Записываем данные о выполнении задания
	taskCompletionQuery := `INSERT INTO tasks_progress (user_id, task_id) values ($1, $2)`
	_, err = tx.Exec(taskCompletionQuery, userID, taskID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if isLastTask == true {
		// Завершаем квест
		questCompletionQuery := `INSERT INTO quests_progress (user_id, quest_id) values ($1, $2)`
		_, err = tx.Exec(questCompletionQuery, userID, questID)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Обновляем баланс пользователя
		updateBalanceQuery := `UPDATE users SET balance = balance + $2 + (SELECT cost FROM quests WHERE id = $3) WHERE id = $1;`
		_, err = tx.Exec(updateBalanceQuery, userID, taskCost, questID)
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}

	// Обновляем баланс пользователя
	updateBalanceQuery := `UPDATE users SET balance = balance + $2  WHERE id = $1;`
	_, err = tx.Exec(updateBalanceQuery, userID, taskCost)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}
