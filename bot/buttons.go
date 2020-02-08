package main

import (
	"github.com/Syfaro/telegram-bot-api"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Посмотреть организаторов"),
		tgbotapi.NewKeyboardButton("Посмотреть все предстоящие концерты"),
	),
)

var cityRegKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/city"),
	),
)