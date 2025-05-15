// Author: Hakan Gunay
// Date: 2025-05-15

package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramBot struct {
	BotToken string
	ChatID   string
}

func NewTelegramBot(token, chatID string) *TelegramBot {
	return &TelegramBot{
		BotToken: token,
		ChatID:   chatID,
	}
}

func (bot *TelegramBot) SendMessage(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot.BotToken)

	payload := map[string]string{
		"chat_id": bot.ChatID,
		"text":    message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram error: status %d", resp.StatusCode)
	}

	return nil
}
