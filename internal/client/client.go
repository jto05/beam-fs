package client

import (
	"log"
	"os"

	"beam-fs/internal/session"
)

func Start() {
	host := "localhost"
	port := "3000"
	in := os.Stdin
	out := os.Stdout

	s := session.NewSession(host, port, in, out)

	ui := NewUserInterface(&s)

	for {
		_, err := ui.Prompt()
		if err != nil {
			log.Println(err)
		}
	}
}
