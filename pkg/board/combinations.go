package board

func intCombinations(all []int, n int) [][]int {
	result := [][]int{}
	intCombinationsHelper([]int{}, all, n, &result)
	return result
}

func intCombinationsHelper(set []int, all []int, n int, result *[][]int) {
	if n == 0 {
		*result = append(*result, set)
		return
	}
	length := len(all)
	for i := 0; i < length; i++ {
		if length-i < n {
			continue
		}
		intCombinationsHelper(append(set, all[i]), all[i+1:], n-1, result)
	}
}

func pointCombinations(points []*Point, n int) [][]*Point {
	result := [][]*Point{}
	pointCombinationsHelper([]*Point{}, points, n, &result)
	return result
}

func pointCombinationsHelper(set []*Point, all []*Point, n int, result *[][]*Point) {
	if n == 0 {
		*result = append(*result, set)
		return
	}
	length := len(all)
	for i := 0; i < length; i++ {
		if length-i < n {
			continue
		}
		pointCombinationsHelper(append(set, all[i]), all[i+1:], n-1, result)
	}
}

func colorTrapCombinations(colorTraps []ColorTrap, n int) [][]ColorTrap {
	result := [][]ColorTrap{}
	colorTrapCombinationsHelper([]ColorTrap{}, colorTraps, n, &result)
	return result
}

func colorTrapCombinationsHelper(set []ColorTrap, all []ColorTrap, n int, result *[][]ColorTrap) {
	if n == 0 {
		*result = append(*result, set)
		return
	}
	length := len(all)
	for i := 0; i < length; i++ {
		if length-i < n {
			continue
		}
		colorTrapCombinationsHelper(append(set, all[i]), all[i+1:], n-1, result)
	}
}
