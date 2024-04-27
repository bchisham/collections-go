# Collections Go

## Description
Collections Go is a library that provides a set of collections for Go.
It provides methods to manipulate collections in a functional way.

## Sequence
A sequence is a collection that has an order. It is possible to access elements by index.
The sequence interface is defined as follows:

```go
package contracts
type Sequence[T any] interface {
	ToSlice() []T
	Each(f ApplyFunc[T]) error
	Where(f UnaryPredicate[T]) (Sequence[T], error)
	WhereMust(f UnaryPredicateMust[T]) Sequence[T]
	FirstWhereMust(predicate UnaryPredicateMust[T]) (result T, found bool)
	WithContext(ctx context.Context) ContextualSequence[T]
}
```

A sequence can be created from a slice using the `sequence.New` function:

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
```


The `ToSlice` method returns a slice with the elements of the sequence. This method is useful when you need to pass the elements of the sequence to a function that expects a slice.

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
slice := seq.ToSlice()
```

The `Each` method applies a function to each element of the sequence. The function must have the following signature:

```go  
type ApplyFunc[T any] func(T) error
```

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
err := seq.Each(func(i int) error {
    fmt.Println(i)
    return nil
})
```


The `Where` method filters the elements of the sequence based on a predicate. The predicate must have the following signature:

```go
type UnaryPredicate[T any] func(T) bool
```

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
filtered, err := seq.Where(func(i int) bool {
    return i%2 == 0
})
```

The `WhereMust` method filters the elements of the sequence based on a predicate that may return an error. The predicate must have the following signature:

```go
type UnaryPredicateMust[T any] func(T) (bool, error)
```

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
filtered := seq.WhereMust(func(i int) (bool, error) {
    return i%2 == 0, nil
})
```

The `FirstWhereMust` method returns the first element that satisfies a predicate that may return an error. The predicate must have the following signature:

```go
type UnaryPredicateMust[T any] func(T) (bool, error)
```

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
result, found := seq.FirstWhereMust(func(i int) (bool, error) {
    return i%2 == 0, nil
})
```

### Transforming a sequence
A sequence transformer is defined by the following interface:

```go
package sequence
type Transformer[T any, U any] interface {
	TransformMust(f TransformFunc[T, U]) Type[U]
}
```

The `TransformMust` method applies a function to each element of the sequence and returns a new sequence with the transformed elements. The function must have the following signature:

```go
type TransformFunc[T any, U any] func(T) (U, error)
```

```go
seq := sequence.New([]int{1, 2, 3, 4, 5})
transformed := seq.TransformMust(func(i int) (string, error) {
    return strconv.Itoa(i), nil
})
```
