package association

import (
	"collections-go/association"
	"collections-go/examples"
	"collections-go/pair"
	"fmt"
)

func Examples() {
	mapOfInts := association.FromMap(map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5})
	fmt.Println("map of ints")
	_ = mapOfInts.Each(examples.PrintItem[int])
	fmt.Println("Every")
	//every item is greater than 0
	every, err := mapOfInts.Every(func(k int, v int) (bool, error) {
		return v > 0, nil
	})
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("every item is greater than 0: ", every)
	}
	fmt.Println("Filter")
	//filter the map
	filtered, err := mapOfInts.Where(func(k int, v int) (bool, error) {
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
}