package callmebot

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Config struct {
	PhoneNumber string
	ApiKey      string
}

type Client interface {
	SendMessage(ctx context.Context, msg Message) error
}

type Message struct {
	Content string
}

type CallMeBotClient struct {
	phoneNumber string
	apiKey      string
	httpClient  *http.Client
}

func NewClient(cfg Config) Client {
	return &CallMeBotClient{
		phoneNumber: cfg.PhoneNumber,
		apiKey:      cfg.ApiKey,
		httpClient:  &http.Client{},
	}
}

func (c *CallMeBotClient) SendMessage(ctx context.Context, msg Message) error {
	values := url.Values{}
	values.Set("phone", c.phoneNumber)
	values.Set("text", msg.Content)
	values.Set("apikey", c.apiKey)

	apiURL := fmt.Sprintf("https://api.callmebot.com/whatsapp.php?%s", values.Encode())

	fmt.Printf("CallMeBot API URL: %s\n", apiURL)

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("[Error] (callmebot.SendMessage): %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("[Error] (callmebot.SendMessage): %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[Error] (callmebot.SendMessage): %w", err)
	}

	fmt.Printf("CallMeBot API Response: %s\n", string(body))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("[Error] (callmebot.SendMessage): unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
