package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Получить баланс пользователя
// @Tags users
// @Description Получить баланс пользователя
// @ID get-users-id-balance
// @Accept  json
// @Produce  json
// @Param user_id path int true "user_id"
// @Success 200 {object} Response
// @Failure 400,401,403,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /users/{user_id}/balance [get]
func (h *Handler) GetBalance(ctx *gin.Context) {
	// Получение userID из параметров запроса
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp := Response{
			Message: "Неверный ID пользователя",
		}
		resp.SendError(ctx, err, 400)
		return
	}
	// Получение баланса пользователя
	balance, tasks, err := h.services.User.GetBalanceAndHistoryTasks(userID)
	if err != nil {
		resp := Response{
			Message: "Не удалось получить баланс пользователя",
		}
		resp.SendError(ctx, err, 500)
		return
	}
	// Отправка ответа
	resp := Response{
		Message: "Баланс пользователя",
		Details: map[string]interface{}{
			"balance": balance,
			"tasks":   tasks,
		},
	}
	resp.Send(ctx, 200)
}
