package main

import (
	"fmt"
	"sync"
	"worker/data"
	"worker/process"
)

func main() {
	taskList := data.ReadData()
	var answers []int
	ch := make(chan int, len(taskList))
	wg := sync.WaitGroup{}
	wg.Add(len(taskList))
	for _, task := range taskList {
		go process.DataProcessConcurrent(&wg, ch, task)

	}
	go func() {
		for i := 0; i < len(taskList); i++ {
			answers = append(answers, <-ch)
		}
	}()

	wg.Wait()
	fmt.Println(answers)
}
