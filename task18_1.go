package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		goroutine("Гоурутина 1", ch)
		wg.Done()
	}()
	go func() {
		goroutine("Гоурутина 2", ch)
		wg.Done()
	}()

	ch <- 1

	wg.Wait()
	fmt.Println("Программа завершена!")
}

func goroutine(name string, ch chan int) {
	for {
		value, ok := <-ch
		if !ok {
			fmt.Printf("%s: Канал закрыт, перезагрузка\n", name)
			return
		}
		fmt.Printf("%s: Получено: %d\n", name, value)
		if value == 10 {
			fmt.Printf("%s: Достигнут 10. Канал закрывается", name)
			close(ch)
			return
		}
		value++
		ch <- value
	}
}
