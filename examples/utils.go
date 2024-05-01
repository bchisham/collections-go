package examples

import (
	"fmt"
)

func PrintItem[T any](item T) error {
	_, err := fmt.Println("Item: ", item)
	return err
}
