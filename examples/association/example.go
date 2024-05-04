package association

import (
	"fmt"
	"github.com/bchisham/collections-go/association"
	"github.com/bchisham/collections-go/examples"
	"github.com/bchisham/collections-go/pair"
	"github.com/bchisham/collections-go/sequence"
)

func Examples() {
	mapOfInts := association.FromMap(map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5})
	fmt.Println("map of ints")
	_ = mapOfInts.Each(examples.PrintItem[int])
	fmt.Println("Every")
	//every item is greater than 0
	every, err := mapOfInts.Every(func(v int) (bool, error) {
		return v > 0, nil
	})
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("every item is greater than 0: ", every)
	}
	fmt.Println("Filter")
	//filter the map
	filtered, err := mapOfInts.Where(func(v int) (bool, error) {
		return v > 3, nil
	})
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("filtered map")
		_ = filtered.Each(examples.PrintItem[int])
	}

	mapOfIntToString := association.NewMapTransform[string](mapOfInts).TransformMust(func(v int) string {
		return fmt.Sprintf("%d", v)
	})
	fmt.Println("map of strings")
	_ = mapOfIntToString.Each(examples.PrintItem[string])

	fmt.Println("join")
	joiner := association.NewJoiner[string](mapOfInts)
	joined, err := joiner.Join(mapOfIntToString)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("joined map")
		_ = joined.Each(examples.PrintItem[pair.Type[int, string]])
	}

	fmt.Println("Zip")
	listOfInts := sequence.FromSlice([]int{1, 2, 3, 4, 5})
	listOfStrings := sequence.FromSlice([]string{"one", "two", "three", "four", "five"})
	zipped := association.Zip(listOfInts, listOfStrings)
	fmt.Println("zipped map")
	_ = zipped.Each(examples.PrintItem[string])

}
