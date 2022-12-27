package main

import (
	"net/http"

	"gin-practice/models"
	"gin-practice/repository"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.POST("/me", func(ctx *gin.Context) {
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
	})

	r.GET("/me/:id", func(ctx *gin.Context) {

		var id int64

		if err := ctx.ShouldBindUri(&id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "id not found - " + err.Error(),
			})
			return
		}

		res, err := repository.GetName(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})

	r.Run(":5050")
}
