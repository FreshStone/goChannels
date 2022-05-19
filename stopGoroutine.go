package main

import (
	"fmt"
//	"time"
)

func Generator() chan int {
        ch := make(chan int)
        go func() {
                n := 5
                for {
                        select {
                          case ch <- n:
                                n++
                          case <-ch:
                                return //brek will cause exit from select but not from for
                                //break
                        }
                }
        }()
        return ch
}

func main() {
        number := Generator()
        fmt.Println(<-number)
	fmt.Println(<-number)
        close(number)
//time.Sleep(1*time.Second)//causes "panic- send on closed channel; better to use context"
        // …
}
