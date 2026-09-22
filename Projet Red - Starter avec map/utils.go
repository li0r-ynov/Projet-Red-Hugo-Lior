package main

import (
	"os"
	"os/exec"
)

func clearCmd() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
