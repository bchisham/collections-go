package sequence

import "github.com/bchisham/collections-go/contracts"

func DotProduct[T any, N contracts.NumericType](a []T, b []T, f func(T) N) N {
	var sum N
	for i, item := range a {
		sum += f(item) * f(b[i])
	}
	return sum
}

func CrossProduct[T any, N contracts.NumericType](a []T, b []T, f func(T) N) N {
	var sum N
	for i, item := range a {
		sum += f(item) * f(b[i])
	}
	return sum
}

func ToMatrix[T any, N contracts.NumericType](seq []T, f func(T) N, rows, cols int) [][]N {
	matrix := make([][]N, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]N, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = f(seq[i*cols+j])
		}
	}
	return matrix
}

func MatrixProduct[T any, N contracts.NumericType](a [][]N, b [][]N) [][]N {
	rows := len(a)
	cols := len(b[0])
	product := make([][]N, rows)
	for i := 0; i < rows; i++ {
		product[i] = make([]N, cols)
		for j := 0; j < cols; j++ {
			var sum N
			for k := 0; k < len(b); k++ {
				sum += a[i][k] * b[k][j]
			}
			product[i][j] = sum
		}
	}
	return product
}
