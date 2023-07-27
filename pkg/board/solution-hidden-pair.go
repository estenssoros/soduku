package board

// fore each box, if we have two points who's values can only be the same two candidates
// then we can eliminate those candidates from the other points in the box
func solutionHiddenPair(space [][]*Point) {
	hiddenPairBox(space)
	hiddenPairRow(space)
	hiddenPairColumn(space)
}

func hiddenPairBox(space [][]*Point) {
	for _, boxCenter := range boxCenters {
		// map of candidate to points it is in
		lookup := map[int][]*Point{}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				row, col := boxCenter.Row+i, boxCenter.Col+j
				point := space[row][col]
				if point == nil {
					continue
				}
				for _, candidate := range point.Candidates {
					if lookup[candidate] == nil {
						lookup[candidate] = []*Point{point}
						continue
					}
					if !containsPoint(lookup[candidate], point) {
						lookup[candidate] = append(lookup[candidate], point)
					}
				}
			}
		}

		for candidate, points := range lookup {
			if len(points) != 2 {
				delete(lookup, candidate)
			}
		}

		switch len(lookup) {
		case 0, 1:
			continue
		}
		candidates := []int{}
		for k := range lookup {
			candidates = append(candidates, k)
		}

		for _, choice := range intCombinations(candidates, 2) {
			if matchingPair(lookup[choice[0]], lookup[choice[1]]) {
				point1, point2 := lookup[choice[0]][0], lookup[choice[0]][1]
				point1.Candidates = []int{choice[0], choice[1]}
				point2.Candidates = []int{choice[0], choice[1]}

				// remove from rest of box
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						row, col := boxCenter.Row+i, boxCenter.Col+j
						point := space[row][col]
						if point == nil {
							continue
						}
						if point == point1 || point == point2 {
							continue
						}
						point.RemoveCandidate(choice[0])
						point.RemoveCandidate(choice[1])
					}
				}
				// remove from row
				if point1.Row == point2.Row {
					for col := 0; col < 9; col++ {
						point := space[point1.Row][col]
						if point == nil {
							continue
						}
						if point == point1 || point == point2 {
							continue
						}
						point.RemoveCandidate(choice[0])
						point.RemoveCandidate(choice[1])
					}
				}
				// remove from column
				if point1.Col == point2.Col {
					for row := 0; row < 9; row++ {
						point := space[row][point1.Col]
						if point == nil {
							continue
						}
						if point == point1 || point == point2 {
							continue
						}
						point.RemoveCandidate(choice[0])
						point.RemoveCandidate(choice[1])
					}
				}

			}
		}
	}
}

func hiddenPairRow(space Space) {
	for i, row := range space {
		var _ = i
		lookup := map[int][]*Point{}
		for _, point := range row {
			if point == nil {
				continue
			}
			for _, candidate := range point.Candidates {
				lookup[candidate] = append(lookup[candidate], point)
			}
		}
		for candidate, points := range lookup {
			if len(points) != 2 {
				delete(lookup, candidate)
			}
		}
		if len(lookup) != 2 {
			continue
		}
		candidates := []int{}
		for k := range lookup {
			candidates = append(candidates, k)
		}
		if !PointsAreSame(lookup[candidates[0]], lookup[candidates[1]]) {
			continue
		}
		for _, point := range lookup[candidates[0]] {
			point.Candidates = []int{candidates[0], candidates[1]}
		}
		for _, point := range row {
			if point == nil {
				continue
			}
			if point == lookup[candidates[0]][0] || point == lookup[candidates[0]][1] {
				continue
			}
			point.RemoveCandidates(candidates...)
		}
	}
}

func hiddenPairColumn(space Space) {
	for _, column := range space.Columns() {
		lookup := map[int][]*Point{}
		for _, point := range column {
			if point == nil {
				continue
			}
			for _, candidate := range point.Candidates {
				lookup[candidate] = append(lookup[candidate], point)
			}
		}
		for candidate, points := range lookup {
			if len(points) != 2 {
				delete(lookup, candidate)
			}
		}
		if len(lookup) != 2 {
			continue
		}
		candidates := []int{}
		for k := range lookup {
			candidates = append(candidates, k)
		}
		if !PointsAreSame(lookup[candidates[0]], lookup[candidates[1]]) {
			continue
		}
		for _, point := range lookup[candidates[0]] {
			point.Candidates = []int{candidates[0], candidates[1]}
		}
		for _, point := range column {
			if point == nil {
				continue
			}
			if point == lookup[candidates[0]][0] || point == lookup[candidates[0]][1] {
				continue
			}
			point.RemoveCandidates(candidates...)
		}
	}
}

func containsPoint(points []*Point, p *Point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}

func matchingPair(points1, points2 []*Point) bool {
	if len(points1) != 2 || len(points2) != 2 {
		return false
	}
	if points1[0] == points2[0] && points1[1] == points2[1] {
		return true
	}
	if points1[0] == points2[1] && points1[1] == points2[0] {
		return true
	}
	return false
}
