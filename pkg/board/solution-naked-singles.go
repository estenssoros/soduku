package board

func solutionNakedSingles(space [][]*Point) []Solution {
	solutions := []Solution{}
	for _, row := range space {
		for _, point := range row {
			if point == nil {
				continue
			}
			if len(point.Candidates) == 1 {
				solutions = append(solutions, Solution{
					Row: point.Row,
					Col: point.Col,
					Val: point.Candidates[0],
				})
			}
		}
	}

	return solutions

}
