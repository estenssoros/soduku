package board

import (
	"github.com/pkg/errors"
)

var ErrNoSolutionsFound = errors.New("no solution found")

var (
	SolutionTypeNakedSingle        = "naked single"
	SolutionTypeHiddenBoxSingle    = "hidden box single"
	SolutionTypeHiddenRowSingle    = "hidden row single"
	SolutionTypeHiddenColumnSingle = "hidden column single"
	SolutionTypeLockedCandidate    = "locked candidate"
	SolutionTypeHiddenPair         = "hidden pair"
)

func Heuristics(b Board) (solutions []Solution, err error) {

	space := b.solutionSpace()
	defer func() {
		if err != nil {
			space.Export()
		}
	}()
	// space.Export()

	var hash string
	newHash := hashSpace(space)
	for hash != newHash {
		if solutions := solutionNakedSingles(space); len(solutions) > 0 {
			return dedupSolutions(solutions), nil
		}
		if solutions := solutionHiddenSingles(space); len(solutions) > 0 {
			return dedupSolutions(solutions), nil
		}

		solutionLockedCandidates(space)
		solutionHiddenPair(space)
		solutionNakedPair(space)
		solutionXyWing(space)
		solutionXyzWing(space)
		solutionColorTrap(space)

		// space.Export()
		hash = newHash
		newHash = hashSpace(space)
	}

	return nil, ErrNoSolutionsFound
}
