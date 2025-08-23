package iter

import (
	"fmt"
	"github.com/bchisham/collections-go/iter"
)

func Examples() {
	// Example of using Inplace
	iter.Inplace([]int{1, 2, 3})(func(v int) int {
		return v * 2
	})

	// Example of using Each
	if _, err := iter.Each([]string{"a", "b", "c"})(func(v string) error {
		println(v)
		return nil
	}); err != nil {
		println("Error:", err.Error())
	}

	var strs = iter.MapMust[[]string]([]int{1, 2, 3})(func(v int) string {
		return fmt.Sprintf("%06d", v)
	})

	iter.EachMust(strs)(func(v string) {
		fmt.Println(v)
	})

	strs = []string{}
	ints := []int{1, 2, 3, 4, 5}
	strs = iter.MapAppendMust(ints, strs)(func(v int) string {
		return fmt.Sprintf("Number: %d", v)
	})

	iter.EachMust(strs)(func(v string) {
		fmt.Println(v)
	})

	// Example of using Filter
	ilist := iter.Filter([]int{1, 2, 3, 4})(func(v int) bool {
		return v%2 == 0
	})

	iter.EachMust(ilist)(func(v int) {
		fmt.Println("Even number:", v)
	})

	// Example of using Map
	mapped, err := iter.Map[[]string]([]int{1, 2, 3})(func(v int) (string, error) {
		return fmt.Sprintf("Number: %d", v), nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	iter.EachMust(mapped)(func(v string) {
		fmt.Println(v)
	})
}
