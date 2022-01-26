package main

func main() {

	telegramBot := TelegramBot{}
	fileDownloader := FileDownloader{}
	cliPrinter := CLIPrinter{}

	fileDownloader.setup()
	cliPrinter.setup()

	telegramBot.CreateBot()

	telegramBot.setDocInputHandler(&fileDownloader, &cliPrinter)

	telegramBot.StartBot()
}
