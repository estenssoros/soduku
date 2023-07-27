package board

func solutionNakedPair(space Space) {
	solutionNakedPairRow(space)
	solutionNakedPairColumns(space)
}
func solutionNakedPairRow(space Space) {
	for _, row := range space {
		points := []*Point{}
		for _, point := range row {
			if point == nil {
				continue
			}
			if len(point.Candidates) == 2 {
				points = append(points, point)
			}
		}
		if len(points) != 2 {
			continue
		}
		if !PointsSameCandidates(points...) {
			continue
		}
		for _, point := range row {
			if point == nil {
				continue
			}
			if point == points[0] || point == points[1] {
				continue
			}
			point.RemoveCandidates(points[0].Candidates...)
		}
	}
}

func solutionNakedPairColumns(space Space) {
	for _, column := range space.Columns() {

		points := []*Point{}
		for _, point := range column {
			if point == nil {
				continue
			}
			if len(point.Candidates) == 2 {
				points = append(points, point)
			}
		}
		if len(points) != 2 {
			continue
		}
		if !PointsSameCandidates(points...) {
			continue
		}
		for _, point := range column {
			if point == nil {
				continue
			}
			if point == points[0] || point == points[1] {
				continue
			}
			point.RemoveCandidates(points[0].Candidates...)
		}

	}
}
