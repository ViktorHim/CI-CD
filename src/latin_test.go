package main

import "testing"

func TestIsLatinSquare_Valid2x2(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{2, 1},
	}
	if !isLatinSquare(matrix) {
		t.Error("2x2 латинский квадрат не распознан")
	}
}

func TestIsLatinSquare_Valid3x3(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}
	if !isLatinSquare(matrix) {
		t.Error("3x3 латинский квадрат не распознан")
	}
}

func TestIsLatinSquare_Invalid_RowRepeat(t *testing.T) {
	matrix := [][]int{
		{1, 1, 3},
		{2, 3, 1},
		{3, 2, 2},
	}
	if isLatinSquare(matrix) {
		t.Error("повтор в строке должен вернуть false")
	}
}

func TestIsLatinSquare_Invalid_ColRepeat(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{3, 1, 2},
	}
	if isLatinSquare(matrix) {
		t.Error("повтор в столбце должен вернуть false")
	}
}

func TestIsLatinSquare_Empty(t *testing.T) {
	matrix := [][]int{}
	if isLatinSquare(matrix) {
		t.Error("пустая матрица должна вернуть false")
	}
}

func TestIsLatinSquare_1x1(t *testing.T) {
	matrix := [][]int{{1}}
	if !isLatinSquare(matrix) {
		t.Error("1x1 матрица с [1] — валидный латинский квадрат")
	}
}

func TestIsLatinSquare_WrongValues(t *testing.T) {
	matrix := [][]int{
		{0, 2},
		{2, 0},
	}
	if isLatinSquare(matrix) {
		t.Error("значения вне диапазона 1..n должны вернуть false")
	}
}

func TestIsLatinSquare_NotSquare(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{2, 3, 1},
	}
	if isLatinSquare(matrix) {
		t.Error("не квадратная матрица должна вернуть false")
	}
}