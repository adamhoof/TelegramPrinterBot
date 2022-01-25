package main

import (
	"log"
	"os/exec"
)

type CLICommandExecutor struct {
}

func (commandExecutor *CLICommandExecutor) setup() {

}

func (commandExecutor *CLICommandExecutor) print(fileName string) {

	cmd := exec.Command("lp", "Docs/"+fileName)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
