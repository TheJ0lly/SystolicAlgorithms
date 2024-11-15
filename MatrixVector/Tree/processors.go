package main

import "fmt"

type Processor struct {
	Value      int
	Multiplier int
}

type Tree struct {
	// The processors of the tree.
	P       [][]Processor
	M       *Matrix
	Height  int
	Row     int
	Results []int
}

func CreateTree(size int, M *Matrix) *Tree {
	tr := &Tree{M: M, Row: 0}

	rowsize := size
	row := 0

	for {
		tr.P = append(tr.P, make([]Processor, rowsize))

		if rowsize == 1 {
			row++
			break
		}

		if rowsize%2 == 1 {
			tr.P[row] = append(tr.P[row], Processor{})
			rowsize++
		}

		row++
		rowsize /= 2
	}

	tr.Results = make([]int, M.Height)

	tr.Height = row
	return tr
}

func (tr *Tree) LoadVector(v *Vector) {
	rowsize := len(tr.P[0])
	for i := 0; i < rowsize; i++ {
		tr.P[0][i].Multiplier = v.Data[i]
	}
}

func (tr *Tree) Next() {
	for i := tr.Height - 1; i > 0; i-- {
		rsize := len(tr.P[i])

		// We hold the first processor to be added.
		// We do +1 to find the sibling.
		first := 0
		for j := 0; j < rsize; j++ {
			tr.P[i][j].Value = tr.P[i-1][first].Value + tr.P[i-1][first+1].Value

			// We increment by 2, so we move to the next pair of child processors.
			first += 2
		}
	}

	// We handle the first layer(the leaves of the tree) like this
	// because this layer is the border between the system and user.
	rsize := len(tr.P[0])

	for i := 0; i < rsize; i++ {
		tr.P[0][i].Value = tr.P[0][i].Multiplier * tr.M.Data[tr.Row][i]
	}

	// We advance the column
	tr.Row++

	if tr.P[tr.Height-1][0].Value != 0 {
		tr.Results = append(tr.Results, tr.P[tr.Height-1][0].Value)
	}
}

func (tr *Tree) ShouldStop() bool {
	return tr.Row > tr.M.Height
}

func (tr *Tree) PrintState() {
	// Processor Index - just for printing purposes
	pi := 1
	for i := 0; i < tr.Height; i++ {
		rowsize := len(tr.P[i])
		for j := 0; j < rowsize; j++ {
			fmt.Printf("P%d: %d    ", pi, tr.P[i][j])
			pi++
		}
		fmt.Println()
	}
}

func (tr *Tree) PrintResults() {
	for i := 0; i < len(tr.Results); i++ {
		if tr.Results[i] != 0 {
			fmt.Printf("%d  ", tr.Results[i])
		}
	}
	fmt.Println()
}
