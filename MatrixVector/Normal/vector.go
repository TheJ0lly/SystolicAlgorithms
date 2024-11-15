package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Vector struct {
	Size int
	Data []int
}

func GetVectorFromFile(fp string) (*Vector, error) {
	b, err := os.ReadFile(fp)

	if err != nil {
		return nil, err
	}

	type VectorJSON struct {
		Size int   `json:"Size"`
		Data []int `json:"Data"`
	}

	var v VectorJSON

	err = json.Unmarshal(b, &v)

	if err != nil {
		return nil, err
	}

	if v.Size != len(v.Data) {
		return nil, errors.New("the size of the vector is different from the size of the data")
	}

	var toret Vector
	toret.Size = v.Size

	toret.Data = append(toret.Data, v.Data...)

	return &toret, nil
}

func (v *Vector) Print() {
	for i := 0; i < v.Size; i++ {
		fmt.Printf("%d ", v.Data[i])
	}
	fmt.Println()
}
