package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var subscriptionID string

func main() {

	token := os.Getenv("VULNERS_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Println("[***] Started bot")
	bot.Debug = true
	log.Printf("[***] Authorized account %s\n", bot.Self.UserName)
	go getSubscription(bot, &subscriptionID)

	for {
		getUpdates(bot, &subscriptionID) // need to implement go routine and have a string chan getting data out, this way of doing things sucks
		time.Sleep(360 * time.Minute)
	}

}

func getUpdates(bot *tgbotapi.BotAPI, sub *string) {
	endpoint := "https://vulners.com/api/v3/subscriptions/webhook?newest_only=true&subscriptionid=" + subscription[*sub]
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var Data VulnersResponse
	err = json.Unmarshal(body, &Data)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("[***] Results fetched: %d\n", len(Data.Data.Result))
	if len(Data.Data.Result) == 0 {
		log.Println("[***] No new data retrieved")
		return

	} else {
		for _, data := range Data.Data.Result {
			if data.Source.Title == "" || data.Source.Href == "" {
				log.Println("[***] Empty msg string")
				fmt.Println(data)
				return
			}
			msg := data.Source.Href
			toSend := tgbotapi.NewMessageToChannel("@vulnerschan", msg)
			time.Sleep(5 * time.Second)
			_, err = bot.Send(toSend)
			if err != nil {
				log.Panic(err)
			}
		}
	}

}

func getSubscription(bot *tgbotapi.BotAPI, sub *string) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "help":
				msg.Text = "Use /subscription to subscribe to a topic"
			case "start":
				msg.Text = "Simple bot for getting security news. Use /subscription to subscribe to a topic"
			case "status":
				msg.Text = "Bot is running"

			case "subscription":
				msg.ReplyMarkup = numericKeyboard
				msg.Text = "Please select a topic to subscribe"

			default:
				msg.Text = "Invalid command"
			}

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {

			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			subscriptionID = update.CallbackQuery.Data
			log.Printf("[***] Subscribed to: %s\n", subscriptionID)
			*sub = subscriptionID

			data := fmt.Sprintf("You have subscribed to %s\n", update.CallbackQuery.Data)
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}

		}
	}
}
