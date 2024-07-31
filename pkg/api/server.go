package api

import (
	"github.com/gin-gonic/gin"
	"messagio/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.Handler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	engine.GET("/message", userHandler.GetPeople)
	engine.POST("/message", userHandler.Post)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run("0.0.0.0:8001")
}
