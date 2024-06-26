# Collections Go

## Description
Collections Go is a library that provides a set of collections for Go.
It provides methods to manipulate collections.

## Sequence
A sequence is a collection that has an order.
The sequence interface is defined as follows:

```go
package contracts
type Sequence[T any] interface {
	ToSlice() []T
	Each(f ApplyFunc[T]) error
	Where(f UnaryPredicate[T]) (Sequence[T], error)
	WhereMust(f UnaryPredicateMust[T]) Sequence[T]
	FirstWhereMust(predicate UnaryPredicateMust[T]) (result T, found bool)
}
```

A sequence can be created from a slice using the `sequence.FromSlice` function:

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
```


The `ToSlice` method returns a slice with the elements of the sequence. This method is useful when you need to pass the elements of the sequence to a function that expects a slice.

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
slice := seq.ToSlice()
```

The `Each` method applies a function to each element of the sequence. The function must have the following signature:

```go  
type ApplyFunc[T any] func(T) error
```

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
err := seq.Each(func(i int) error {
    fmt.Println(i)
    return nil
})
```


The `WhereMust` method filters the elements of the sequence based on a predicate. The predicate must have the following signature:

```go
type UnaryPredicateMust[T any] func(T) bool
```

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
filtered, err := seq.WhereMust(func(i int) bool {
    return i%2 == 0
})
```

The `Where` method filters the elements of the sequence based on a predicate that may return an error. The predicate must have the following signature:

```go
type UnaryPredicate[T any] func(T) (bool, error)
```

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
filtered := seq.Where(func(i int) (bool, error) {
    return i%2 == 0, nil
})
```

The `FirstWhere` method returns the first element that satisfies a predicate that may return an error. The predicate must have the following signature:

```go
type UnaryPredicate[T any] func(T) (bool, error)
```

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
result, found := seq.FirstWhere(func(i int) (bool, error) {
    return i%2 == 0, nil
})
```

The `EveryMust` method returns true if all elements of the sequence satisfy a predicate. The predicate must have the following signature:

```go

type UnaryPredicate[T any] func(T) bool
```

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
result := seq.EveryMust(func(i int) bool {
    return i > 0
})
```

### Transforming a sequence
A sequence transformer is defined by the following interface:

```go
package sequence
type Transformer[T any, U any] interface {
	TransformMust(f TransformFuncMust[T, U]) Type[U]
}
```

The `Transform` method applies a function to each element of the sequence and returns a new sequence with the transformed elements. The function must have the following signature:

```go
type TransformFunc[T any, U any] func(T) (U, error)
```

```go
seq := sequence.FromSlice([]int{1, 2, 3, 4, 5})
transformed := seq.Transform(func(i int) (string, error) {
    return strconv.Itoa(i), nil
})
```

### Contextual Sequence 

This type extends the basic Sequence concepts but the function signatures but the corresponding function signatures accept a context.Context.

```go
sequence.WithContext(context.TODO()).Build({}int{1, 2, 3, 4, 5}).Where(func (ctx context.Context, i int) (bool, error){return i%2, nil })
```
## Map
Is an abstraction of an associative container, and implemented for the golang map[K]V type. The `Map[K,V]` contract defines`Where`, `WhereMust`,
`Each`, and `Every` operations analogous to those defined for the `Sequence[T]` contract.

### Joiner
The `Joiner` contract defines the operations that join the given maps by key producing a map of keys to a join product as a `Map[K, pair.Type[FirstType, SecondType]`

```go
type Joiner[SecondType,K, FirstType] {
   Join(other Map[K, SecondType]) Map[K, pair.Pair[FirstType, SecondType]], error
   JoinMust(other Map[K, SecondType]) Map[pair.Pair[FirstType, SecondType]]
}
```

```go
func main() {
	mapOfInts := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	mapOfIntToString := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
    joiner := association.NewJoiner[string](mapOfInts)
	joined, err := joiner.Join(mapOfIntToString)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("joined map")
		_ = joined.Each(examples.PrintItem[pair.Type[int, string]])
	}
}
   ```

```plain
joined map
1: {1 one}
2: {2 two}
3: {3 three}
4: {4 four}
5: {5 five}
```

### Map Transformer

Map transformer operates analogously to the sequence Transform.

```go
func main() {
    mapOfIntToString := association.NewMapTransform[string](mapOfInts).TransformMust(func(v int) string {
    return fmt.Sprintf("%d", v)
    })
}
```

```plain
1: 1
2: 2
3: 3
4: 4
5: 5
```

## Channel 

The `Channel` contract defines the operations that can be performed on a channel. The `Channel[T]` contract defines Send, Receive, and Close operations.

```go
func main() {
var err error
	ctx := context.TODO()
	dlCtx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	defer cancelFunc()

	ch := stream.NewChan[int]()

	sliceOfInts := sequence.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = stream.SendSequence(ch, sliceOfInts)
		if err != nil {
			fmt.Println(err)
		}
	}()
	outputSequence, _ := stream.IteratorToSlice(dlCtx, ch)

	wg.Wait()

	err = outputSequence.Each(examples.PrintItem[int])
	if err != nil {
		fmt.Println(err)
	}

}