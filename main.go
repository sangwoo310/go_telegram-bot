package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	log.Println("Telegram alarm bot program is start")

	// bot token value 1209633330:AAEArq32PNoeMex1IhlH1mKENmLbWyrwE-Y
	bot, err := tgbotapi.NewBotAPI("토큰값")
	if err != nil {
		log.Fatalln("Telegram Bot API connect token connect fail")
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	//updates, err := bot.GetUpdatesChan(u)

	msg := tgbotapi.NewMessage(-499899417, "맥주병으로 헤드샷 날리게???")
	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)


	//for update := range updates {
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//	}
	//
	//	log.Println("updateChatId : ",  update.Message.Chat.ID)
	//
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "test message send")
	//	msg.ReplyToMessageID = update.Message.MessageID
	//
	//	bot.Send(msg)
	//}
}