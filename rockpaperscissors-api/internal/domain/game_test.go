package domain_test

import (
	"rockpaperscissors-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name   string
		rounds int
	}{
		{"Create new game successfully", 1},
		{"Create new game with less than 1 round", -1},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			game := domain.NewGame(tc.rounds)

			assert.NotEmpty(t, game.ID, "Expected game.ID not to be empty")
			expectedRounds := tc.rounds
			if tc.rounds < 1 {
				expectedRounds = 1
			}
			assert.Equalf(t, expectedRounds, game.RoundsNumber, "Expected game.RoundsNumber to be %d, but got %d\n", expectedRounds, game.RoundsNumber)
			assert.Equalf(t, expectedRounds, len(game.Rounds), "Expected game.Rounds length to be %d, but got %d\n", expectedRounds, len(game.Rounds))
		})
	}
}
