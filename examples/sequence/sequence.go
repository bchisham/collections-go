package sequence

import (
	"collections-go/sequence"
	"fmt"
	"strconv"
)

func PrintItem[T any](item T) error {
	_, err := fmt.Println("Item: ", item)
	return err
}

func Examples() {
	listOfInts := sequence.FromSlicef([]int{1, 2, 3, 4, 5})
	fmt.Println("list of ints")
	_ = listOfInts.Each(PrintItem[int])
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
		return PrintItem(i)
	})
	fmt.Println("TransformMust")
	//transform to a list of strings
	listOfStrings := sequence.NewTransformer[int, string](listOfInts).TransformMust(strconv.Itoa)
	fmt.Println("list of strings")
	_ = listOfStrings.Each(PrintItem[string])
}
