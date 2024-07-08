# Hopping Race Tracks

Here you can find a variation of _Race Tracks_ game.

**Hopping Race Tracks** is played on a rectangular grid, where each
square on the grid is either empty or occupied.
While hoppers can fly over any square, they can only
land on empty squares.
At any point in time, a hopper has a velocity `(x,y)`, where `x` and `y` are the
speed (in squares) parallel to the grid.
Thus, a speed of `(2,1)` corresponds to a knight jump, as does `(-
2,1)` and six other speeds.

To determine the hops a hopper can make, we need to know how much speed he can pick up or lose:
either `-1`, `0`, or `1` square in both directions.
Thus, while having speed `(2,1)`, the hopper can change it to `(1,0)`, `(1,1)`, `(1,2)`, `(2,0)`, `(2,1)`, `(2,2)`, `(3,0)`, `(3,1)` and `(3,2)`.
It is impossible for the hopper to
get a velocity of `4` in either direction, so the `x` and `y` components will stay between `-3` and `3`
inclusive.

The goal of **Hopping Race Tracks** is to get from start to finish as quickly as possible (i.e., in the least
number of hops) without landing on occupied squares.

## Input

The solution should accept a text file as input.

Please be rational and consistent with the input file format
as its variations will cause the solution to fail immediately in the attempt to parse it.

### First line: the number of test cases
The first line contains the number of test cases (`N`) the solution has to process.

### Test Case

| Line               | Content                                                                                                                                                                                                                                                                                                                                                                                  | Example   |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------|
| 1                  | The _width_ `X` (`1 ≤ X ≤ 30`) and _height_ `Y` (`1 ≤ Y ≤ 30`) _of the grid_. <br/> `X` and `Y` values must be positive integers separated by a _single_ whitespace.                                                                                                                                                                                                                     | `5 5`     |
| 2                  | The _start_ and the _end_ _position_ of the hopper. <br/> This line contains _four_ positive integers separated by a _single_ whitespace. <br/> The first two numbers `(x1, y1)` indicate the start point (`0 ≤ x1 < X`, `0 ≤ y1 < Y`). <br/> The second two numbers `(x2, y2)` indicate the end point (`0 ≤ x2 < X`, `0 ≤ y2 < Y`).                                                     | `4 0 4 4` |
| 3                  | The _number of obstacles_ `P` in the grid.                                                                                                                                                                                                                                                                                                                                               | `1`       |
| 4 to (`4 + P - 1`) | _Obstacle_ specification. <br/> Each line contains _four_ positive integers separated by a _single_ whitespace: `x1`, `x2`, `y1`, and `y2` (in this exact order). <br/> This numbers indicate that all squares `(x,y)` with `x1 ≤ x ≤ x2` and `y1 ≤ y ≤ y2` are occupied. <br/> The start point will never be occupied. <br/> The limitations are: `0 ≤ x1 ≤ x2 < X`, `0 ≤ y1 ≤ y2 < Y`. | `1 4 2 3` |

### Example Input File Content

```
2
5 5
4 0 4 4
1
1 4 2 3
3 3
0 0 2 2
2
1 1 0 2
0 2 1 1
```

This input file contains `2` test cases:
1. A `5x5` grid with a hopper starting at `(4,0)` and ending at `(4,4)`. The following squares are occupied: `(1,2)`, `(2,2)`, `(3,2)`, `(4,2)`, `(1,3)`, `(2,3)`, `(3,3)`, and `(4,3)`.
2. A `3x3` grid with a hopper starting at `(0,0)` and ending at `(2,2)`. The following squares are occupied: `(1,0)`, `(1,0)`, `(1,1)`, `(2,1)`, and `(1,2)`.

## Output

The solution should output the minimum number of hops the hopper needs to reach the end position.
If the hopper cannot reach the end position, the output be `No solution`.

### Example Output

The result for the example input file above should be:

```
Test case #1: Optimal solution takes 7 hops.
Test case #2: No solution.
```

## Implementation Details

The solution implementation is based on [A* algorithm](https://theory.stanford.edu/~amitp/GameProgramming/AStarComparison.html).

In the case of Hopping Race Tracks, the algorithm uses a priority queue to determine the next best hopper position to explore.

The value of each square's `GCost` is the _number of hops_ from the start position to the current square. Thus, if the solution exists, the end point's `GCost` will hold a minimal number of hops we're looking for.

The `HCost` is the `Diagonal distance` from the current square to the end position, because the hopper is able to move not only in the cardinal directions but also diagonally.
In the case of Hopping Race Tracks, the cost of straightforward and diagonal movements are equal, so we can utilize `Chebyshev Distance` formula for calculating `HCost`.

The `FCost` is the sum of `GCost` and `HCost`. It is the main basis for the priority queue to determine the next best square to explore. If the `FCost` is equal for two squares, the square with the lower `HCost` is chosen.