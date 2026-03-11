package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		strNums := strings.Fields(line)
		row := make([]int, len(strNums))
		for i, s := range strNums {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			row[i] = num
		}
		matrix = append(matrix, row)
	}
	return matrix, scanner.Err()
}

func isLatinSquare(matrix [][]int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}

	// Проверка, что все строки и столбцы имеют длину n
	for _, row := range matrix {
		if len(row) != n {
			return false
		}
	}

	expected := make(map[int]bool)
	for i := 1; i <= n; i++ {
		expected[i] = true
	}

	// Проверка строк
	for _, row := range matrix {
		seen := make(map[int]bool)
		for _, val := range row {
			if !expected[val] || seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	// Проверка столбцов
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
	files := []string{"matrix_examples/matrix1.txt", "matrix_examples/matrix2.txt", "matrix_examples/matrix3.txt", "matrix_examples/matrix4.txt", "matrix_examples/matrix5.txt"}

	fmt.Println("Выберите файл для проверки матрицы:")
	for i, f := range files {
		fmt.Printf("%d. %s\n", i+1, f)
	}

	var choice int
	fmt.Scan(&choice)
	if choice < 1 || choice > 5 {
		fmt.Println("Некорректный выбор")
		return
	}

	matrix, err := readMatrix(files[choice-1])
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
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