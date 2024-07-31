package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"messagio/pkg/domain/message"
)

func (h *Handler) Post(c *gin.Context) {
	var msg message.Message
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(400, gin.H{"error bind Post message": err.Error()})
		log.Error("error bind Post message %v", err.Error())
		return
	}
	result, err := h.service.PostMessage(c.Request.Context(), msg)
	if err != nil {
		c.JSON(500, gin.H{"error Post message": err.Error()})
		log.Error("error service Post message %v", err.Error())
		return
	}

	c.JSON(201, result)
	log.Info("Success registration: %v", result)
}

func (h *Handler) GetPeople(c *gin.Context) {
	messages, err := h.service.GetStaticsMessage(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Error(err.Error())
		return
	}
	c.JSON(200, gin.H{"data": messages})
	log.Info("Success GetMessage %v", messages)
	return
}
