package main
import (
    "fmt"
    "math/rand"
    "runtime"
    "time"
)
func query() int {
    n := rand.Intn(50)
    time.Sleep(time.Duration(n) * time.Millisecond)
    return n
}
func queryAll() int {
    ch := make(chan int)
    go func() { ch <- query() }()
    go func() { ch <- query() }()
    go func() { ch <- query() }()
   	return <-ch //causes leaky G's, consume with separate G's
}
func main() {
    for i := 0; i < 4; i++ {
	fmt.Println(queryAll())
	    //fmt.Println(a)
        fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
    }
}
