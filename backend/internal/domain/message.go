package domain

import (
	"context"
	"time"
)

type Message struct {
	ID          string
	PhoneNumber string
	Content     string
	ScheduledAt time.Time
	Status      string
	CreatedAt   time.Time
}

type MessageRepository interface {
	GetScheduledMessages(ctx context.Context) ([]Message, error)
	UpdateMessageStatus(ctx context.Context, id string, status string) error
}

type MessageService interface {
	SendScheduledMessage(ctx context.Context) error
}
