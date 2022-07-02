package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("pdfcpu split test.pdf .")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
