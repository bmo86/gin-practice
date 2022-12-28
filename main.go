package main

import (
	"context"
	"gin-practice/handlers"
	server "gin-practice/server"
	"log"

	"github.com/gin-gonic/gin"
)

func bindRoutes(s server.Server, r *gin.Engine) {
	r.GET("/home", handlers.HomeHandler(s))
	r.POST("/me", handlers.CreatedMeHandler(s))
	r.GET("/me/:id", handlers.GetNameHandler(s))
	r.GET("/ws", handlers.HandlerWsGin(s))
}

func main() {

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        "5050",
		DatabaseUrl: "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable",
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(bindRoutes)
}
