package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChebyshevDistance(t *testing.T) {
	tests := []struct {
		name string
		a    *Cell
		b    *Cell
		want int
	}{
		{
			name: "nil cells",
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			name: "nil cell A",
			a:    nil,
			b:    &Cell{X: 1, Y: 2},
			want: 0,
		},
		{
			name: "nil cell B",
			a:    &Cell{X: 1, Y: 2},
			b:    nil,
			want: 0,
		},
		{
			name: "same cell",
			a:    &Cell{X: 1, Y: 2},
			b:    &Cell{X: 1, Y: 2},
			want: 0,
		},
		{
			name: "different cells",
			a:    &Cell{X: 1, Y: 2},
			b:    &Cell{X: 3, Y: 4},
			want: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ChebyshevDistance(test.a, test.b)
			assert.Equal(t, test.want, got)
		})
	}
}
