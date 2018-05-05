package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

// Game game settings
type Game struct {
	Timeout        time.Duration
	Words          []string
	NumOfQuestions int
}

// NewGame constractor for Game
func NewGame(timeout time.Duration, numOfQuestions int) *Game {
	g := new(Game)
	g.Timeout = timeout
	for i := 0; i < numOfQuestions; i++ {
		g.Words = append(g.Words, randomdata.Adjective())
	}
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
		<-time.After(time.Duration(g.Timeout) * time.Second)
		timeoutCh <- struct{}{}
	}()

	ch := input(os.Stdin)
	okCount := 0
	questionCount := 0

QUESTION_LOOP:
	for _, word := range g.Words {
		fmt.Printf("question %d: %s\n", questionCount+1, word)
		fmt.Print("> ")

		select {
		case v, ok := <-ch:
			if ok {
				if word == v {
					okCount++
				} else {
					fmt.Println("miss")
				}
				questionCount++
			} else {
				break QUESTION_LOOP
			}
		case <-timeoutCh:
			fmt.Print("\nTimeup\n\n")
			break QUESTION_LOOP
		}
	}
	return questionCount, okCount
}
