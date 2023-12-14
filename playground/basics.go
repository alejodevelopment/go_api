package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	myMap := make(map[string]int) // dict
	myMap["6"] = 6
	fmt.Println(myMap)
	s := []int{1, 2, 3} // slices (lists, arrays)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 14, 45, 12)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	//c := make(chan int) // channel
	//go doSomething(c) // implement channel with go routine
	//<-c
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hola 3 seconds after")
	c <- 1
}
