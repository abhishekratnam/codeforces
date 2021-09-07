package main

import "fmt"

func main() {
	var rows, columns int

	var matrix [10][10]int
	var matrixtranspose [10][10]int

	fmt.Print("Enter the Matrix rows and Columns ")
	fmt.Scan(&rows, &columns)

	fmt.Println("Enter Matrix Items to Transpose ")
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matrixtranspose[j][i] = matrix[i][j]
		}
	}
	fmt.Print("Transpose of Matrix ")
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			fmt.Print(matrixtranspose[i][j], " ")
		}
		fmt.Println()
	}
}
