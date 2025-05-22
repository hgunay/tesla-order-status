// main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"tesla-order-status/internal/auth"
	"tesla-order-status/internal/client"
	"tesla-order-status/internal/notify"
	"tesla-order-status/internal/order"
	"tesla-order-status/internal/store"
	"tesla-order-status/internal/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Continuing without it.")
	}
}

func main() {
	dataDir := filepath.Join(".", "data")
	tokenPath := filepath.Join(dataDir, "tesla_tokens.json")
	storePath := filepath.Join(dataDir, "tesla_stores.json")
	orderPath := filepath.Join(dataDir, "tesla_orders.json")

	fmt.Println(utils.ColorText("\n> Retrieving Tesla order information...\n", "94"))
	codeVerifier, codeChallenge := auth.GenerateCodeVerifierAndChallenge()

	tokens, err := auth.LoadTokensFromFile(tokenPath)
	var accessToken string

	if err == nil && tokens.AccessToken != "" {
		if !auth.IsTokenValid(tokens.AccessToken) {
			fmt.Println(utils.ColorText("> Access token expired. Refreshing...", "93"))
			tokens = auth.RefreshTokens(tokens.RefreshToken)
			auth.SaveTokensToFile(tokens, tokenPath)
		}
		accessToken = tokens.AccessToken
	} else {
		authCode := auth.GetAuthCode(codeChallenge)
		tokens = auth.ExchangeCodeForTokens(authCode, codeVerifier)
		auth.SaveTokensToFile(tokens, tokenPath)
		accessToken = tokens.AccessToken
	}

	err = store.LoadStoreData(storePath)
	if err != nil {
		log.Fatalf("Mağaza bilgileri yüklenemedi: %v", err)
	}

	ordersList := client.RetrieveOrders(accessToken)

	var detailedOrders []order.DetailedOrder
	for _, o := range ordersList {
		details, err := client.GetOrderDetails(o.ReferenceNumber, accessToken)
		if err != nil {
			fmt.Println(utils.ColorText("Failed to get details for order "+o.ReferenceNumber, "91"))
			continue
		}
		detailedOrders = append(detailedOrders, order.DetailedOrder{Order: o, Details: details})
	}

	var differencesTelegram []string
	var differencesTerminal []string

	oldOrders, _ := order.LoadOrdersFromFile(orderPath)
	if oldOrders != nil {
		for i := range oldOrders {
			if i < len(detailedOrders) {
				fmt.Println(strings.Repeat("-", 45))
				fmt.Println(utils.CenterText("GÜNCELLEMELER", 45))
				fmt.Println(strings.Repeat("-", 45))

				diffTerm := order.CompareDicts(
					oldOrders[i].Details,
					detailedOrders[i].Details,
					fmt.Sprintf("Order %d", i),
					utils.TerminalFormatter{},
				)
				diffTelegram := order.CompareDicts(
					oldOrders[i].Details,
					detailedOrders[i].Details,
					fmt.Sprintf("Order %d", i),
					utils.TelegramFormatter{},
				)

				differencesTerminal = append(differencesTerminal, diffTerm...)
				differencesTelegram = append(differencesTelegram, diffTelegram...)

				for _, d := range diffTerm {
					fmt.Println(d)
				}
			}
		}
		_ = order.SaveOrdersToFile(detailedOrders, orderPath)
	} else {
		fmt.Print(utils.ColorText("Save current order data for future diff? (y/n): ", "93"))
		answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(strings.ToLower(answer)) == "y" {
			_ = order.SaveOrdersToFile(detailedOrders, orderPath)
		}
	}

	for _, d := range detailedOrders {
		client.DisplayAdditionalOrderDetails(d.Details)
		fmt.Println(strings.Repeat("-", 45) + "\n")
	}

	bot := notify.NewTelegramBot(os.Getenv("TELEGRAM_BOT_TOKEN"), os.Getenv("TELEGRAM_CHAT_ID"))
	if len(differencesTelegram) > 0 {
		msg := "📦🚨 Tesla sipariş bilgilerinde güncelleme var:\n\n" + strings.Join(differencesTelegram, "\n")
		fmt.Printf("\n> Sending message to Telegram bot...\n")
		fmt.Println(utils.ColorText(msg, "92"))

		_ = bot.SendMessage(msg)
	} else {
		_ = bot.SendMessage("📦🚫 Tesla sipariş bilgilerinde güncelleme yok.")
	}

	fmt.Println(utils.ColorText(fmt.Sprintf("\n> Finished processing orders at %s.\n", time.Now().Format("02.01.2006 15:04")), "94"))
}
