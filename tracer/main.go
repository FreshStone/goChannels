package main

import (
	"log"
	"os"
	"runtime/trace"
//	"time"
)

func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()

	const n = 3

	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go pass(left, right)
		left = right//pass by reference
	}

	go sendFirst(right)
	log.Println(<-leftmost)
//	time.Sleep(1*time.Second)
}

func pass(left, right chan int) {
	v := 1 + <-right
	left <- v
}

func sendFirst(ch chan int) { ch <- 0 }
