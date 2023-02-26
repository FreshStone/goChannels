package main

import (
	"fmt"
	"time"
)

func Generator() (chan int, chan bool) {
        ch := make(chan int)
	closechan := make(chan bool)
        go func() {
                n := 5
                for {
                        select {
                          case ch <- n:
                                n++
                          case <-closechan:
				fmt.Println("closing oprations")
                                return //break will cause exit from select but not from for
                                //break
			 default:
				 fmt.Println("waiting")
                        }
                }
        }()
        return ch, closechan
}

func main() {
        number, closechan := Generator()
	time.Sleep(1*time.Millisecond)
        fmt.Println(<-number)
	fmt.Println(<-number)
        closechan <- true
	close(number)
	time.Sleep(1*time.Second)//causes panic( send on closed channel) if number channel not closed gracefully; better to use context/closechan as used above
        // …
}
