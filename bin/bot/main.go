package main

import (
	"log"
//	"os"
	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	//botToken := os.Getenv("TG_TOKEN")
	bot, err := tgbotapi.NewBotAPI("1044318649:AAFWlUdqpoBanDcpCHx6A3PHqh0hnHSWbdk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Регистрация пользователя в системе."
			case "about":
				msg.Text = "Информация о боте."
			case "event-places":
				msg.Text = "Посмотреть список площадок."
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}

	}
}
