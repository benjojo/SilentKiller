package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 0 {
		log.Fatal("Please put a command after this.")
	}
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func monitor(poke chan bool) {
	for {
		select {
		case <-poke:
			// a read from ch has occurred
		case <-time.Sleep(time.Minute):
			// the read from ch has timed out
		}
	}
}
