package sequence

import (
	"github.com/bchisham/collections-go/contracts"
	"strings"
)

type matrix[T contracts.NumericType] struct {
	rows, cols int
	basis      contracts.Sequence[contracts.Vector[T]]
}

func FromBasis[T contracts.NumericType](m contracts.Sequence[contracts.Vector[T]]) (contracts.Matrix[T], error) {
	if err := validateBasis(m); err != nil {
		return nil, err
	}
	rows := m.Length()
	firstCol, _ := m.First()
	cols := firstCol.Length()

	return matrix[T]{
		rows:  rows,
		cols:  cols,
		basis: m,
	}, nil
}

func FromSlices[T contracts.NumericType](slices [][]T) (contracts.Matrix[T], error) {
	basis := make([]contracts.Vector[T], len(slices))
	for i := 0; i < len(slices); i++ {
		basis[i] = FromNumericSlice(slices[i])
	}
	seqBases := FromSlice(basis)
	return matrix[T]{
		rows:  len(slices),
		cols:  basis[0].Length(),
		basis: seqBases,
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
	return m.basis
}

func (m matrix[T]) Rows() int {
	return m.rows
}

func (m matrix[T]) Cols() int {
	return m.cols
}

func (m matrix[T]) Column(index int) (contracts.Vector[T], error) {
	column := make([]T, m.rows)

	c, err := m.basis.At(index)
	if err != nil {
		return nil, err
	}

	for i := 0; i < m.rows; i++ {
		column[i], _ = c.ToSequence().At(i)
	}
	return FromNumericSlice(column), nil
}

func (m matrix[T]) ColumnMust(index int) contracts.Vector[T] {
	column := make([]T, m.rows)
	c, _ := m.basis.At(index)
	for i := 0; i < m.rows; i++ {
		column[i], _ = c.ToSequence().At(i)
	}
	return FromNumericSlice(column)
}

func (m matrix[T]) Row(index int) (contracts.Vector[T], error) {
	row := make([]T, m.cols)
	for i := 0; i < m.cols; i++ {
		c, err := m.basis.At(i)
		if err != nil {
			return nil, err
		}
		row[i], _ = c.ToSequence().At(index)
	}
	return FromNumericSlice(row), nil
}

func (m matrix[T]) RowMust(index int) contracts.Vector[T] {
	row := make([]T, m.cols)
	for i := 0; i < m.cols; i++ {
		c, _ := m.basis.At(i)
		row[i], _ = c.ToSequence().At(index)
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
	product := make([][]T, m.rows)
	for i := 0; i < m.rows; i++ {
		product[i] = make([]T, other.Cols())
	}

	for c := 0; c < m.cols; c++ {
		for r := 0; r < m.rows; r++ {
			row, _ := m.Row(r)
			column, _ := other.Column(c)
			product[r][c] = row.DotProduct(column)
		}
	}

	resultBasis := make([]contracts.Vector[T], m.rows)
	for i := 0; i < m.cols; i++ {
		resultBasis[i] = FromNumericSlice(product[i])
	}
	return FromBasis(FromSlice(resultBasis))
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
	transpose := make([][]T, m.cols)
	for i := 0; i < m.cols; i++ {
		transpose[i] = make([]T, m.rows)
	}

	rhs := m.basis.ToSlice()

	for c := 0; c < len(rhs); c++ {
		for r := 0; r < m.rows; r++ {
			v, _ := rhs[c].ToSequence().At(r)
			transpose[r][c] = v
		}
	}

	resultBasis := make([]contracts.Vector[T], len(transpose))
	for i := 0; i < m.cols; i++ {
		resultBasis[i] = FromNumericSlice(transpose[i])
	}
	result, _ := FromBasis(FromSlice(resultBasis))
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
