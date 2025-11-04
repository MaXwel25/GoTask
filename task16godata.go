package main


import (
	"fmt"
	"math/rand"
	"sync"
)

var numbers []int

var mutex sync.Mutex

func main(){
	const grs = 3

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
	}

	wg.Wait()

	for i, number := range numbers {
		fmt.Println(i, number)
	}

}

func random(amout int) {
	for i := 0; i < amout; i++ {
		n := rand.Intn(100)

		mutex.Lock()
		numbers = append(numbers, n)
		mutex.Unlock()
	}
}
