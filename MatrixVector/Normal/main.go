package main

import (
	"flag"
	"fmt"
)

func main() {
	matrixFile := flag.String("m", "", "the json file containing the matrix")
	vectorFile := flag.String("v", "", "the json file containing the vector")
	allStates := flag.Bool("a", false, "print the state after each step")

	flag.Parse()

	m, err := GetMatrixFromFile(*matrixFile)

	if err != nil {
		fmt.Printf("matrix file error: %s\n", err)
		return
	}

	m.Adjust()

	v, err := GetVectorFromFile(*vectorFile)

	if err != nil {
		fmt.Printf("vector file error: %s\n", err)
		return
	}

	// We pass the matrix height as the number of processors because we need a proc/row.
	p := CreateProcessorList(m.Height, m)

	// We process the vector.
	for i := 0; i < v.Size; i++ {
		p.Next(v.Data[i])
		if *allStates {
			p.PrintState()
			fmt.Println()
		}
	}

	// Then we resume execution until the whole matrix is processed.
	for !p.ShouldStop() {
		p.Next(0)
		if *allStates {
			p.PrintState()
			fmt.Println()
		}
	}

	fmt.Println("Result:")
	p.PrintState()
}
