package domain_test

import (
	"rockpaperscissors-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeats(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name   string
		move   domain.Move
		vsMove domain.Move
		result bool
	}{
		{"Rock beats scissors", domain.MoveRock, domain.MoveScissors, true},
		{"Rock doesn't beat paper", domain.MoveRock, domain.MovePaper, false},
		{"Rock doesn't beat rock", domain.MoveRock, domain.MoveRock, false},
		{"Rock doesn't beat invalid move", domain.MoveRock, domain.Move("invalid"), false},

		{"Paper doesn't beats scissors", domain.MovePaper, domain.MoveScissors, false},
		{"Paper doesn't beat paper", domain.MovePaper, domain.MovePaper, false},
		{"Paper doesn't beat rocks", domain.MovePaper, domain.MoveRock, true},
		{"Paper doesn't beat invalid move", domain.MoveRock, domain.Move("invalid"), false},

		{"Scissors doesn't beats scissors", domain.MoveScissors, domain.MoveScissors, false},
		{"Scissors beat paper", domain.MoveScissors, domain.MovePaper, true},
		{"Scissors doesn't beat rocks", domain.MoveScissors, domain.MoveRock, false},
		{"Scissors doesn't beat invalid move", domain.MoveRock, domain.Move("invalid"), false},

		{"Invalid move doesn't beats scissors", domain.Move("invalid"), domain.MoveScissors, false},
		{"Invalid move doesn't paper", domain.Move("invalid"), domain.MovePaper, false},
		{"Invalid move doesn't beat rocks", domain.Move("invalid"), domain.MoveRock, false},
		{"Invalid move doesn't beat invalid move", domain.Move("invalid"), domain.Move("invalid"), false},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.result, tc.move.Beats(tc.vsMove))
		})
	}
}
