package main

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"io/ioutil"
	"strings"
	"time"
)

type TelegramBot struct {
	bot *tb.Bot
}

const homiesId = "-764263562"
const telegramServer = "https://api.telegram.org/file/bot"

type Group struct {
	id string
}

func (group *Group) Recipient() string {
	return group.id
}

var homies = Group{id: homiesId}

func (telegramBot *TelegramBot) CreateBot() {

	token, err := ioutil.ReadFile("Auth/BotToken")
	formattedToken := strings.Split(string(token), "\n")
	telegramBot.bot, err = tb.NewBot(tb.Settings{
		Token: formattedToken[0],
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
		Verbose: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("bot created")
}

func (telegramBot *TelegramBot) setDocInputHandler(fileDownloader *FileDownloader, commandExecutor *CLICommandExecutor) {

	telegramBot.bot.Handle(tb.OnDocument, func(c tb.Context) error {

		file, err := telegramBot.bot.FileByID(c.Message().Document.FileID)
		if err != nil {
			return err
		}

		fileURL := telegramServer + telegramBot.bot.Token + "/" + file.FilePath

		fileName := c.Message().Document.FileName
		err = fileDownloader.download(fileURL, fileName)
		commandExecutor.print(fileName)
		if err != nil {
			return err
		}
		return nil
	})
}

func (telegramBot *TelegramBot) StartBot() {
	telegramBot.bot.Start()
}
