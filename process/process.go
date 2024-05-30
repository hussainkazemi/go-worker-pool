package process

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DataProcess(n int) int {
	fmt.Printf("data %d start process\n", n)
	time.Sleep(5 * time.Second)
	fmt.Printf("data %d end process \n", n)
	return rand.Intn(n) + 1
}

func DataProcessConcurrent(wg *sync.WaitGroup, ch chan int, n int) {
	fmt.Printf("data %d start process\n", n)
	time.Sleep(5 * time.Second)
	fmt.Printf("data %d end process \n", n)
	wg.Done()
	ch <- rand.Intn(n) + 1
}
