package service

import (
	"context"

	"auto-messenger/internal/domain"
	"auto-messenger/pkg/callmebot"
	"auto-messenger/pkg/logger"
)

const MessageContent = `Hey there! Just a little reminder that Iftar time is almost here! 🌅✨ Time to get your meal ready and enjoy your break.
	
_You are receiving this as part of your daily Iftar reminder!_`

type messageService struct {
	client callmebot.Client
	logger logger.Logger
}

func NewMessageService(client callmebot.Client, logger logger.Logger) domain.MessageService {
	return &messageService{
		client: client,
		logger: logger,
	}
}

func (s *messageService) SendScheduledMessage(ctx context.Context) error {
	msg := callmebot.Message{
		Content: MessageContent,
	}

	if err := s.client.SendMessage(ctx, msg); err != nil {
		s.logger.Error("(messageService.SendScheduledMessage):" + err.Error())
		return err
	}

	s.logger.Info("[Success] (messageService.SendScheduledMessage)")
	return nil
}
