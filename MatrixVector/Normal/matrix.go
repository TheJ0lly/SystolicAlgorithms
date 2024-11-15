package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
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

func (m *Matrix) Adjust() {
	for r := 0; r < m.Height; r++ {
		// We insert r number of 0 at the begining of the row
		// to create a delay in the matrix.
		for c := 0; c < r; c++ {
			m.Data[r] = slices.Insert(m.Data[r], c, 0)
		}

		// We append r number of 0 at the end of the row
		// to make the matrix even.
		for c := r; c < m.Height; c++ {
			m.Data[r] = append(m.Data[r], 0)
		}
	}

	// We update the matrix width.
	m.Width = len(m.Data[0])
}

func (m *Matrix) Print() {
	for r := 0; r < m.Height; r++ {
		for c := 0; c < m.Width; c++ {
			fmt.Printf("%d ", m.Data[r][c])
		}
		fmt.Println()
	}
}
