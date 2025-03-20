package domain

import (
	"context"
	"time"
)

// Message represents the core message entity
type Message struct {
	ID          string
	PhoneNumber string
	Content     string
	ScheduledAt time.Time
	Status      string
	CreatedAt   time.Time
}

// MessageRepository defines the interface for message storage operations
type MessageRepository interface {
	GetScheduledMessages(ctx context.Context) ([]Message, error)
	UpdateMessageStatus(ctx context.Context, id string, status string) error
}

// MessageService defines the interface for message business logic
type MessageService interface {
	SendScheduledMessage(ctx context.Context) error
}
