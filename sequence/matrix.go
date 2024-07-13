package sequence

import (
	"github.com/bchisham/collections-go/contracts"
	"strings"
)

type matrix[T contracts.NumericType] struct {
	rows, cols int
	basis      []contracts.Vector[T]
}

func FromBasis[T contracts.NumericType](m contracts.Sequence[contracts.Vector[T]]) (contracts.Matrix[T], error) {
	if err := validateBasis(m); err != nil {
		return nil, err
	}
	cols := m.Length()
	firstCol, _ := m.First()
	rows := firstCol.Length()

	return matrix[T]{
		rows:  rows,
		cols:  cols,
		basis: m.ToSlice(),
	}, nil
}

func FromSlices[T contracts.NumericType](slices [][]T) (contracts.Matrix[T], error) {
	basis := make([]contracts.Vector[T], len(slices))
	for i := 0; i < len(slices); i++ {
		basis[i] = FromNumericSlice(slices[i])
	}
	return matrix[T]{
		rows:  basis[0].Length(),
		cols:  len(slices),
		basis: basis,
	}, nil
}

func (m matrix[T]) Add(other contracts.Matrix[T]) (contracts.Matrix[T], error) {
	if m.rows != other.Rows() || m.cols != other.Cols() {
		return nil, ErrInvalidDimensions
	}
	resultBasis := make([]contracts.Vector[T], m.cols)
	for i := 0; i < m.cols; i++ {
		v1 := m.ColumnMust(i)
		v2 := other.ColumnMust(i)
		resultBasis[i] = v1.Add(v2)
	}
	result, _ := FromBasis(FromSlice(resultBasis))
	return result, nil
}

func (m matrix[T]) Subtract(other contracts.Matrix[T]) (contracts.Matrix[T], error) {
	if m.rows != other.Rows() || m.cols != other.Cols() {
		return nil, ErrInvalidDimensions
	}
	resultBasis := make([]contracts.Vector[T], m.cols)
	for i := 0; i < m.cols; i++ {
		v1 := m.ColumnMust(i)
		v2 := other.ColumnMust(i)
		resultBasis[i] = v1.Subtract(v2)
	}
	result, _ := FromBasis(FromSlice(resultBasis))
	return result, nil
}

func (m matrix[T]) ToBasis() contracts.Sequence[contracts.Vector[T]] {
	return FromSlice(m.basis)
}

func (m matrix[T]) Rows() int {
	return m.rows
}

func (m matrix[T]) Cols() int {
	return m.cols
}

func (m matrix[T]) Column(index int) (contracts.Vector[T], error) {

	return m.basis[index], nil
}

func (m matrix[T]) ColumnMust(index int) contracts.Vector[T] {

	return m.basis[index]
}

func (m matrix[T]) Row(index int) (_ contracts.Vector[T], err error) {
	if index >= m.rows {
		return nil, ErrIndexOutOfBounds
	}
	row := make([]T, m.cols)
	for i := 0; i < m.cols; i++ {
		row[i], err = m.basis[i].ToSequence().At(index)
		if err != nil {
			return nil, err
		}
	}
	return FromNumericSlice(row), nil
}

func (m matrix[T]) RowMust(index int) contracts.Vector[T] {
	row := make([]T, m.cols)
	for i := 0; i < m.cols; i++ {
		row[i], _ = m.basis[i].ToSequence().At(index)
	}
	return FromNumericSlice(row)
}

func (m matrix[T]) CanMultiply(other contracts.Matrix[T]) bool {
	return m.rows == other.Cols()
}

func (m matrix[T]) Multiply(other contracts.Matrix[T]) (contracts.Matrix[T], error) {
	if !m.CanMultiply(other) {
		return nil, ErrInvalidDimensions
	}
	product := make([][]T, other.Cols())
	for i := 0; i < other.Cols(); i++ {
		product[i] = make([]T, m.Rows())
	}

	for c := 0; c < other.Cols(); c++ {
		for r := 0; r < m.Rows(); r++ {
			row, err := m.Row(r)
			if err != nil {
				return nil, err
			}
			column, err := other.Column(c)
			if err != nil {
				return nil, err
			}
			product[c][r] = row.DotProduct(column)
		}
	}

	result, _ := FromSlices(product)
	return result, nil
}

func (m matrix[T]) ScalarMultiply(scalar T) contracts.Matrix[T] {

	var resultBasis = make([]contracts.Vector[T], m.cols)
	for i := 0; i < m.cols; i++ {
		resultBasis[i] = m.ColumnMust(i).Scale(scalar)
	}
	result, _ := FromBasis(FromSlice(resultBasis))
	return result
}

func (m matrix[T]) Transpose() contracts.Matrix[T] {

	transpose := make([][]T, m.Rows())

	for i := 0; i < m.Rows(); i++ {
		transpose[i] = make([]T, m.Cols())
	}
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			transpose[r][c] = m.basis[c].ToSequence().ToSlice()[r]
		}
	}
	result, _ := FromSlices(transpose)
	return result
}

func (m matrix[T]) String() string {

	buf := strings.Builder{}
	buf.WriteString("[")
	for r := 0; r < m.rows; r++ {
		row := m.RowMust(r)
		buf.WriteString(row.String())
		if r < m.rows-1 {
			buf.WriteString(", ")

		}
	}
	buf.WriteString("]")
	return buf.String()
}

func validateBasis[T contracts.NumericType](basis contracts.Sequence[contracts.Vector[T]]) error {
	if basis.Length() == 0 {
		return ErrEmptyBasis
	}
	first, _ := basis.First()
	if first.Length() == 0 {
		return ErrEmptyVector
	}
	length := first.Length()
	if ok := basis.EveryMust(func(v contracts.Vector[T]) bool {
		return v.Length() == length
	}); !ok {
		return ErrAssertionFailed
	}
	return nil
}
