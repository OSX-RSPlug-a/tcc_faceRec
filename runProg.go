package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("go", "run", "./humanRec/humRec.go", "0", "./humanRec/haarcascade_profileface.xml")

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
