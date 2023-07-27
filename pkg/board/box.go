package board

var boxCenters = []Point{
	{Row: 1, Col: 1},
	{Row: 1, Col: 4},
	{Row: 1, Col: 7},
	{Row: 4, Col: 1},
	{Row: 4, Col: 4},
	{Row: 4, Col: 7},
	{Row: 7, Col: 1},
	{Row: 7, Col: 4},
	{Row: 7, Col: 7},
}

func boxPoints(boxCenter Point) []Point {
	neighbors := []Point{}
	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			neighbors = append(neighbors, Point{
				Row: boxCenter.Row + row,
				Col: boxCenter.Col + col,
			})
		}
	}
	return neighbors
}

func (p *Point) boxNeighbors(space Space) []*Point {
	boxCenter := boxCenters[p.BoxIdx()]
	neighbors := []*Point{}
	for _, boxPoint := range boxPoints(boxCenter) {
		point := space[boxPoint.Row][boxPoint.Col]
		if point == nil {
			continue
		}
		neighbors = append(neighbors, point)
	}
	return neighbors
}

func (p *Point) boxNeighborsExclude(space Space, excludePoint ...*Point) []*Point {
	exclude := map[*Point]struct{}{}
	for _, point := range excludePoint {
		exclude[point] = struct{}{}
	}
	boxCenter := boxCenters[p.BoxIdx()]
	neighbors := []*Point{}
	for _, boxPoint := range boxPoints(boxCenter) {
		point := space[boxPoint.Row][boxPoint.Col]
		if point == nil {
			continue
		}
		if _, ok := exclude[point]; ok {
			continue
		}
		neighbors = append(neighbors, point)
	}
	return neighbors
}
