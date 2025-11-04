package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const goroutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			n := rand.Intn(1000)
			if n%2 == 0 {
				fmt.Printf("Гоурутина %d: %d - чётное, пропускаем")
				return
			}
			fmt.Printf("Гоурутина %d: отправка нечётного числа %d\n", id, n)
			ch <- n
		}(i)
	}	

		go func(){
			wg.Wait()
			close(ch)
		}()

		var numbers []int
		for n := range ch {
			numbers = append(numbers, n)
		}

		fmt.Printf("\nПолучено %d нечёиных чисел:\n", len(numbers))
		for i, num := range numbers {
			fmt.Printf("numbers[%d] = %d\n", i, num)
		}

}
