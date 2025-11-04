package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main(){
	values := make(chan int)

	shutdown := make(chan struct{})

	poolSize := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				n := rand.Intn(1000)

				select {
				case values <- n:
				case <-shutdown:
					fmt.Printf("Рабочий %d: Заыершение работы\n", id)
					return
				}
			}
		}(i)
	}

	var numbers []int

	for n := range values {
		if n%2 == 0{
			fmt.Printf("Отбрасываем чётное число: %d\n", n)
			continue
		}

		numbers = append(numbers, n)
		fmt.Printf("Полученное нечётное число: %d (всего: %d)\n", n, len(numbers))
		if len(numbers) == 100 {
			fmt.Println("Достигнуто 100 нечётных чисел. Остановка")
			break
		}
	}

	close(shutdown)
	wg.Wait()

	fmt.Printf("\nСобранно %d нечётных чисел:\n", len(numbers))
	for i, num := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, num)
	}
}

