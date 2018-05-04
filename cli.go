package main

import (
	"flag"
	"fmt"
	"io"
	"time"

	"github.com/shiimaxx/typing-game/game"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	var timeout int
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.outStream)
	flags.IntVar(&timeout, "timeout", 60, "timeout")
	flags.IntVar(&timeout, "t", 60, "timeout(Short)")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	g := game.NewGame(time.Duration(timeout))
	questionCount, okCount := g.Run()
	rate := int((float64(okCount) / float64(questionCount)) * 100)
	var message string
	switch {
	case rate == 100:
		message = "Excellent!!"
	case rate >= 80:
		message = "Great!"
	case rate >= 60:
		message = "Good"
	case rate >= 50:
		message = "Soso"
	default:
		message = "Bad"
	}

	fmt.Fprintf(c.outStream, "result: %d/%d (rate: %d%%) %s\n", okCount, questionCount, rate, message)
	return ExitCodeOK
}
