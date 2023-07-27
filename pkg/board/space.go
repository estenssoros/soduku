package board

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Space [][]*Point

func (s Space) VisiblePointsShared(points ...*Point) []*Point {
	exclude := map[*Point]struct{}{}
	for _, point := range points {
		exclude[point] = struct{}{}
	}
	pointCount := map[*Point]int{}
	for _, point := range points {
		for _, visible := range s.VisiblePoints(point) {
			if _, ok := exclude[visible]; !ok {
				pointCount[visible]++
			}
		}
	}
	out := []*Point{}
	for point, count := range pointCount {
		if count == len(points) {
			out = append(out, point)
		}
	}

	return out
}

func (s Space) VisiblePointsSharedExclude(points []*Point, excludePoints ...*Point) []*Point {
	exclude := map[*Point]struct{}{}
	for _, point := range points {
		exclude[point] = struct{}{}
	}
	for _, point := range excludePoints {
		exclude[point] = struct{}{}
	}

	pointCount := map[*Point]int{}
	for _, point := range points {
		for _, visible := range s.VisiblePoints(point) {
			if _, ok := exclude[visible]; !ok {
				pointCount[visible]++
			}
		}
	}
	out := []*Point{}
	for point, count := range pointCount {
		if count == len(points) {
			out = append(out, point)
		}
	}
	return out
}

func (s Space) VisiblePoints(p *Point) []*Point {
	points := []*Point{}
	for _, row := range s {
		for _, point := range row {
			if point == nil {
				continue
			}
			if p.CanSee(*point) && p != point {
				points = append(points, point)
			}
		}
	}

	return points
}

func (s Space) Flatten() []*Point {
	points := []*Point{}
	for _, row := range s {
		for _, point := range row {
			if point == nil {
				continue
			}
			points = append(points, point)
		}
	}
	return points
}
func (s Space) String() string {
	var b strings.Builder
	b.WriteString("|" + strings.Repeat("-", 9*10-1) + "|\n")
	for _, row := range s {
		b.WriteString("|")
		for _, cell := range row {
			if cell == nil {
				b.WriteString(strings.Repeat(" ", 9) + "|")
				continue
			}
			candidates := joinIntsString(cell.Candidates)
			remainder := 9 - len(candidates)
			b.WriteString(candidates + strings.Repeat(" ", remainder) + "|")
		}
		b.WriteString("\n")

		for i := 0; i < 2; i++ {
			b.WriteString("|")
			for cell := 0; cell < 9; cell++ {
				b.WriteString(strings.Repeat(" ", 9) + "|")
			}
			b.WriteString("\n")
		}

		b.WriteString("|" + strings.Repeat("-", 9*10-1) + "|\n")
	}
	return b.String()
}

func (s Space) Export() error {
	f, err := os.Create("space.txt")
	if err != nil {
		return errors.Wrap(err, "os.Create")
	}
	defer f.Close()
	_, err = f.WriteString(s.String())
	return err
}

func joinIntsString(ints []int) string {
	out := []string{}
	for _, i := range ints {
		out = append(out, strconv.Itoa(i))
	}
	return strings.Join(out, "")
}

func unsolvedRowCount(row []*Point) int {
	count := 0
	for _, point := range row {
		if point != nil {
			count++
		}
	}
	return count
}

func (s Space) Columns() [][]*Point {
	columns := [][]*Point{}
	for i := 0; i < 9; i++ {
		column := make([]*Point, 9)
		for _, row := range s {
			column = append(column, row[i])
		}
		columns = append(columns, column)
	}
	return columns
}

func (s Space) Boxes() [][]*Point {
	boxes := [][]*Point{}
	for _, boxCenter := range boxCenters {
		box := []*Point{}
		for _, point := range boxPoints(boxCenter) {
			box = append(box, s[point.Row][point.Col])
		}
		boxes = append(boxes, box)
	}
	return boxes
}
