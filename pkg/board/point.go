package board

import (
	"fmt"
	"sort"
)

type Point struct {
	Row        int
	Col        int
	Candidates []int
}

func (p Point) Equals(other Point) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (p Point) HasCandidate(val int) bool {
	for _, v := range p.Candidates {
		if v == val {
			return true
		}
	}
	return false
}

func (p *Point) RemoveCandidate(val int) {
	for i := 0; i < len(p.Candidates); i++ {
		if p.Candidates[i] == val {
			p.Candidates = append(p.Candidates[:i], p.Candidates[i+1:]...)
			return
		}
	}
}
func (p *Point) RemoveCandidates(vals ...int) {
	for _, val := range vals {
		p.RemoveCandidate(val)
	}
}

func (p Point) String() string {
	return fmt.Sprintf("R%dC%d candidates: %v", p.Row+1, p.Col+1, p.Candidates)
}

func (p *Point) removeCandidates(board [][]int) {
	p.removeCandidatesVertically(board)
	p.removeCandidatesHorizontally(board)
	p.removeCandidatesBox(board)
}

func (p *Point) removeCandidatesVertically(board [][]int) {
	for row := 0; row < 9; row++ {
		p.RemoveCandidate(board[row][p.Col])
	}
}

func (p *Point) removeCandidatesHorizontally(board [][]int) {
	for col := 0; col < 9; col++ {
		p.RemoveCandidate(board[p.Row][col])
	}
}

func (p *Point) removeCandidatesBox(board [][]int) {
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			p.RemoveCandidate(board[p.Row-p.Row%3+k][p.Col-p.Col%3+l])
		}
	}
}

func (p Point) BoxIdx() int {
	return p.Row/3*3 + p.Col/3
}

func (p Point) CanSee(other Point) bool {
	return p.Row == other.Row || p.Col == other.Col || p.BoxIdx() == other.BoxIdx()
}

func (p *Point) VisiblePointsExclude(space [][]*Point, exclude []*Point) []*Point {
	lookup := map[*Point]struct{}{
		p: {},
	}
	for _, point := range exclude {
		lookup[point] = struct{}{}
	}
	for _, row := range space {
		for _, point := range row {
			if point == nil {
				continue
			}
			if _, ok := lookup[point]; ok {
				continue
			}
			if p.CanSee(*point) {
				lookup[point] = struct{}{}
			}
		}
	}

	return nil
}

func (p Point) SharedCandidateCount(other Point) int {
	candidates := map[int]int{}
	for _, candidate := range p.Candidates {
		candidates[candidate]++
	}
	for _, candidate := range other.Candidates {
		candidates[candidate]++
	}
	var candidateCount int
	for _, count := range candidates {
		if count == 2 {
			candidateCount++
		}
	}
	return candidateCount
}

func PointsSameCandidates(points ...*Point) bool {
	candidates := map[int]int{}
	for _, point := range points {
		for _, candidate := range point.Candidates {
			candidates[candidate]++
		}
	}
	for _, count := range candidates {
		if count != len(points) {
			return false
		}
	}
	return true
}

func PointsSharedCandidates(points ...*Point) []int {
	candidates := map[int]int{}
	for _, point := range points {
		for _, candidate := range point.Candidates {
			candidates[candidate]++
		}
	}
	var sharedCandidates []int
	for candidate, count := range candidates {
		if count == len(points) {
			sharedCandidates = append(sharedCandidates, candidate)
		}
	}
	sort.Ints(sharedCandidates)
	return sharedCandidates
}

func PointsAllCandidates(points ...*Point) []int {
	candidates := map[int]struct{}{}
	for _, point := range points {
		for _, candidate := range point.Candidates {
			candidates[candidate] = struct{}{}
		}
	}
	var allCandidates []int
	for candidate := range candidates {
		allCandidates = append(allCandidates, candidate)
	}
	sort.Ints(allCandidates)
	return allCandidates
}

func PointsAreSame(points1, points2 []*Point) bool {
	if len(points1) != len(points2) {
		return false
	}
	PointsSort(points1)
	PointsSort(points2)
	for i := range points1 {
		if !points1[i].Equals(*points2[i]) {
			return false
		}
	}
	return true
}

func PointsSort(points []*Point) {
	sort.Slice(points, func(i, j int) bool {
		if points[i].Row == points[j].Row {
			return points[i].Col < points[j].Col
		}
		return points[i].Row < points[j].Row
	})
}
