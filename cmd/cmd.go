package cmd

import (
	"fmt"

	"github.com/estenssoros/soduku/pkg/board"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(
		// add commands here
		newCmd,
		loadCmd,
	)
}

var cmd = &cobra.Command{
	Use:   "soduku",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := board.NewFromInput()
		if err != nil {
			return errors.Wrap(err, "board.NewFromInput")
		}
		b.Display()
		if err := b.Save("board.txt"); err != nil {
			return errors.Wrap(err, "b.Save")
		}
		if err := b.Save("board2.txt"); err != nil {
			return errors.Wrap(err, "b.Save")
		}
		return nil
	},
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("please supply one file arg")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := board.Read(args[0])
		if err != nil {
			return errors.Wrap(err, "board.Read")
		}
		for {
			solutions, err := board.Heuristics(b)
			if err != nil {
				b.Display()
				return errors.Wrap(err, "solve.Heuristics")
			}
			for i, solution := range solutions {
				b.DisplayWithSolution(solution)
				fmt.Printf("%d/%d - %s\n", i+1, len(solutions), solution)
				fmt.Printf("next? ")
				fmt.Scanln()
				b.Set(solution.Row, solution.Col, solution.Val)
				if err := b.Save("board.txt"); err != nil {
					return errors.Wrap(err, "b.Save")
				}
			}
		}
		return nil
	},
}
