package main

import (
	"fmt"
	"time"
)

func main(){
//	var ch chan int
//	c := make(chan, int)
	a := [][]int{{145,2546,3345,43} ,{5,6,7}, {1,2,3,4,9}}
	//a := []int{1,2,3,4,5}
	sum := 0
	for i := range a {
		fmt.Println(a[i])
		time.Sleep(1*time.Millisecond)
		go func(i int) {   //shared variable i is sent as param to avoid DATA RACE
//			fmt.Printf("length : %d\n", len(a[i])) //  DATA RACE
//			time.Sleep(1*time.Millisecond)
			for j := range a[i] {
//			fmt.Printf("length : %d\n", len(a[i]))
sum += a[i][j] //DATA RACE 
			fmt.Println(a[i][j])
			}
		}(i)
	}
time.Sleep(10*time.Millisecond)
fmt.Println(sum)
}


