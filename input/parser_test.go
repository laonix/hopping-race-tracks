package input

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestParseTestCases(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     []*TestCase
		err      error
	}{
		{
			name:     "valid test cases input file",
			filePath: "../test/resource/valid.txt",
			want: []*TestCase{
				{
					ID:       1,
					GridRows: 5,
					GridCols: 5,
					Start:    CellCoordinates{X: 4, Y: 0},
					End:      CellCoordinates{X: 4, Y: 4},
					Obstacles: []Obstacle{
						{X1: 1, X2: 4, Y1: 2, Y2: 3},
					},
				},
				{
					ID:       2,
					GridRows: 3,
					GridCols: 3,
					Start:    CellCoordinates{X: 0, Y: 0},
					End:      CellCoordinates{X: 2, Y: 2},
					Obstacles: []Obstacle{
						{X1: 1, X2: 1, Y1: 0, Y2: 2},
						{X1: 0, X2: 2, Y1: 1, Y2: 1},
					},
				},
			},
			err: nil,
		},
		{
			name:     "invalid test cases input file path",
			filePath: "../test/resource/invalid_path.txt",
			want:     nil,
			err:      errors.New("failed to get file content"),
		},
		{
			name:     "empty test cases input file",
			filePath: "../test/resource/invalid_empty.txt",
			want:     nil,
			err:      errors.New("file is empty"),
		},
		{
			name:     "invalid test cases count (cannot parse)",
			filePath: "../test/resource/invalid_count_1.txt",
			want:     nil,
			err:      errors.New("failed to parse cases count"),
		},
		{
			name:     "invalid test cases count (less than 1)",
			filePath: "../test/resource/invalid_count_2.txt",
			want:     nil,
			err:      errors.New("no test cases provided"),
		},
		{
			name:     "invalid test case grid (cannot parse)",
			filePath: "../test/resource/invalid_grid_1.txt",
			want:     nil,
			err:      errors.New("failed to parse grid rows and columns"),
		},
		{
			name:     "invalid test case grid (less than 1)",
			filePath: "../test/resource/invalid_grid_2.txt",
			want:     nil,
			err:      errors.New("invalid grid size"),
		},
		{
			name:     "invalid test case start or end coordinates (cannot parse)",
			filePath: "../test/resource/invalid_start_or_end.txt",
			want:     nil,
			err:      errors.New("failed to parse start and end coordinates"),
		},
		{
			name:     "invalid test case start",
			filePath: "../test/resource/invalid_start.txt",
			want:     nil,
			err:      errors.New("invalid start coordinates"),
		},
		{
			name:     "invalid test case end",
			filePath: "../test/resource/invalid_end.txt",
			want:     nil,
			err:      errors.New("invalid end coordinates"),
		},
		{
			name:     "invalid test case obstacles count (cannot parse)",
			filePath: "../test/resource/invalid_obstacles_count_1.txt",
			want:     nil,
			err:      errors.New("failed to parse obstacles count"),
		},
		{
			name:     "invalid test case obstacles count (less than 0)",
			filePath: "../test/resource/invalid_obstacles_count_2.txt",
			want:     nil,
			err:      errors.New("invalid obstacles count"),
		},
		{
			name:     "invalid test case obstacle (cannot parse)",
			filePath: "../test/resource/invalid_obstacle_1.txt",
			want:     nil,
			err:      errors.New("failed to parse obstacle"),
		},
		{
			name:     "invalid test case obstacle (invalid coordinates)",
			filePath: "../test/resource/invalid_obstacle_2.txt",
			want:     nil,
			err:      errors.New("invalid obstacle"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ParseTestCases(test.filePath)

			if test.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want, got)
		})
	}
}
