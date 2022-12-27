package main

import (
	"gin-practice/database"
	"gin-practice/handlers"
	"gin-practice/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	addrPostgres := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"

	repo, err := database.NewConnectionDatabase(addrPostgres)
	if err != nil {
		log.Fatal(err.Error() + "aqui")
	}

	repository.SetRepository(repo)

	r := gin.Default()
	r.GET("/home", handlers.HomeHandler())
	r.POST("/me", handlers.CreatedMeHandler())
	r.GET("/me/:id", handlers.GetNameHandler())

	r.Run(":5050")
}
