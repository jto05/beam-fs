package client

import (
	"bufio"
	"fmt"

	"beam-fs/internal/session"
)

type UserInterface struct {
	session *session.Session
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func NewUserInterface(s *session.Session) UserInterface {
	sc := bufio.NewScanner(s.GetReader())
	w := bufio.NewWriter(s.GetWriter())

	return UserInterface{
		session: s,
		scanner: sc,
		writer:  w,
	}
}

func (ui *UserInterface) Prompt() (string, error) {
	var line string
	appName := "beam-fs"
	prompt := fmt.Sprintf("<%s> ", appName)

	// display pronpt
	_, err := ui.writer.WriteString(prompt)
	if err != nil {
		return err
	}

	// get input
	line = ui.getInput()
	fmt.Fprintf(ui.writer, "you wrote: %s\n", line)
	return line
}

func (ui *UserInterface) getInput() string {
	if !ui.scanner.Scan() {
		return ""
	}
	input := ui.scanner.Text()
	return input
}
