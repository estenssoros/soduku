package board

func solutionXyzWing(space Space) {
	choices := space.Flatten()
	combos := pointCombinations(choices, 3)
	combos = combosThatShareNumCandidates(combos, 3)
	pivots, pinchers := xyzPivotsAndPinchers(combos)

	for i, pivot := range pivots {
		shared := space.VisiblePointsShared(pivot, pinchers[i][0], pinchers[i][1])
		if len(shared) == 0 {
			continue
		}
		// fmt.Println("pivot", pivot)
		// fmt.Println("\tpincher1", pinchers[i][0])
		// fmt.Println("\tpincher1", pinchers[i][1])
		// fmt.Println("\tshared", shared)
		// os.Exit(1)

		zCandidates := PointsSharedCandidates(pivot, pinchers[i][0], pinchers[i][1])

		for _, point := range shared {
			for _, z := range zCandidates {
				point.RemoveCandidate(z)
			}
		}
	}
}

func xyzPivotsAndPinchers(combos [][]*Point) ([]*Point, [][]*Point) {
	pivots, pincers := []*Point{}, [][]*Point{}
	for _, combo := range combos {
		for i := 0; i < 3; i++ {
			pivot := combo[i]
			if len(pivot.Candidates) != 3 {
				continue
			}
			pincer1, pincer2 := combo[(i+1)%3], combo[(i+2)%3]
			if validXyzPivotPincher(pivot, []*Point{pincer1, pincer2}) {
				pivots = append(pivots, pivot)
				pincers = append(pincers, []*Point{pincer1, pincer2})
			}
		}
	}
	return pivots, pincers
}

func validXyzPivotPincher(pivot *Point, pinchers []*Point) bool {
	if !pivot.CanSee(*pinchers[0]) || !pivot.CanSee(*pinchers[1]) {
		return false
	}
	if pivot.SharedCandidateCount(*pinchers[0]) != 2 || pivot.SharedCandidateCount(*pinchers[1]) != 2 {
		return false
	}
	if len(PointsSharedCandidates(pinchers...)) == 2 {
		return false
	}
	return true
}
