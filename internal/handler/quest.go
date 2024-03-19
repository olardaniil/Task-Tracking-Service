package handler

import (
	"github.com/gin-gonic/gin"
	"task_tracking_service/internal/entity"
)

// @Summary Создание квеста
// @Tags quests
// @Description Создание квеста. В квесте может быть несколько задач. Каждая задача может быть выполнена один или несколько раз в квесте - зависит от параметра is_reusable.
// @ID post-quests
// @Accept  json
// @Produce  json
// @Param input body entity.QuestInput true "body"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /quests/ [post]
func (h *Handler) CreateQuest(ctx *gin.Context) {
	var input entity.QuestInput
	// Получение тела запроса
	if err := ctx.BindJSON(&input); err != nil {
		resp := Response{
			Message: "Неверное тело запроса",
		}
		resp.SendError(ctx, err, 400)
		return
	}
	// Валидация
	err := input.Validate()
	if err != nil {
		resp := Response{
			Message: err.Error(),
		}
		resp.Send(ctx, 400)
		return
	}
	// Создание квеста
	_, err = h.services.Quest.CreateQuest(&input)
	if err != nil {
		resp := Response{
			Message: "Не удалось создать квест",
		}
		resp.SendError(ctx, err, 500)
		return
	}
	// Отправка ответа
	resp := Response{
		Message: "Квест успешно создан",
	}
	resp.Send(ctx, 201)
	return
}

// @Summary Получить квесты
// @Tags quests
// @Description Получить квесты
// @ID get-quests
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /quests/ [get]
func (h *Handler) GetQuests(ctx *gin.Context) {
	// Получение квестов и их заданий
	quests, err := h.services.Quest.GetQuestsAndTasks()
	if err != nil {
		resp := Response{
			Message: "Не удалось получить квесты",
		}
		resp.SendError(ctx, err, 500)
		return
	}
	// Отправка ответа
	resp := Response{
		Message: "Квесты",
		Details: quests,
	}
	resp.Send(ctx, 200)
	return
}
