# Auto Messenger Bot

This is a simple Go-based bot that sends scheduled messages via WhatsApp using the CallMeBot API triggered by a cron job.

## Setup

### 1. Environment Variables

Create a `.env` file in the root of the project with the following variables:
```env
CMB_NUMBER=your_whatsapp_number_with_country_code # e.g., 60123456789
CMB_API_KEY=your_callmebot_api_key
LOG_LEVEL=info # or debug for more verbose logging
```

### 2. CallMeBot Setup

1. Register your phone number with CallMeBot:
   - Visit [CallMeBot](https://www.callmebot.com/blog/free-api-whatsapp-messages/)
   - Follow the instructions to link your WhatsApp number
   - Get your API key

2. Ensure your phone number is in international format:
   - Example: `60123456789` for Malaysia
   - Do not include `+` or `00` prefixes

### 3. Running the Bot

1. Install dependencies
```bash
go mod tidy
```

2. Run the bot
```bash
make run
```


## Configuration

### Environment Variables

| Variable     | Description                          | Example           |
|--------------|--------------------------------------|-------------------|
| CMB_NUMBER   | Your WhatsApp number with country code | 60123456789       |
| CMB_API_KEY  | Your CallMeBot API key               | 1234567890abcdef  |
| LOG_LEVEL    | Logging level (info/debug)           | info              |

## Message Content

The message content is defined in `internal/service/message_service.go`. You can modify the `MessageContent` constant to customize the message:

```go
const MessageContent = `Hey there! Just a little reminder that Iftar time is almost here! 🌅✨ Time to get your meal ready and enjoy your break.

_You are receiving this as part of your daily Iftar reminder!_`
```

## Troubleshooting

- **Apostrophe not showing**: Ensure you're using a regular apostrophe (`'`) in the message content
- **Message not delivered**: Verify your CallMeBot API key and ensure your number is properly registered
- **Invalid phone number**: Ensure the number is in international format without `+` or `00`

## Contributing

Pull requests are welcome as this repositories are not implemented following any design pattern and can be refactored as well as improved. (_E.g. implemeting cron value in environment variable_)

