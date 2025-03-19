package repository

import (
	"context"
	"database/sql"
	"strconv"

	"auto-messenger/internal/db"
	"auto-messenger/internal/domain"
)

type messageRepository struct {
	queries *db.Queries
}

func NewMessageRepository(dbConn *sql.DB) domain.MessageRepository {
	return &messageRepository{
		queries: db.New(dbConn),
	}
}

func (r *messageRepository) GetScheduledMessages(ctx context.Context) ([]domain.Message, error) {
	dbMessages, err := r.queries.GetScheduledMessages(ctx)
	if err != nil {
		return nil, err
	}

	var messages []domain.Message
	for _, dbMsg := range dbMessages {
		messages = append(messages, domain.Message{
			ID:          strconv.Itoa(int(dbMsg.ID)),
			PhoneNumber: dbMsg.PhoneNumber,
			Content:     dbMsg.Content,
		})
	}

	return messages, nil
}

func (r *messageRepository) UpdateMessageStatus(ctx context.Context, id string, status string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return r.queries.UpdateMessageStatus(ctx, db.UpdateMessageStatusParams{
		ID:     int32(idInt),
		Status: sql.NullString{String: status, Valid: true},
	})
}
