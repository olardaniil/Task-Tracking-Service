package repository

import (
	"github.com/jmoiron/sqlx"
	"task_tracking_service/internal/entity"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *entity.UserInput) (int, error) {
	var id int
	query := `INSERT INTO users (username, balance) values ($1, $2) RETURNING id`

	row := r.db.QueryRow(query, user.UserName, 0)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetUserBalance(userID int) (int, error) {
	var balance int
	query := `SELECT balance FROM users WHERE id = $1`
	err := r.db.Get(&balance, query, userID)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r *UserRepo) GetUserTasksHistoryByUserID(userID int) ([]entity.Task, error) {
	var tasks []entity.Task
	taskQuery := `
		SELECT t.id, t.quest_id, t.name, t.is_reusable, t.cost
    	FROM tasks t
    	LEFT JOIN tasks_progress tp on t.id = tp.task_id where user_id=$1;
	`
	err := r.db.Select(&tasks, taskQuery, userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
