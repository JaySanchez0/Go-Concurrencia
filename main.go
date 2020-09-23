package main

import (
	"sync"
	"strconv"
	"fmt"
)

var cont int = 0
var wg sync.WaitGroup

func par(i int, j int){
	fmt.Println("Start")
	for k:=i;k<=j;k++{
		if k%2==0 {
			fmt.Println(strconv.Itoa(k))
		}
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go par(0,100)
	go par(101,200)
	wg.Wait()
	fmt.Println("End!")
}