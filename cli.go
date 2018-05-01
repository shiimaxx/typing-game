package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Pallinder/go-randomdata"
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

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	ch := input(os.Stdin)
	var okCount int

	for {
		word := randomdata.Adjective()
		fmt.Println(word)
		fmt.Print(">")

		if v, ok := <-ch; ok {
			if word == v {
				fmt.Println("ok")
				okCount++
			} else {
				fmt.Println("ng")
			}
		} else {
			break
		}
	}

	fmt.Fprintln(c.outStream, okCount)

	return ExitCodeOK
}
