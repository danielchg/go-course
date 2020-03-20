package main

import (
	"fmt"
	"sync"
)

func main() {
	input := make([]int, 0, 100)

	for i := 0; i != cap(input); i++ {
		input = append(input, i)
	}
	fmt.Printf("content of input: %v", input)

	channels := make([]chan int, 5)
	for i := range channels {
		channels[i] = make(chan int)
	}

	go func() {
		defer func() {
			for _, v := range channels {
				close(v)
			}
		}()
		for _, v := range input {
			select {
			case channels[0] <- v:
			case channels[1] <- v:
			case channels[2] <- v:
			case channels[3] <- v:
			case channels[4] <- v:
			}
		}
	}()

	w := sync.WaitGroup{}
	w.Add(5)
	for i := range channels {
		go func(i int) {
			defer w.Done()
			for v := range channels[i] {
				fmt.Printf("Content channel %d: %d \n", i, v)

			}
		}(i)
	}

	w.Wait()

}
