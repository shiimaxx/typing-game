package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

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
	var timeout int
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.outStream)
	flags.IntVar(&timeout, "timeout", 60, "timeout")
	flags.IntVar(&timeout, "t", 60, "timeout(Short)")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	ch := input(os.Stdin)
	var okCount int

	timeoutCh := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		timeoutCh <- struct{}{}
	}()

	count := 1
QUESTION_LOOP:
	for ; ; count++ {
		word := randomdata.Adjective()
		fmt.Printf("question %d: %s\n", count, word)
		fmt.Print("> ")

		select {
		case v, ok := <-ch:
			if ok {
				if word == v {
					okCount++
				} else {
					fmt.Println("ng")
				}
			} else {
				break QUESTION_LOOP
			}
		case <-timeoutCh:
			fmt.Print("Timeup\n\n")
			break QUESTION_LOOP
		}
	}
	fmt.Fprintf(c.outStream, "result: %d/%d\n", okCount, count)

	return ExitCodeOK
}
