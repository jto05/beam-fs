package protocol

import (
	"errors"
	"io"
	"strings"
)

type Command struct {
	Name  string
	Usage string
	Short string
	Run   func(s *Session, args []string) error
}

type Session struct {
	Out io.Reader
}

var commands = []*Command{
	{
		Name:  "ls",
		Usage: "ls [path]",
		Short: "List remote files",
		Run:   runLS,
		// TODO: write runLS()
	},

	{
		Name:  "get",
		Usage: "get [filepath]",
		Short: "Get file at remote file path",
		Run:   runGET,
	},

	{
		Name:  "put",
		Usage: "put [local filepath] [remote filepath]",
		Short: "Put local file into remote dir",
		Run:   runPUT,
	},

	{
		Name:  "cd",
		Usage: "cd [path]",
		Short: "Change remote directory",
		Run:   runCD,
	},

	{
		Name:  "exit",
		Usage: "exit",
		Short: "Exit remote",
		Run:   runEXIT,
	},
}

func runLS(s *Session, args []string) error {
	return nil
}

func runCD(s *Session, args []string) error {
	return nil
}

func runGET(s *Session, args []string) error {
	return nil
}

func runPUT(s *Session, args []string) error {
	return nil
}

func runEXIT(s *Session, args []string) error {
	return nil
}

func parseArgs(name string, args []string) ([]string, error) {
	var validName bool

	// handle validity of command name
	validName = false
	for _, cmd := range commands {
		if cmd.Name == name {
			validName = true
		}
	}

	if !validName {
		return nil, errors.New("not a valid command name")
	}

	return args, nil
}

func ParseCommand(s *Session, line string) (Command, error) {
	splitLine := strings.Split(line, " ")
	cmd_name := splitLine[0]

	// TODO: handle arguments
	_, err := parseArgs(cmd_name, splitLine[0:])
	if err != nil {
		return Command{}, nil
	}

	return Command{}, nil
}
