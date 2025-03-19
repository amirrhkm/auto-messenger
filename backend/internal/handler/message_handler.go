package handler

import (
	"auto-messenger/internal/domain"
	"auto-messenger/pkg/logger"
)

type MessageHandler struct {
	service domain.MessageService
	logger  logger.Logger
}

func NewMessageHandler(service domain.MessageService, logger logger.Logger) *MessageHandler {
	return &MessageHandler{
		service: service,
		logger:  logger,
	}
}
