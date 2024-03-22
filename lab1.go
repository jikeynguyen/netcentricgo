package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	// Calculating differences in two equally long strings.
	if len(a) != len(b) {
		return 0, errors.New("Different Lengths")
	}

	differences := 0
	for i := range a {
		if a[i] != b[i] {
			differences++
		}
	}
	return differences, nil
}

func main() {
	a := "GAGCCTACTAACGGGAT"
	b := "CATCGTAATGACGGCCT"
	distance, err := Distance(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The Hamming Distance is %d\n", distance)
}
