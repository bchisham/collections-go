package sequence

import (
	"fmt"
	"github.com/bchisham/collections-go/contracts"
	"github.com/bchisham/collections-go/examples"
	"github.com/bchisham/collections-go/sequence"
	"strconv"
)

func Examples() {
	listOfInts := sequence.FromSlice([]int{1, 2, 3, 4, 5})
	fmt.Println("list of ints")
	_ = listOfInts.Each(examples.PrintItem[int])
	fmt.Println("Sum")
	//sum of the list
	sum := sequence.Sum(listOfInts.ToSlice(), func(i int) int {
		return i
	})

	fmt.Println("sum: ", sum)
	fmt.Println("Mean")
	//mean of the list
	mean := sequence.Mean(listOfInts.ToSlice(), func(i int) int {
		return i

	})
	fmt.Println("mean: ", mean)

	fmt.Println("FirstWhereMust")
	//first greater than 3
	first, found := listOfInts.FirstWhereMust(func(i int) bool {
		return i > 3
	})
	if !found {
		fmt.Println("no item found")
	} else {
		fmt.Println("first item greater than 3: ", first)
	}
	fmt.Println("WhereMust")
	//filter the list
	filtered := listOfInts.WhereMust(func(i int) bool {
		return i > 3
	})
	fmt.Println("filtered list")
	_ = filtered.Each(func(i int) error {
		return examples.PrintItem(i)
	})
	fmt.Println("TransformMust")
	//transform to a list of strings
	listOfStrings := sequence.NewTransformer[int, string](listOfInts).TransformMust(strconv.Itoa)
	fmt.Println("list of strings")
	_ = listOfStrings.Each(examples.PrintItem[string])

	smallInts, largeInts := sequence.PartitionMust(listOfInts, func(i int) bool {
		return i < 3
	})
	fmt.Println("small ints")
	_ = smallInts.Each(examples.PrintItem[int])
	fmt.Println("large ints")
	_ = largeInts.Each(examples.PrintItem[int])

	chunk := sequence.Chunk(listOfInts, 3)
	fmt.Println("chunks of 3")
	_ = chunk.Each(func(seq contracts.Sequence[int]) error {
		fmt.Println("chunk")
		_ = seq.Each(examples.PrintItem[int])
		return nil
	})

	v1 := sequence.FromNumericSlice([]int{1, 2, 3, 4, 5})
	v2 := sequence.FromNumericSlice([]int{1, 2, 3, 4, 5})
	fmt.Println("dot product")
	dotProduct := v1.DotProduct(v2)
	fmt.Println("dot product: ", dotProduct)
	fmt.Println("cross product")
	crossProduct := v1.CrossProduct(v2)
	fmt.Println("cross product: ", crossProduct.String())
	fmt.Println("add")
	added := v1.Add(v2)
	fmt.Println("added", added.String())
	fmt.Println("subtract")
	subtracted := v1.Subtract(v2)
	fmt.Println("subtracted", subtracted.String())
	fmt.Println("scale")
	scaled := v1.Scale(2)
	fmt.Println("scaled", scaled.String())

	vecX := sequence.FromNumericSlice([]int{1, 0, 0})
	vecY := sequence.FromNumericSlice([]int{0, 1, 0})
	vexZ := sequence.FromNumericSlice([]int{0, 0, 1})

	matrix, _ := sequence.FromBasis(sequence.FromSlice([]contracts.Vector[int]{vecX, vecY, vexZ}))
	fmt.Println("matrix")
	fmt.Println(matrix.String())
	fmt.Println("transpose")
	transposed := matrix.Transpose()
	fmt.Println("transposed")
	fmt.Println(transposed.String())
	fmt.Println("scalar multiply")
	twiceScaled := matrix.ScalarMultiply(2)
	fmt.Println("scalar multiplied")
	fmt.Println(twiceScaled.String())
	fmt.Println("multiply")
	threeXScaled := matrix.ScalarMultiply(3)
	multiplied, _ := twiceScaled.Multiply(threeXScaled)
	fmt.Println("multiplied")
	fmt.Println(multiplied.String())
	//
	//m1, _ := sequence.FromBasis(sequence.FromItems(v1, v2))
	//m2 := m1.Transpose()

	//fmt.Println("multiplication when r <> c")
	//nsMultiply, _ := m1.Multiply(m2)
	//fmt.Println("result")
	//fmt.Println(nsMultiply.String())

}
