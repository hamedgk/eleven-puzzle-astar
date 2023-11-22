package puzzle

import (
	"math"
)

func Heuristic(coords DigestCoords, buffer PuzzleBuffer) int {
	var sum int

	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			number := buffer[i][j]
			if number == Blank {
				continue
			}
			digested := coords[number]
			min := MaxInt
			for _, value := range digested {
				xDiff := math.Abs(float64(i - value.x))
				yDiff := math.Abs(float64(j - value.y))
				manhatan := int(xDiff) + int(yDiff)
				if manhatan < min {
					min = int(manhatan)
				}
			}
			sum += min
		}
	}

	return sum
}
