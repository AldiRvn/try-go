package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//? ref: https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de
//? ref: https://gist.github.com/YumaInaura/8d52e73dac7dc361745bf568c3c4ba37

//* Case: Need to process 1000 data, each 100 data will be processed conccurently.
//*		then the result is merged into 1 list

func asChan(listNum ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, num := range listNum {
			channel <- num
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(channel)
	}()
	return channel
}
func main() {
	fmt.Println("Start")
	defer func(now time.Time) {
		fmt.Println("Done in", time.Since(now))
	}(time.Now())

	a := asChan(1, 2, 3)
	b := asChan(4, 5, 6, 7, 8, 9, 10)
	c := asChan(123, 12, 31, 2312, 3, 14, 1, 9)
	for v := range merge(a, b, c) {
		fmt.Println(v)
	}
}

func merge(listChannel ...<-chan int) <-chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(listChannel))
	for _, channel := range listChannel {
		go func(listChannelValue <-chan int) {
			for channelValue := range listChannelValue {
				res <- channelValue
			}
			wg.Done()
		}(channel)
	}

	go func() {
		wg.Wait()
		// close(res)
	}()
	return res
}
