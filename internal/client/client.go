package client

import (
	"fmt"

	"beam-fs/internal/session"
)

func Start() {
	s := session.Session{}

	ui := NewUserInterface(s * session.Session)

	for {
		fmt.Fprintf(ui.writer, "you wrote: %s\n", line)
	}
}
