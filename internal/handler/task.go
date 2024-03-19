package handler

import (
	"github.com/gin-gonic/gin"
	"task_tracking_service/internal/entity"
)

// @Summary		Завершение задачи
// @Tags			tasks
// @Description	Завершение задачи. Задача может быть выполнена несколько раз - зависит от параметра is_reusable.
// @ID				post-tasks-progress
// @Accept			json
// @Produce		json
// @Param			input			body		entity.TaskProgress	true	"body"
// @Success		200				{object}	Response
// @Failure		400,401,403,404	{object}	Response
// @Failure		500				{object}	Response
// @Failure		default			{object}	Response
// @Router			/task-progress/ [post]
func (h *Handler) TaskCompletion(ctx *gin.Context) {
	var input entity.TaskProgress
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
	// Завершение задания
	err = h.services.Task.TaskCompletion(&input)
	if err != nil {
		if err.Error() == "Вы уже выполнили это задание" || err.Error() == "Задание не найдено" {
			resp := Response{
				Message: err.Error(),
			}
			resp.Send(ctx, 200)
			return
		}
		resp := Response{
			Message: "Не удалось завершить задание",
		}
		resp.SendError(ctx, err, 500)
		return
	}
	// Отправка ответа
	resp := Response{
		Message: "Задание успешно завершено",
	}
	resp.Send(ctx, 200)
	return
}
