package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

type Game struct {
	Timeout time.Duration
}

// NewGame constractor for Game
func NewGame(timeout time.Duration) *Game {
	g := new(Game)
	g.Timeout = timeout
	return g
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

// Run run game
func (g *Game) Run() (int, int) {
	timeoutCh := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(g.Timeout) * time.Second)
		timeoutCh <- struct{}{}
	}()

	ch := input(os.Stdin)
	var okCount int

	questionCount := 0
QUESTION_LOOP:
	for ; ; questionCount++ {
		word := randomdata.Adjective()
		fmt.Printf("question %d: %s\n", questionCount+1, word)
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
	return questionCount, okCount
}
