package main

import (
	"fmt"
//	"runtime/pprof"
//	"os"
)
/*
func NewCounter() {
	n := 0
	func(n int){
		n += 2
		fmt.Println(n)
	}(n)
	fmt.Println(n)
}
*/


func main() {
//	NewCounter()
//	NewCounter()
//	NewCounter()
/*
	if *flagCpuprofile != "" {
    f, err := os.Create(*flagCpuprofile)
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
}
*/
	var a [5]int
	for i := range a {
		defer func(i int) {// pass i and then check
			fmt.Println(i)
		}(i)
		//fmt.Println(i)
}
}


