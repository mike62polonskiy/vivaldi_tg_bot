package main

import (
	"fmt"
	"log"
	"os"
	"github.com/Syfaro/telegram-bot-api"
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButton("start")
)

func main() {
	botToken := os.Getenv("TG_TOKEN")
	
	bot, err := tgbotapi.NewBotAPI(botToken)
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

		switch update.Message.Text {
		case "/start":
			userName := update.Message.Chat.UserName
			existUser := checkExistUser(userName)
			if existUser == "" {
				fmt.Println(existUser)
				userReg(update.Message.Chat.ID, update.Message.Chat.UserName)
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
			bot.Send(msg)
		case "/get_groups":
			grTags := getPlaces()
			fmt.Println(grTags)
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, grTags)
		}

	}
}