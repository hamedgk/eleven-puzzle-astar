package data_structures

import (
	"container/heap"
	"eleven-puzzle/data_structures/puzzle"
	"fmt"
	"os"

	"text/tabwriter"

	"github.com/fatih/color"
)

type Node struct {
	Parent        *Node
	Direction     puzzle.Direction
	Puzzle        puzzle.Puzzle
	Cost          int
	HeuristicCost int
}

func (node *Node) Expand(queue *PriorityQueue, explored map[puzzle.PuzzleBuffer]bool, digested puzzle.DigestCoords) {
	possibleMoves := node.Puzzle.PossibleBlankMoves()
	for _, direction := range possibleMoves {
		copyNode := *node
		copyNode.Direction = direction
		copyNode.Parent = node
		copyNode.Puzzle.MoveBlank(direction)
		copyNode.Cost = node.Cost + 1
		copyNode.HeuristicCost = copyNode.Cost + puzzle.Heuristic(digested, copyNode.Puzzle.Buffer)
		if explored[copyNode.Puzzle.Buffer] {
			continue
		}
		heap.Push(queue, copyNode)
		explored[copyNode.Puzzle.Buffer] = true
	}
}

func (node *Node) IsGoal(buffer puzzle.PuzzleBuffer) bool {
	return node.Puzzle.Buffer == buffer
}

func TraceBack(node Node) {
	if node.Parent == nil {
		node.Print()
		return
	}

	TraceBack(*node.Parent)
	node.Print()

}

func (node Node) Print() {
	puzzleColor := color.New(color.FgGreen, color.Bold)
	switch node.Direction {
	case puzzle.Up:
		color.Red("Up")
	case puzzle.Down:
		color.Red("Down")
	case puzzle.Right:
		color.Red("Right")
	case puzzle.Left:
		color.Red("Left")
	}
	w := tabwriter.NewWriter(os.Stdout, 4, 1, 2, ' ', 0)
	for i := 0; i < puzzle.Rows; i++ {
		for j := 0; j < puzzle.Cols; j++ {
			puzzleColor.Fprintf(w, "%v\t", node.Puzzle.Buffer[i][j])
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	color.Cyan("----------------------------")
}
