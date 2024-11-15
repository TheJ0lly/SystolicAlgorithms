package main

import "fmt"

type Processor struct {
	Pass  int
	Value int
}

type ProcessorList struct {
	P      []Processor
	M      *Matrix
	Size   int
	Column int
}

func CreateProcessorList(size int, M *Matrix) *ProcessorList {
	return &ProcessorList{P: make([]Processor, size), Size: size, M: M, Column: 0}
}

func (pl *ProcessorList) Next(val int) {
	// For each processor, we get the value of the processor before it,
	// and perform the computation.
	for i := pl.Size - 1; i > 0; i-- {
		pl.P[i].Pass = pl.P[i-1].Pass
		pl.P[i].Value += pl.P[i].Pass * pl.M.Data[i][pl.Column]
	}

	// The first processor gets the input value.
	pl.P[0].Pass = val

	// We also compute the first processor value, because we treat it specially
	// since its the line between the sytem and the user.
	pl.P[0].Value += pl.P[0].Pass * pl.M.Data[0][pl.Column]

	// We advance the Matrix column index
	pl.Column++
}

// If the index of the Column is the same, or bigger than the width of the matrix
// we can stop.
func (pl *ProcessorList) ShouldStop() bool {
	return pl.Column >= pl.M.Width
}

func (pl *ProcessorList) PrintState() {
	for i := 0; i < pl.Size; i++ {
		fmt.Printf("P%d: Value: %d - Pass: %d\n", i, pl.P[i].Value, pl.P[i].Pass)
	}
}
