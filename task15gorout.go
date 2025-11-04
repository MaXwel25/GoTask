package main 

import (
		"fmt"
		"runtime"
		"sync"
)

func init(){
	runtime.GOMAXPROCS(1) // 1 логический процессор
}

func main () {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 	100; i >= 0; i-- {
			fmt.Printf("Гоурутина 1 (низ) - ID: %d, Value: %d\n", getGoroutineID(), i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Printf("Гоурутина 2 (верх) - ID: %d, Value: %d\n", getGoroutineID(), i)
		}
		wg.Done()
	}()
	wg.Wait()
}

func getGoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	id := -1
	fmt.Sscanf(string(buf[:n]), "Гоурутина %d", &id)
	return id
}
