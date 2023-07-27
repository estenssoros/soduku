package board

// solutionLockedCandidates
func solutionLockedCandidates(space [][]*Point) {
	solutionLockedCandidatesBoxRow(space)
	solutionLockedCandidatesBoxCol(space)
	solutionLockedCandidateRow(space)
	solutionLockedCandidateCol(space)
}

// if we can can say a candidate only goes in one row of a box
// then we can eliminate the candidate from the rest of the row
func solutionLockedCandidatesBoxRow(space [][]*Point) {
	for _, boxCenter := range boxCenters {
		// map of candidate to rows it is in
		lookup := map[int][]int{}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				row, col := boxCenter.Row+i, boxCenter.Col+j
				point := space[row][col]
				if point == nil {
					continue
				}
				for _, candidate := range point.Candidates {
					if lookup[candidate] == nil {
						lookup[candidate] = []int{row}
						continue
					}
					if !contains(lookup[candidate], row) {
						lookup[candidate] = append(lookup[candidate], row)
					}
				}
			}
		}
		for candidate, rows := range lookup {
			if len(rows) == 1 {
				// candidate is only in one row of the box
				// so we can eliminate it from the rest of the row
				row := rows[0]
				for col := 0; col < 9; col++ {
					point := space[row][col]
					if point == nil {
						continue
					}
					if boxCenter.BoxIdx() != point.BoxIdx() {
						point.RemoveCandidate(candidate)
					}
				}
			}
		}
	}
}

func solutionLockedCandidatesBoxCol(space [][]*Point) {
	for _, boxCenter := range boxCenters {
		lookup := map[int][]int{}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				row, col := boxCenter.Row+i, boxCenter.Col+j
				point := space[row][col]
				if point == nil {
					continue
				}
				for _, candidate := range point.Candidates {
					if lookup[candidate] == nil {
						lookup[candidate] = []int{col}
						continue
					}
					if !contains(lookup[candidate], col) {
						lookup[candidate] = append(lookup[candidate], col)
					}
				}
			}
		}
		for candidate, cols := range lookup {
			if len(cols) == 1 {
				col := cols[0]
				for row := 0; row < 9; row++ {
					point := space[row][col]
					if point == nil {
						continue
					}
					if boxCenter.BoxIdx() != point.BoxIdx() {
						point.RemoveCandidate(candidate)
					}
				}
			}
		}
	}
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// for each row if the unsolved points are in the same box and have the same candidates,
// then we can remove the from the rest of the box
func solutionLockedCandidateRow(space Space) {
	for _, row := range space {
		lookup := map[int][]*Point{}
		for _, point := range row {
			if point == nil {
				continue
			}
			boxIdx := point.BoxIdx()
			if lookup[boxIdx] == nil {
				lookup[boxIdx] = []*Point{point}
				continue
			}
			lookup[boxIdx] = append(lookup[boxIdx], point)
		}

		for boxIdx, points := range lookup {
			candidates := PointsAllCandidates(points...)
			if len(candidates) != len(points) {
				continue
			}
			for _, point := range boxCenters[boxIdx].boxNeighborsExclude(space, points...) {
				point.RemoveCandidates(candidates...)
			}
		}
	}
}

func solutionLockedCandidateCol(space Space) {
	for col := 0; col < 9; col++ {
		lookup := map[int][]*Point{}
		for _, row := range space {
			point := row[col]
			if point == nil {
				continue
			}
			boxIdx := point.BoxIdx()
			if lookup[boxIdx] == nil {
				lookup[boxIdx] = []*Point{point}
				continue
			}
			lookup[boxIdx] = append(lookup[boxIdx], point)
		}

		for boxIdx, points := range lookup {
			candidates := PointsAllCandidates(points...)
			if len(candidates) != len(points) {
				continue
			}
			for _, point := range boxCenters[boxIdx].boxNeighborsExclude(space, points...) {
				point.RemoveCandidates(candidates...)
			}
		}
	}
}
