package main

import (
	"context"
	"fmt"
	"runtime"
	"reflect"
	"time"
)

/*
// leaking G's
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
//		n := 0
		for {
			ch <- n
			n = n+2
		}
	}()
//	fmt.Printf("%d\n", reflect.TypeOf(ch).Size())
	fmt.Println(reflect.TypeOf(ch))

	return ch
}

func main() {
//	runtime.StartTrace()
for n := range gen() {
    fmt.Println(n)
    fmt.Printf("number of G's : %d\n", runtime.NumGoroutine())
	if n == 6 {
        break
    }
}
fmt.Printf("number of G's : %d\n", runtime.NumGoroutine())
//runtime.StopTrace()
}

*/


//Correction

func gen(ctx context.Context)chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	//	fmt.Printf("%d\n", reflect.TypeOf(ch).Size())
	}()
	fmt.Printf("%d\n", reflect.TypeOf(ch).Size())
	return ch
}

func main() {
//	runtime.StartTrace()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		fmt.Printf("number of G's : %d\n", runtime.NumGoroutine())
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1*time.Millisecond) //causes other routine to return so #G's = 1
	fmt.Printf("number of G's : %d\n", runtime.NumGoroutine())// #2
//	runtime.StopTrace()
}

