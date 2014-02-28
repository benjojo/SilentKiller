package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
