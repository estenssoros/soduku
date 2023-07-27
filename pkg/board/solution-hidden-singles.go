package board

func solutionHiddenSingles(space [][]*Point) []Solution {
	solutions := hiddenSolutionsVertical(space)
	solutions = append(solutions, hiddenSolutionsHorizontal(space)...)
	solutions = append(solutions, hiddenSolutionsBox(space)...)
	return dedupSolutions(solutions)
}

func hiddenSolutionsVertical(space [][]*Point) []Solution {
	solutions := []Solution{}
	for i := 0; i < 9; i++ {
		val := i + 1
		for col := 0; col < 9; col++ {
			values := []*Point{}
			for row := 0; row < 9; row++ {
				point := space[row][col]
				if point == nil {
					continue
				}
				if point.HasCandidate(val) {
					values = append(values, point)
				}
			}
			if len(values) == 1 {
				solutions = append(solutions, Solution{
					Row: values[0].Row,
					Col: values[0].Col,
					Val: val,
				})
			}
		}
	}
	return solutions
}

func hiddenSolutionsHorizontal(space [][]*Point) []Solution {
	solutions := []Solution{}
	for i := 0; i < 9; i++ {
		val := i + 1
		for row := 0; row < 9; row++ {
			values := []*Point{}
			for col := 0; col < 9; col++ {
				point := space[row][col]
				if point == nil {
					continue
				}
				if point.HasCandidate(val) {
					values = append(values, point)
				}
			}
			if len(values) == 1 {
				solutions = append(solutions, Solution{
					Row: values[0].Row,
					Col: values[0].Col,
					Val: val,
				})
			}
		}
	}
	return solutions
}

func hiddenSolutionsBox(space [][]*Point) []Solution {
	solutions := []Solution{}
	for _, boxCenter := range boxCenters {
		for i := 0; i < 9; i++ {
			val := i + 1
			values := []*Point{}
			for _, boxPoint := range boxPoints(boxCenter) {
				point := space[boxPoint.Row][boxPoint.Col]
				if point == nil {
					continue
				}
				if point.HasCandidate(val) {
					values = append(values, point)
				}
			}
			if len(values) == 1 {
				solutions = append(solutions, Solution{
					Row: values[0].Row,
					Col: values[0].Col,
					Val: val,
				})
			}
		}
	}
	return solutions
}
