package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var saldo int64 = 0

// solution1
var isLock bool = false

// solution2
var lock sync.Mutex

func addSaldo(){
	for i :=0; i< 50000; i++{
		atomic.AddInt64(&saldo, 1)
		saldo++
	}
	
}

// solution 1
// func Lock(lockedFunc func()){
// 	for isLock{
// 		if isLock{
// 			break
// 		}
// 	}
// 	isLock = true
// 	lockedFunc()
// 	isLock = false
// }



func main(){

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func ()  {
		defer wg.Done()
		// solution1
		// Lock(addSaldo)

		// solution2 
		// lock.Lock()
		// addSaldo()
		// lock.Unlock()

		addSaldo()
	}()
	
	go func ()  {
		defer wg.Done()
		// solution1
		// Lock(addSaldo)

		// solution2 
		// lock.Lock()
		// addSaldo()
		// lock.Unlock()
		addSaldo()
	}()

	wg.Wait()
	fmt.Println(saldo)
}