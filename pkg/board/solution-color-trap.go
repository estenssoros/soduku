package board

import (
	"fmt"
	"os"
)

func solutionColorTrap(space Space) {
	for candidate := 1; candidate <= 9; candidate++ {
		colorTraps := getColorTrapsRow(space, candidate)
		colorTraps = append(colorTraps, getColorTrapsColumn(space, candidate)...)
		colorTraps = append(colorTraps, getColorTrapsBox(space, candidate)...)

		combos := colorTrapCombinations(colorTraps, 2)
		if candidate == 3 {
			// for _, ct := range colorTraps {
			// 	fmt.Println(ct)
			// }
			fmt.Println("combos", len(combos))
			os.Exit(1)
		}

	}
	// get color traps for row
	// get color traps for columns
	// get color traps for boxes
}

type ColorTrap struct {
	Candidate int
	Points    []*Point
}

func getColorTrapsRow(space Space, candidate int) []ColorTrap {
	out := []ColorTrap{}
	for _, row := range space {
		points := []*Point{}
		for _, point := range row {
			if point == nil {
				continue
			}
			if point.HasCandidate(candidate) {
				points = append(points, point)
			}
		}
		if len(points) != 2 {
			continue
		}
		out = append(out, ColorTrap{candidate, points})
	}
	return out
}

func getColorTrapsColumn(space Space, candidate int) []ColorTrap {
	out := []ColorTrap{}
	for _, column := range space.Columns() {
		points := []*Point{}
		for _, point := range column {
			if point == nil {
				continue
			}
			if point.HasCandidate(candidate) {
				points = append(points, point)
			}
		}
		if len(points) != 2 {
			continue
		}
		out = append(out, ColorTrap{candidate, points})
	}
	return out
}

func getColorTrapsBox(space Space, candidate int) []ColorTrap {
	out := []ColorTrap{}
	for _, box := range space.Boxes() {
		points := []*Point{}
		for _, point := range box {
			if point == nil {
				continue
			}
			if point.HasCandidate(candidate) {
				points = append(points, point)
			}
		}
		if len(points) != 2 {
			continue
		}
		out = append(out, ColorTrap{candidate, points})
	}
	return out
}
