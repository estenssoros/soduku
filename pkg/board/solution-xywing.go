package board

// choose pivot and 2 "visible" cells that have the same candidates
// cells that both pivots can see can't have those candidates
// pivot > pincers it can see > sells pincers see
func solutionXyWing(space Space) {
	choices := []*Point{}
	for _, row := range space {
		for _, cell := range row {
			if cell == nil {
				continue
			}
			if len(cell.Candidates) == 2 {
				choices = append(choices, cell)
			}
		}
	}
	combos := pointCombinations(choices, 3)
	combos = combosThatShareNumCandidates(combos, 3)
	pivots, pinchers := xyPivotsAndPincers(combos)
	for i, pivot := range pivots {
		candidates := sharedCandidates(pinchers[i])
		visibleSharedWithoutPivot := space.VisiblePointsSharedExclude(pinchers[i], pivot)
		for _, point := range visibleSharedWithoutPivot {
			for _, candidate := range candidates {
				point.RemoveCandidate(candidate)
			}
		}
		// fmt.Println("pivot", pivot)
		// fmt.Println("\tpincher1", pinchers[i][0])
		// fmt.Println("\tpincher1", pinchers[i][1])
		// fmt.Println("\tcandidates", candidates)
	}
	// os.Exit(1)
}

func combosThatShareNumCandidates(combos [][]*Point, num int) [][]*Point {
	out := [][]*Point{}
	for _, combo := range combos {
		candidateCount := map[int]struct{}{}
		for _, point := range combo {
			for _, candidate := range point.Candidates {
				candidateCount[candidate] = struct{}{}
			}
		}
		if len(candidateCount) == num {
			out = append(out, combo)
		}
	}
	return out
}

func xyPivotsAndPincers(combos [][]*Point) ([]*Point, [][]*Point) {
	pivots, pincers := []*Point{}, [][]*Point{}
	for _, combo := range combos {
		for i := 0; i < 3; i++ {
			pivot := combo[i]
			pincer1, pincer2 := combo[(i+1)%3], combo[(i+2)%3]
			if validXyPivotPincher(pivot, []*Point{pincer1, pincer2}) {
				pivots = append(pivots, pivot)
				pincers = append(pincers, []*Point{pincer1, pincer2})
			}
		}

	}
	return pivots, pincers
}

func validXyPivotPincher(pivot *Point, pinchers []*Point) bool {
	if !pivot.CanSee(*pinchers[0]) || !pivot.CanSee(*pinchers[1]) {
		return false
	}
	if pivot.SharedCandidateCount(*pinchers[0]) != 1 || pivot.SharedCandidateCount(*pinchers[1]) != 1 {
		return false
	}
	if len(PointsSharedCandidates(pinchers...)) == 2 {
		return false
	}

	return true
}

func sharedCandidates(points []*Point) []int {
	out := []int{}
	candidates := map[int]int{}
	for _, point := range points {
		for _, candidate := range point.Candidates {
			candidates[candidate]++
		}
	}
	for candidate, count := range candidates {
		if count == len(points) {
			out = append(out, candidate)
		}
	}
	return out
}
