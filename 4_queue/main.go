package main

import (
	"fmt"
	"DA/4_queue/queue"
	"DA/4_queue/usage"
)

func main() {
	queue.OqTest()
	queue.LqTest()
	usage.PqTest()
	usage.DqTest()
	
	fmt.Println()
}