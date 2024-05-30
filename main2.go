package main

import (
	"fmt"
	"worker/data"
	"worker/process"
)

func main() {
	taskList := data.ReadData()
	var answers []int
	for _, task := range taskList {

		answer := process.DataProcess(task)
		answers = append(answers, answer)
	}

	fmt.Println(answers)

}
