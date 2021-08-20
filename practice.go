package main

import (
	"fmt"
	"strconv"
)


func longEquation(m []string) float64 {
	if len(m)%2 != 1 {
		return 0
	}

	x := m[:3]
	y := m[3:]
	val, _ := strconv.ParseFloat(x[0], 64)
	val2, _ := strconv.ParseFloat(x[2], 64)
	storage := process(val, x[1], val2)

	if len(y) == 0 {
		return storage
	}

	var ResultStorage []string
	ResultStorage = append(ResultStorage, fmt.Sprintf("%g", storage))
	ResultStorage = append(ResultStorage, y...)
	return longEquation(ResultStorage)
}
