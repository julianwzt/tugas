package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	matrix2 := [][]int{{10},
		{13},
		{16}}

	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	rows2 := len(matrix2)
	cols2 := len(matrix2[0])

	if cols1 != rows2 {
		fmt.Println("Cannot multiply the matrices. Invalid dimensions.")
		return
	}

	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}

	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	fmt.Println("Result:")
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
