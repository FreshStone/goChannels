package main

import (
  "fmt"
  "time"
)

func main(){
  var v int
  var isChOpen bool
  ch := make(chan int, 3)
  go func(){
      for i := 1; i < 4; i++{
        ch <- i
        //fmt.Println(i)
      }
      time.Sleep(1*time.Second)
      close(ch)
  }()
  /*
  for v := range ch{
    //what if ch(buffered) still has some unread values ??
    time.Sleep(1*time.Second)
    fmt.Println(v)
  }*/

  for{
    select{
      case v, isChOpen = <-ch:
	fmt.Println("d")
        time.Sleep(2*time.Second)//simulating consumer delay
        fmt.Println(v, isChOpen)
	if !isChOpen{
		fmt.Println("exiting")
		return //why "break" doesn't works here???
	}
     default:
	     fmt.Println("in default")
	     break
    }
  }

  /*
   time.Sleep(1*time.Second)
   v, isChOpen  = <-ch
   fmt.Println(v, isChOpen)
   v, isChOpen = <-ch
   fmt.Println(v, isChOpen)
   v, isChOpen = <-ch
   fmt.Println(v, isChOpen)*/

}
