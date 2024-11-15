package main

import (
	"flag"
	"fmt"
)

func main() {
	fm := flag.String("m1", "", "the filepath to the first matrix")
	sm := flag.String("m2", "", "the filepath to the second matrix")
	allStates := flag.Bool("a", false, "print the state after each step")

	flag.Parse()

	m1, err := GetMatrixFromFile(*fm, false)

	if err != nil {
		fmt.Printf("get matrix from file: %s\n", err)
		return
	}

	// We do the transpose of the second matrix, to read it better from the top.
	m2, err := GetMatrixFromFile(*sm, true)

	if err != nil {
		fmt.Printf("get matrix from file: %s\n", err)
		return
	}

	// Firstly we create the ProcMatrix
	pm := CreateProcMatrix(m1.Height, m2.Width, m1, m2)

	// Then we adjust both matrixes
	m1.Adjust()
	m2.Adjust()

	for !pm.ShouldStop() {
		pm.Next()
		if *allStates {
			pm.PrintState()
			fmt.Println()
		}
	}

	fmt.Println("Result:")
	pm.PrintResult()
}
