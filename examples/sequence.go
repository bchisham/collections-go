package examples

import (
	"collections-go/sequence"
	"fmt"
	"strconv"
)

func PrintItem[T any](item T) error {
	_, err := fmt.Println("Item: ", item)
	return err
}

func SequenceExamples() {
	listOfInts := []int{1, 2, 3, 4, 5}
	fmt.Println("list of ints")
	_ = sequence.Apply(listOfInts, func(i int) error {
		return PrintItem(i)
	})
	fmt.Println("Sum")
	//sum of the list
	sum := sequence.Sum(listOfInts, func(i int) int {
		return i
	})

	fmt.Println("sum: ", sum)
	fmt.Println("Mean")
	//mean of the list
	mean := sequence.Mean(listOfInts, func(i int) int {
		return i

	})
	fmt.Println("mean: ", mean)

	fmt.Println("FirstWhere")
	//first greater than 3
	first, found := sequence.FirstWhere(listOfInts, func(i int) bool {
		return i > 3
	})
	if !found {
		fmt.Println("no item found")
	} else {
		fmt.Println("first item greater than 3: ", first)
	}
	fmt.Println("Where")
	//filter the list
	filtered := sequence.Where(listOfInts, func(i int) bool {
		return i > 3
	})
	fmt.Println("filtered list")
	_ = sequence.Apply(filtered, func(i int) error {
		return PrintItem(i)
	})
	fmt.Println("Transform")
	//transform to a list of strings
	listOfStrings := sequence.Transform(listOfInts, func(i int) string {
		return "'" + strconv.Itoa(i) + "'"
	})
	fmt.Println("list of strings")
	_ = sequence.Apply(listOfStrings, func(s string) error {
		return PrintItem(s)
	})

}
