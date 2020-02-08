package main

import (
	"fmt"
	"log"
	"os"
	"bytes"
	"text/template"
	"github.com/mike62polonskiy/vivaldi_tg_bot/chat"
	"github.com/Syfaro/telegram-bot-api"
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
		case "Указать город" :
			city := update.Message.Text
			userName := update.Message.Chat.UserName
			updateUserCity(city, userName)
			key := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
			key.ReplyMarkup = numericKeyboard
			reply = "Город "+city+" успешно выбран"
			bot.Send(key)
		case "Посмотреть организаторов" :
			var groupList = template.Must(template.New("grList").
				Parse(msgTemplate))
			grResult, err := msgGroups()
			if err != nil {
				log.Panic(err)
			}
			var byteMsg bytes.Buffer
			groupList.Execute(&byteMsg, grResult)
			reply = byteMsg.String()
		case "Посмотреть все предстоящие концерты" :
			var groupList = template.Must(template.New("eventList").
				Parse(eventMsgTemplate))
			eventResult, err := msgEvents()
			if err != nil {
				log.Panic(err)
			}
			var byteMsg bytes.Buffer
			groupList.Execute(&byteMsg, eventResult)
			reply = byteMsg.String()
		}


		key := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
		switch update.Message.Command() {
		case "/start":
			reply = chat.WelcomeMessage
			userName := update.Message.Chat.UserName
			sqlQuery := "SELECT username FROM tg_bot_users where username = $1"
			existUser := checkExist(sqlQuery, userName)
			chatID := update.Message.Chat.ID
			fmt.Println(chatID)
			//userReg(chatID, userName)
			if existUser == "" {
				fmt.Println("Пользователя не существует")
				userReg(chatID, userName)
			}
			key.ReplyMarkup = cityRegKeyboard
			bot.Send(key)
		case "/cityupdate":
			city := update.Message.Text
			userName := update.Message.Chat.UserName
			sqlQuery := "SELECT city FROM tg_bot_cities WHERE city ILIKE $1"
			existCity = checkExist(sqlQuery, city)

			if existCity != "" {
				updateUserCity(city, userName)
				reply = chat.SuccessCityUpdateMessage
			} else {
				reply = chat.ErrorCityUpdateMessage
			}
		case "/close":
			key.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(key)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)

	}
}