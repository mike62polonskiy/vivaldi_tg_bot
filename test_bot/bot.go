package main

import (
	"fmt"
	"log"
	"os"
	"github.com/mike62polonskiy/vivaldi_tg_bot/chat"
	"github.com/Syfaro/telegram-bot-api"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Посмотреть организаторов"),
	),
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

		var reply string

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "Посмотреть организаторов":
			grTags := getPlaces()
			fmt.Println(len(grTags))
		}


		key := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
		switch update.Message.Command() {
		case "start":
			reply = "Привет, я бот который поможет тебе не пропустить концерты в твоем городе!"
			userName := update.Message.Chat.UserName
			existUser := checkExistUser(userName)
			if existUser == "" {
				fmt.Println(existUser)
				userReg(update.Message.Chat.ID, update.Message.Chat.UserName)
			}
			key.ReplyMarkup = numericKeyboard
			bot.Send(key)
		//case "/get_groups":
		//	grTags := getPlaces()
		//	fmt.Println(grTags)
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, grTags)
		case "close":
			key.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(key)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)

	}
}