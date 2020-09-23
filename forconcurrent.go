package main

import (
	"sync"
	"fmt"
	"strconv"
)

var wg sync.WaitGroup

func printInterval(a int, b int){
	for i:=a;i<b;i++{
		fmt.Println(strconv.Itoa(i))
	}
	wg.Done()
}


func main(){
	for i:=0;i<10;i++{
		wg.Add(1)
		go printInterval(i*10,(i+1)*10)
	}
	wg.Wait()
}