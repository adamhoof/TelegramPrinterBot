package main

func main() {

	telegramBot := TelegramBot{}
	fileDownloader := FileDownloader{}
	commandExecutor := CLIPrinter{}

	fileDownloader.setup()
	commandExecutor.setup()

	telegramBot.CreateBot()

	telegramBot.setDocInputHandler(&fileDownloader, &commandExecutor)

	telegramBot.StartBot()
}
