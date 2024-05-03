package stream

import (
	"collections-go/examples"
	"collections-go/sequence"
	"collections-go/stream"
	"context"
	"fmt"
	"sync"
	"time"
)

func ChannelExample() {
	var err error
	ctx := context.TODO()
	dlCtx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	defer cancelFunc()

	ch := stream.NewChan[int]()
	defer func() { _ = ch.Close() }()

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
