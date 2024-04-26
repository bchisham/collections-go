package examples

import (
	"collections-go/sequence"
	"fmt"
	"strconv"
)

func main() {
	listOfInts := []int{1, 2, 3, 4, 5}

	_ = sequence.Apply(listOfInts, func(i int) error {
		_, err := fmt.Println("Item: ", i)
		return err
	})

	//transform to a list of strings
	listOfStrings := sequence.Transform(listOfInts, func(i int) string {
		return strconv.Itoa(i)
	})

	_ = sequence.Apply(listOfStrings, func(s string) error {
		_, err := fmt.Println("Item: ", s)
		return err
	})

}
