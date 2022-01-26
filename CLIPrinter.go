package main

import (
	"log"
	"os/exec"
)

type CLIPrinter struct {
}

func (cliPrinter *CLIPrinter) setup() {
	cmd := exec.Command("lpoptions", "-d", "Printur")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func (cliPrinter *CLIPrinter) print(fileName string) {

	cmd := exec.Command("lp", "Docs/"+fileName)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
