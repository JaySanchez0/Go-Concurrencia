package main

import (
	"sync"
	"strconv"
	"fmt"
	"time"
)

var cond *sync.Cond

var list []int

var wg sync.WaitGroup

func consume(){
	for{
		fmt.Println("<<<< Consume")
		cond.L.Lock()
		if len(list)== 0 {
			cond.Wait()
		}else{
			var l int= list[len(list)-1]
			list = list[:len(list)-1]
			fmt.Println(strconv.Itoa(l))
		}
		cond.L.Unlock()
		time.Sleep(5000 * time.Millisecond)
	}
	wg.Done()
}

func produce(){
	for{
		fmt.Println(">>>>> Produce")
		cond.L.Lock()
		list = append(list,7)
		cond.Signal()
		cond.L.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}


func main(){
	wg.Add(2)
	m:= sync.Mutex{}
	cond = sync.NewCond(&m)
	go consume()
	go produce()
	wg.Wait()
	fmt.Println("end")
}