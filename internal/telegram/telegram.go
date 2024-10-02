package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-pay-reminder/internal/logger"
	"go-pay-reminder/internal/sheets"
)

type Payload struct {
	Method    string `json:"method"`
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func SendToBot(token string, chatID int, row *sheets.Row) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/", token)

	payload := Payload{
		Method:    "sendMessage",
		ChatID:    chatID,
		Text:      prepareText(row),
		ParseMode: "HTML",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		logger.Error("Error while marshalling payload", "error", err, "payload", payload)
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		logger.Error("Error to send request to Telegram bot", "error", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		logger.Error("Error to send request to Telegram bot", "error", res.StatusCode)
		return
	}

	logger.Info("Message sent successfully", "name", row.Name)
}

func prepareText(row *sheets.Row) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Time to pay for %s\n", row.Name))
	builder.WriteString(fmt.Sprintf("days left %d\n", row.DaysLeft))
	builder.WriteString(fmt.Sprintf("<b>Login:</b>\n%s\n", row.Login))

	return builder.String()
}
