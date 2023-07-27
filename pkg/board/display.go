package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/estenssoros/soduku/pkg/utils"
	"github.com/fatih/color"
)

func (b Board) Display() {
	display(b.Cells)
}

func display(board [][]int) {
	utils.ClearScreen()
	b := strings.Builder{}
	for i, row := range board {

		if i%3 == 0 {
			b.WriteString(" ---------------------\n")
		}

		for j, cell := range row {
			if j%3 == 0 {
				b.WriteString("|")
			}
			b.WriteString(" ")
			if cell == 0 {
				b.WriteString(" ")
				continue
			}
			b.WriteString(strconv.Itoa(cell))
		}
		b.WriteString(" |\n")
	}
	b.WriteString(" ---------------------\n")
	fmt.Println(b.String())
}

func (b Board) DisplayWithSolution(solution Solution) {
	displayWithSolution(b.Cells, solution)
}

func displayWithSolution(board [][]int, solution Solution) {
	utils.ClearScreen()
	b := strings.Builder{}
	for i, row := range board {

		if i%3 == 0 {
			b.WriteString(" ---------------------\n")
		}

		for j, cell := range row {
			if j%3 == 0 {
				b.WriteString("|")
			}
			b.WriteString(" ")
			if i == solution.Row && j == solution.Col {
				green := color.New(color.FgGreen)
				green.Fprint(&b, strconv.Itoa(solution.Val))
				continue
			}
			if cell == 0 {
				b.WriteString(" ")
				continue
			}
			b.WriteString(strconv.Itoa(cell))
		}
		b.WriteString(" |\n")
	}
	b.WriteString(" ---------------------\n")
	fmt.Println(b.String())
}

func (b Board) DisplayAtCell(row, col int) {
	displayAtCell(b.Cells, row, col)
}

func displayAtCell(board [][]int, rowIdx, colIdx int) {
	utils.ClearScreen()

	b := strings.Builder{}
	for i, row := range board {

		if i%3 == 0 {
			b.WriteString(" ---------------------\n")
		}

		for j, cell := range row {
			if j%3 == 0 {
				b.WriteString("|")
			}
			b.WriteString(" ")
			if i == rowIdx && j == colIdx {
				green := color.New(color.FgGreen)
				if cell == 0 {
					green.Fprint(&b, "X")
				} else {
					green.Fprint(&b, strconv.Itoa(cell))
				}
				continue
			}
			if cell == 0 {
				b.WriteString(" ")
				continue
			}
			b.WriteString(strconv.Itoa(cell))
		}
		b.WriteString(" |\n")
	}
	b.WriteString(" ---------------------\n")
	fmt.Println(b.String())
}
