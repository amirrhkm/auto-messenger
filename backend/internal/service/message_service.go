package service

import (
	"context"

	"auto-messenger/internal/domain"
	"auto-messenger/pkg/logger"
	"auto-messenger/pkg/whatsapp"
)

type messageService struct {
	repo   domain.MessageRepository
	client whatsapp.Client
	logger logger.Logger
}

func NewMessageService(repo domain.MessageRepository, client whatsapp.Client, logger logger.Logger) domain.MessageService {
	return &messageService{
		repo:   repo,
		client: client,
		logger: logger,
	}
}

func toWhatsappMessage(msg domain.Message) whatsapp.Message {
	return whatsapp.Message{
		ID:          msg.ID,
		PhoneNumber: msg.PhoneNumber,
		Content:     msg.Content,
	}
}

func (s *messageService) SendScheduledMessages(ctx context.Context) error {
	messages, err := s.repo.GetScheduledMessages(ctx)
	if err != nil {
		return err
	}

	for _, msg := range messages {
		wMsg := toWhatsappMessage(msg)
		if err := s.client.SendMessage(ctx, wMsg); err != nil {
			s.logger.Error("Failed to send message", "error", err, "messageID", msg.ID)
			if err := s.repo.UpdateMessageStatus(ctx, msg.ID, "failed"); err != nil {
				s.logger.Error("Failed to update message status", "error", err, "messageID", msg.ID)
			}
			continue
		}

		if err := s.repo.UpdateMessageStatus(ctx, msg.ID, "sent"); err != nil {
			s.logger.Error("Failed to update message status", "error", err, "messageID", msg.ID)
		}

		s.logger.Info("Message sent successfully", "messageID", msg.ID)
	}

	return nil
}
