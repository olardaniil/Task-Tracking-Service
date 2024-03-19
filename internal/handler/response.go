package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Details any    `json:"details,omitempty"`
}

func (r *Response) Send(ctx *gin.Context, code int) {
	ctx.JSON(code, r)
}

func (r *Response) SendError(ctx *gin.Context, err error, code int) {
	log.Println(err)
	ctx.JSON(code, r)
}
