package handlers

import (
	"gin-practice/models"
	"gin-practice/repository"
	ws "gin-practice/websocket"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func HomeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello, my API using gin/golang",
		})
	}
}

func CreatedMeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var m models.Me

		err := ctx.ShouldBind(&m)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := models.Me{
			Name:     m.Name,
			Lastname: m.Lastname,
			Age:      m.Age,
		}

		if err = repository.CreatedMe(ctx, &data); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, ResponseMessage{
			Message: "Created Me",
		})
	}
}

func GetNameHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idReq := ctx.Param("id")

		id, err := strconv.ParseInt(idReq, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "not number",
			})
		}

		res, err := repository.GetName(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func HandlerWsGin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws.NewHub().WsHandler(ctx.Writer, ctx.Request)
	}
}
