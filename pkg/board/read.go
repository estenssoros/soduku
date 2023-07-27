package board

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

func Read(fileName string) (Board, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return Board{}, errors.Wrap(err, "os.Open")
	}
	defer f.Close()
	cells := [][]int{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		row := []int{}
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		cells = append(cells, row)
	}
	return Board{Cells: cells}, nil
}
