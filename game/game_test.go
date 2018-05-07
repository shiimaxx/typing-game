package game

import (
	"testing"
	"time"
)

func TestNewGame(t *testing.T) {
	var cases = []struct {
		name           string
		timeout        time.Duration
		numOfQuestions int
	}{
		{
			name:           "default",
			timeout:        time.Duration(60),
			numOfQuestions: 100,
		},
		{
			name:           "default",
			timeout:        time.Duration(0),
			numOfQuestions: 100,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			game := NewGame(c.timeout, c.numOfQuestions)
			if game.Timeout != c.timeout {
				t.Errorf("expected %d to eq %d", game.Timeout, c.timeout)
			}
			if len(game.Words) != c.numOfQuestions {
				t.Errorf("expected %d to eq %d", len(game.Words), c.numOfQuestions)
			}
		})
	}
}

func TestRun(t *testing.T) {

}
