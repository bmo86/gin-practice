package Server

import (
	"context"
	"errors"
	"fmt"
	"gin-practice/database"
	"gin-practice/repository"
	ws "gin-practice/websocket"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *ws.Hub
}

type Broker struct {
	config *Config
	router *gin.Engine
	hub    *ws.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *ws.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	return &Broker{
		config: config,
		router: gin.New(),
		hub:    ws.NewHub(),
	}, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	b.router = gin.New()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)

	repo, err := database.NewConnectionDatabase(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	go b.hub.Run()
	repository.SetRepository(repo)

	fmt.Println("server on : ", b.config.Port)
	if err := http.ListenAndServe(":"+b.config.Port, handler); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}
