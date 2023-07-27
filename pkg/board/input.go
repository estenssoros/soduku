package board

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func NewFromInput() (Board, error) {
	board := New()
	var row, col int
	for {

		board.DisplayAtCell(row, col)
		val, err := getVal()
		if err != nil {
			return Board{}, errors.Wrap(err, "getVal()")
		}
		switch val {
		case 'q':
			return board, nil
		case 'n', '\n':
			row, col = nextRowCol(row, col)
			continue
		case 'w':
			if row > 0 {
				row--
			} else {
				row = 9
			}
			continue
		case 'a':
			if col > 0 {
				col--
			} else {
				col = 8
			}
			continue
		case 's':
			if row < 8 {
				row++
			} else {
				row = 0
			}
			continue
		case 'd':
			if col < 8 {
				col++
			} else {
				col = 0
			}
			continue
		}
		board.Set(row, col, int(val-'0'))
		row, col = nextRowCol(row, col)
	}
}

func nextRowCol(row, col int) (int, int) {
	if col < 8 {
		col++
	} else {
		col = 0
		row++
	}
	if row == 9 {
		row = 0
	}
	return row, col
}

func getVal() (byte, error) {
	fmt.Print("Enter a value (1-9, q to quit, n to skip): ")
	var first string
	_, err := fmt.Scanln(&first)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected newline") {
			return 'n', nil
		}
		return 'x', errors.Wrap(err, "fmt.Scanln")
	}
	val := first[0]
	if val < '1' || val > '9' {
		switch val {
		case 'q', 'n', 'w', 'a', 's', 'd':
			return val, nil
		}
		return 'x', errors.New("invalid value")
	}
	return val, nil
}
