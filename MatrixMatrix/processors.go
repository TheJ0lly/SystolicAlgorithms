package main

import "fmt"

type Processor struct {
	Value int
	PassA int
	PassB int
}

type ProcMatrix struct {
	P      [][]Processor
	Height int
	Width  int

	// The left matrix
	Left *Matrix

	// The top matrix
	Top *Matrix

	// For the M1 Matrix - it comes from the left
	LeftColumn int

	// It still represents the column, but since the M2 is transposed, its more readable - it comes from the top.
	TopRow int
}

func CreateProcMatrix(h, w int, Left, Top *Matrix) *ProcMatrix {
	toret := &ProcMatrix{P: make([][]Processor, h), Height: h, Width: w, Left: Left, Top: Top, LeftColumn: 0, TopRow: 0}

	for i := 0; i < toret.Height; i++ {
		toret.P[i] = make([]Processor, toret.Width)
	}

	return toret
}

func (pm *ProcMatrix) Next() {

	// Handle the interior of the matrix, without the top and left sides.
	for r := pm.Height - 1; r > 0; r-- {
		for c := pm.Width - 1; c > 0; c-- {
			pm.P[r][c].PassA = pm.P[r][c-1].PassA
			pm.P[r][c].PassB = pm.P[r-1][c].PassB
			pm.P[r][c].Value += pm.P[r][c].PassA * pm.P[r][c].PassB
		}
	}

	// Handle left side minus the top left processor
	for r := pm.Height - 1; r > 0; r-- {
		if pm.LeftColumn < pm.Left.Width {
			pm.P[r][0].PassA = pm.Left.Data[r][pm.LeftColumn]
		} else {
			pm.P[r][0].PassA = 0
		}
		pm.P[r][0].PassB = pm.P[r-1][0].PassB
		pm.P[r][0].Value += pm.P[r][0].PassA * pm.P[r][0].PassB
	}

	// Handle top side minus the top left processor
	for c := pm.Width - 1; c > 0; c-- {
		if pm.TopRow < pm.Top.Width {
			pm.P[0][c].PassB = pm.Top.Data[c][pm.TopRow]
		} else {
			pm.P[0][c].PassB = 0
		}
		pm.P[0][c].PassA = pm.P[0][c-1].PassA
		pm.P[0][c].Value += pm.P[0][c].PassA * pm.P[0][c].PassB
	}

	// Handle the top left
	if pm.LeftColumn < pm.Left.Width {
		pm.P[0][0].PassA = pm.Left.Data[0][pm.LeftColumn]
	} else {
		pm.P[0][0].PassA = 0
	}

	if pm.TopRow < pm.Top.Height {
		pm.P[0][0].PassB = pm.Top.Data[0][pm.TopRow]
	} else {
		pm.P[0][0].PassB = 0
	}
	pm.P[0][0].Value += pm.P[0][0].PassA * pm.P[0][0].PassB

	// We advance the states
	pm.TopRow++
	pm.LeftColumn++
}

func (pm *ProcMatrix) ShouldStop() bool {
	return pm.TopRow > pm.Left.Height+1 && pm.LeftColumn > pm.Top.Width+1
}

func (pm *ProcMatrix) PrintState() {
	pi := 0

	for i := 0; i < pm.Height; i++ {
		for j := 0; j < pm.Width; j++ {
			fmt.Printf("P%d: V:%d A:%d B:%d    ", pi, pm.P[i][j].Value, pm.P[i][j].PassA, pm.P[i][j].PassB)
			pi++
		}
		fmt.Println()
	}
}

func (pm *ProcMatrix) PrintResult() {
	for i := 0; i < pm.Height; i++ {
		for j := 0; j < pm.Width; j++ {
			fmt.Printf("%d ", pm.P[i][j].Value)
		}
		fmt.Println()
	}
}
