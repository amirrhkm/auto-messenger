package whatsapp

import (
	"context"
	"fmt"
	"net/http"
)

type Config struct {
	AccessToken   string
	PhoneNumberID string
}

type Client interface {
	SendMessage(ctx context.Context, msg Message) error
}

type Message struct {
	ID          string
	PhoneNumber string
	Content     string
}

type whatsappClient struct {
	accessToken   string
	phoneNumberID string
	httpClient    *http.Client
}

func NewClient(cfg Config) Client {
	return &whatsappClient{
		accessToken:   cfg.AccessToken,
		phoneNumberID: cfg.PhoneNumberID,
		httpClient:    &http.Client{},
	}
}

func (c *whatsappClient) SendMessage(ctx context.Context, msg Message) error {
	fmt.Printf("Sending message to %s: %s\n", msg.PhoneNumber, msg.Content)
	return nil
}
