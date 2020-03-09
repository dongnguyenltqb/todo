package utils

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	Cmd := exec.Command("clear")
	Cmd.Stdout = os.Stdout
	if err := Cmd.Run(); err != nil {
		panic(err)
	}
}
