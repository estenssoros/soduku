package board

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func (b Board) Save(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "os.Create")
	}
	defer f.Close()
	for _, row := range b.Cells {
		for _, cell := range row {
			if _, err := f.WriteString(strconv.Itoa(cell)); err != nil {
				return errors.Wrap(err, "f.Write")
			}
		}
		f.WriteString("\n")
	}
	return nil
}
