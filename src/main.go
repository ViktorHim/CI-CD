package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanMatrix(size int) ([][]int, error) {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	matrix := make([][]int, size)

	fmt.Println("Введите матрицу (каждая строка через пробел):")
	for i := 0; i < size; i++ {
		fmt.Printf("Строка %d: ", i+1)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)

		matrix[i] = make([]int, size)
		for j, p := range parts {
			matrix[i][j], _ = strconv.Atoi(p)
		}
	}
	return matrix, reader.Err()
}

func isLatinSquare(matrix [][]int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}

	for _, row := range matrix {
		if len(row) != n {
			return false
		}
	}

	expected := make(map[int]bool)
	for i := 1; i <= n; i++ {
		expected[i] = true
	}

	for _, row := range matrix {
		seen := make(map[int]bool)
		for _, val := range row {
			if !expected[val] || seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	for col := 0; col < n; col++ {
		seen := make(map[int]bool)
		for row := 0; row < n; row++ {
			val := matrix[row][col]
			if !expected[val] || seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	return true
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Введите размерность матрицы:")
	var n int
	fmt.Scan(&n)

	matrix, err := scanMatrix(n)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Println("Матрица:")
	printMatrix(matrix)

	if isLatinSquare(matrix) {
		fmt.Println("Это латинский квадрат")
	} else {
		fmt.Println("Это НЕ латинский квадрат")
	}
}
