package board

import (
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
)

type Board struct {
	Cells [][]int
}

func New() (b Board) {
	b.Cells = make([][]int, 9)
	for i := range b.Cells {
		b.Cells[i] = make([]int, 9)
	}
	return
}

func (b Board) IsSolved() bool {
	for _, row := range b.Cells {
		for _, val := range row {
			if val != 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) Set(row, col, val int) {
	b.Cells[row][col] = val
}

func (b Board) Get(row, col int) int {
	return b.Cells[row][col]
}

// EmptyPoints finds all points that don't have a value
// and removes candidates from each point based on soduku rules
func (b Board) EmptyPoints() []*Point {
	points := []*Point{}
	for row := 0; row < len(b.Cells); row++ {
		for col := 0; col < len(b.Cells[row]); col++ {
			if b.Get(row, col) == 0 {
				points = append(points, &Point{Row: row, Col: col, Candidates: newCandidates()})
			}
		}
	}
	for _, point := range points {
		point.removeCandidates(b.Cells)
	}
	return points
}

func newCandidates() []int {
	out := make([]int, 9)
	for i := 0; i < 9; i++ {
		out[i] = i + 1
	}
	return out
}

func (b Board) solutionSpace() Space {
	board := make([][]*Point, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]*Point, 9)
	}
	for _, point := range b.EmptyPoints() {
		board[point.Row][point.Col] = point
	}
	return board
}

func hashSpace(space [][]*Point) string {
	hash, err := hashstructure.Hash(space, hashstructure.FormatV2, nil)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d", hash)
}
