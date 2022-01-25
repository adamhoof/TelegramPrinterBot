package main

func main() {

	telegramBot := TelegramBot{}
	fileDownloader := FileDownloader{}
	commandExecutor := CLICommandExecutor{}

	fileDownloader.setup()
	commandExecutor.setup()

	telegramBot.CreateBot()

	telegramBot.setDocInputHandler(&fileDownloader, &commandExecutor)

	telegramBot.StartBot()
}
