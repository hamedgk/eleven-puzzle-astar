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
mainLoop:
	for _, direction := range possibleMoves {
		copyNode := *node
		copyNode.Direction = direction
		copyNode.Parent = node
		copyNode.Puzzle.MoveBlank(direction)
		copyNode.Cost = node.Cost + 1
		copyNode.HeuristicCost = copyNode.Cost + puzzle.Heuristic(digested, copyNode.Puzzle.Buffer)
		if explored[copyNode.Puzzle.Buffer] {
			for index, frontierNode := range *queue {
				if copyNode.Puzzle.Buffer == node.Puzzle.Buffer {
					if frontierNode.HeuristicCost > copyNode.HeuristicCost {
						fmt.Println("********************************")
						(*queue)[index] = Node{}
						heap.Push(queue, copyNode)
						continue mainLoop
					}
				}
			}
			continue
		}
		heap.Push(queue, copyNode)
		explored[copyNode.Puzzle.Buffer] = true
	}
}

func (node *Node) IsGoal(buffer puzzle.PuzzleBuffer) bool {
	return node.Puzzle.Buffer == buffer
}

func TraceBack(node *Node, counter int) {
	if node.Parent == nil {
		fmt.Println(counter)
		node.Print()
		return
	}

	TraceBack(node.Parent, counter+1)
	fmt.Println(counter)
	node.Print()
}

func (node Node) Print() {
	greenBold := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgHiYellow, color.Bold)
	action := color.New(color.FgHiWhite, color.Italic, color.Bold)
	switch node.Direction {
	case puzzle.Up:
		action.Println("Up")
	case puzzle.Down:
		action.Println("Down")
	case puzzle.Right:
		action.Println("Right")
	case puzzle.Left:
		action.Println("Left")
	}
	w := tabwriter.NewWriter(os.Stdout, 4, 1, 2, ' ', 0)
	for i := 0; i < puzzle.Rows; i++ {
		for j := 0; j < puzzle.Cols; j++ {
			if node.Puzzle.Buffer[i][j] == puzzle.Blank {
				red.Fprintf(w, "%v\t", puzzle.BlankStr)
			} else {
				greenBold.Fprintf(w, "%v\t", node.Puzzle.Buffer[i][j])
			}
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	color.Cyan("----------------------------")
}
