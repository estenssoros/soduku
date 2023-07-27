package board

import "fmt"

type Solution struct {
	Row int
	Col int
	Val int
}

func (s Solution) String() string {
	return fmt.Sprintf("row: %d, col: %d, value: %d", s.Row+1, s.Col+1, s.Val)
}

func (s Solution) point() string {
	return fmt.Sprintf("R%dC%d", s.Row+1, s.Col+1)
}

func dedupSolutions(solutions []Solution) []Solution {
	out := []Solution{}
	unique := map[string]struct{}{}
	for _, solution := range solutions {
		p := solution.point()
		if _, ok := unique[p]; !ok {
			out = append(out, solution)
			unique[p] = struct{}{}
		}
	}
	return out
}
