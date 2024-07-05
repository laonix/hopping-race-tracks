package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// TestCase represents a test case with grid, start and end coordinates, and obstacles.
type TestCase struct {
	ID int

	GridRows int
	GridCols int

	Start CellCoordinates
	End   CellCoordinates

	Obstacles []Obstacle
}

// CellCoordinates represents the coordinates of a cell in the grid.
type CellCoordinates struct {
	X int
	Y int
}

// Obstacle represents an area in a grid that is not available for hopping.
type Obstacle struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// ParseTestCases reads the test cases from the specified file and returns them as a slice.
func ParseTestCases(fileName string) ([]*TestCase, error) {
	lines, err := getFileLines(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file content")
	}

	if len(lines) == 0 {
		return nil, errors.New("file is empty")
	}

	casesCount, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse cases count")
	}
	if casesCount == 0 {
		return nil, errors.New("no test cases provided")
	}

	var testCases []*TestCase
	testCaseIdx := 1

	for i := 1; i < len(lines); i++ {
		testCase := &TestCase{
			ID: testCaseIdx,
		}
		testCaseIdx++

		// grid rows and columns
		_, err := fmt.Sscanf(lines[i], "%d %d", &testCase.GridRows, &testCase.GridCols)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("test case %d: failed to parse grid rows and columns", testCase.ID))
		}

		// start and end coordinates
		i++
		_, err = fmt.Sscanf(lines[i], "%d %d %d %d", &testCase.Start.X, &testCase.Start.Y, &testCase.End.X, &testCase.End.Y)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("test case %d: failed to parse start and end coordinates", testCase.ID))
		}

		// obstacles
		i++
		obstaclesCount, err := strconv.Atoi(lines[i])
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("test case %d: failed to parse obstacles count", testCase.ID))
		}

		for j := 0; j < obstaclesCount; j++ {
			i++
			var o Obstacle
			_, err := fmt.Sscanf(lines[i], "%d %d %d %d", &o.X1, &o.X2, &o.Y1, &o.Y2)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("test case %d: failed to parse obstacle %d", testCase.ID, j+1))
			}
			testCase.Obstacles = append(testCase.Obstacles, o)
		}

		testCases = append(testCases, testCase)
	}

	return testCases, nil
}

// getFileLines reads the lines from the specified file and returns them as a slice.
func getFileLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	if scan == nil {
		return nil, errors.New("failed to create scanner")
	}

	scan.Split(bufio.ScanLines)

	var fileLines []string
	for scan.Scan() {
		fileLines = append(fileLines, scan.Text())
	}

	return fileLines, nil
}