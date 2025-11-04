package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numGoroutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ch := make(chan int, numGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			n := rand.Intn(1000)
			fmt.Printf("Гоурутина %d создана: %d\n", id, n)
			ch <- n
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	numbers := make([]int, 0, numGoroutines)
	for n := range ch {
		numbers = append(numbers, n)
	}

	fmt.Printf("\nСобрано %d случайных чисел:\n", len(numbers))
	for i, num := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, num)
	}
}

