package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) == 0 {
		log.Fatal("Please put a command after this.")
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	pkc := make(chan bool)
	go monitor(pkc, cmd)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Caught %s exiting", sig)
			cmd.Process.Kill()
			os.Exit(0)
		}
	}()
	for {
		n, e := stdout.Read(buf)
		if e != nil {
			os.Exit(1)
		}
		os.Stdout.Write(buf[:n])
		pkc <- true
	}
}

func monitor(poke chan bool, cmd *exec.Cmd) {
	for {
		select {
		case <-poke:
			// a read from ch has occurred
		case <-time.After(time.Second * 5):
			// the read from ch has timed out
			log.Fatal("Timed out")
			cmd.Process.Kill()
			os.Exit(0)
		}
	}
}
