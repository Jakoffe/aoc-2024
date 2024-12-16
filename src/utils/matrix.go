package utils

import (
	"fmt"
	"strings"
)

type Matrix[T comparable] struct {
	Data [][]T
}

func NewMatrix[T comparable]() *Matrix[T] {
	return &Matrix[T]{
		Data: [][]T{},
	}
}

func (m *Matrix[T]) Height() int {
	return len(m.Data)
}

func (m *Matrix[T]) Width() int {
	return len(m.Data[0])
}

func (m *Matrix[T]) AddRow(row []T) {
	m.Data = append(m.Data, row)
}

func (m *Matrix[T]) Set(row int, column int, value T) bool {
	// Check if row and column are valid
	if row < 0 || column < 0 || row >= len(m.Data) || column >= len(m.Data[0]) {
		return false
	}

	m.Data[row][column] = value

	return true
}

func (m *Matrix[T]) Get(row int, col int) T {
	return m.Data[row][col]
}

func (m *Matrix[T]) GetRow(y int) []T {
	return m.Data[y]
}

func (m *Matrix[T]) GetColumn(col int) []T {
	result := []T{}
	for _, row := range m.Data {
		result = append(result, row[col])
	}
	return result
}

func (m *Matrix[T]) GetDiagonal(row int, col int, absolute bool) []T {
	result := []T{}

	i, j := row, col
	if absolute {
		for i > 0 && j > 0 {
			i--
			j--
		}
	}

	for i <= m.Height()-1 && j <= m.Width()-1 {
		result = append(result, m.Get(i, j))
		i++
		j++
	}

	return result
}

// GetAntiDiagonal returns the anti-diagonal of the matrix from a given row and column
func (m *Matrix[T]) GetAntiDiagonal(row int, col int, absolute bool) []T {
	result := []T{}

	i, j := row, col
	if absolute {
		for i > 0 && j < m.Width()-1 {
			i--
			j++
		}
	}

	for i <= m.Height()-1 && j >= 0 {
		result = append(result, m.Get(i, j))
		i++
		j--
	}

	return result
}

func (m Matrix[T]) Elements() []T {
	result := []T{}
	for _, row := range m.Data {
		for _, value := range row {
			result = append(result, value)
		}
	}
	return result
}

func (m Matrix[T]) FindAll(value T) []Match[T] {
	result := []Match[T]{}
	for i, row := range m.Data {
		for j, v := range row {
			if v == value {
				result = append(result, NewMatch(v, i, j))
			}
		}
	}

	return result
}

func (m Matrix[T]) Print() {
	for _, row := range m.Data {
		for j, value := range row {
			print(value)
			if j == len(row)-1 {
				print("\n")
			}
		}
	}
}

func (m Matrix[T]) PrintMapWithPosition(x int, y int) {
	println(strings.Repeat("=", len(m.Data[0])*3))
	for i, row := range m.Data {
		for j, value := range row {
			if i == x && j == y {
				fmt.Printf("[%v]", value)
			} else {
				fmt.Printf(" %v ", value)
			}
			if j == len(row)-1 {
				print("\n")
			}
		}
	}
	println(strings.Repeat("=", len(m.Data[0])*3))
}

type Match[T comparable] struct {
	Value  T
	Row    int
	Column int
}

func NewMatch[T comparable](value T, row int, column int) Match[T] {
	return Match[T]{value, row, column}
}

func (m Match[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", m.Value, m.Row, m.Column)
}
