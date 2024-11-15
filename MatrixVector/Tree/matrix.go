package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

type Matrix struct {
	Width  int
	Height int
	Data   [][]int
}

func GetMatrixFromFile(fp string) (*Matrix, error) {
	b, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	type MatrixJSON struct {
		Width  int   `json:"Width"`
		Height int   `json:"Height"`
		Data   []int `json:"Data"`
	}

	var m MatrixJSON

	err = json.Unmarshal(b, &m)

	if err != nil {
		return nil, err
	}

	if m.Height*m.Width != len(m.Data) {
		return nil, errors.New("the size of the matrix is different from the size of the data")
	}

	var toret Matrix
	toret.Width = m.Width
	toret.Height = m.Height

	lastcol := 0

	for r := 0; r < m.Height; r++ {
		// create the first row
		newrow := make([]int, 0)

		// populates the first row
		for c := lastcol; c < m.Width+lastcol; c++ {
			newrow = append(newrow, m.Data[c])
		}
		toret.Data = append(toret.Data, newrow)
		lastcol += m.Width
	}

	return &toret, nil
}

func ComputeHeight(h int) int {
	power := 0.0

	for math.Pow(2.0, power) < float64(h) {
		power += 1.0
	}

	return int(power)
}

func (m *Matrix) Adjust() {
	h := ComputeHeight(m.Width) + 1

	for i := 0; i < m.Height+h; i++ {
		m.Data = append(m.Data, make([]int, m.Width))
	}

	// We update the matrix height.
	m.Height = m.Height + h
}

func (m *Matrix) Print() {
	for r := 0; r < m.Height; r++ {
		for c := 0; c < m.Width; c++ {
			fmt.Printf("%d ", m.Data[r][c])
		}
		fmt.Println()
	}
}
