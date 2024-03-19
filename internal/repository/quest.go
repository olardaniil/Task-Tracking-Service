package repository

import (
	"github.com/jmoiron/sqlx"
	"task_tracking_service/internal/entity"
)

type QuestRepo struct {
	db *sqlx.DB
}

func NewQuestRepo(db *sqlx.DB) *QuestRepo {
	return &QuestRepo{db: db}
}

func (r *QuestRepo) CreateQuest(quest *entity.QuestInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var questID int
	createQuestQuery := `INSERT INTO quests (name, cost) values ($1, $2) RETURNING id`

	row := tx.QueryRow(createQuestQuery, quest.Name, quest.Cost)
	err = row.Scan(&questID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createTaskQuery := `INSERT INTO tasks (quest_id, name, cost) values ($1, $2, $3)`
	for _, task := range quest.Tasks {
		_, err = tx.Exec(createTaskQuery, questID, task.Name, task.Cost)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return questID, tx.Commit()
}

func (r *QuestRepo) GetQuestsAndTasks() ([]entity.Quest, error) {
	type QuestWithTasks struct {
		QuestID        int    `json:"quest_id,omitempty"`
		QuestName      string `json:"quest_name,omitempty"`
		QuestCost      int    `json:"quest_cost,omitempty"`
		TaskID         int    `json:"task_id,omitempty"`
		TaskName       string `json:"task_name,omitempty"`
		TaskIsReusable bool   `json:"task_is_reusable,omitempty"`
		TaskCost       int    `json:"task_cost,omitempty"`
	}
	var questWithTasks []QuestWithTasks
	// Получение всех квестов и их заданий
	questsQuery := `SELECT q.id, q.name, q.cost, t.id, t.name, t.is_reusable, t.cost FROM quests q JOIN tasks t ON q.id = t.quest_id`
	rows, err := r.db.Query(questsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Запись данных в структуру
	for rows.Next() {
		var q QuestWithTasks
		err = rows.Scan(&q.QuestID, &q.QuestName, &q.QuestCost, &q.TaskID, &q.TaskName, &q.TaskIsReusable, &q.TaskCost)
		if err != nil {
			return nil, err
		}
		questWithTasks = append(questWithTasks, q)
	}

	var quests []entity.Quest
	var questsIDs []int

	for _, q := range questWithTasks {
		if !contains(questsIDs, q.QuestID) {
			quests = append(quests, entity.Quest{
				ID:    q.QuestID,
				Name:  q.QuestName,
				Cost:  q.QuestCost,
				Tasks: []entity.Task{},
			})
			questsIDs = append(questsIDs, q.QuestID)
		}
		quests[len(quests)-1].Tasks = append(quests[len(quests)-1].Tasks, entity.Task{
			ID:         q.TaskID,
			Name:       q.TaskName,
			IsReusable: q.TaskIsReusable,
			Cost:       q.TaskCost,
		})
	}

	return quests, nil
}

func contains(ds []int, id int) bool {
	for _, v := range ds {
		if v == id {
			return true
		}
	}
	return false
}
