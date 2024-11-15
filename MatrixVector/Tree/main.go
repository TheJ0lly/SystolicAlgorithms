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

	v, err := GetVectorFromFile(*vectorFile)

	if err != nil {
		fmt.Printf("vector file error: %s\n", err)
		return
	}

	m.Adjust()

	tree := CreateTree(m.Width, m)

	tree.LoadVector(v)

	for !tree.ShouldStop() {
		tree.Next()

		if *allStates {
			tree.PrintState()
			fmt.Println()
		}
	}

	fmt.Println("Result:")
	tree.PrintResults()
}
