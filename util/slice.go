package util

import "fmt"

// 0 <= index <= len(a)
func Insert[T any](a []T, index int, value T) []T {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Print2D[T any](matrix [][]T) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func All[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v != value {
			return false
		}
	}
	return true
}

func ConvertToMatrix(lines []string) [][]string {
	var matrix [][]string
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// order is not important
func Remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// order is important
func RemoveStable[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
