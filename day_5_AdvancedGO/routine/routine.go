package main

import (
	"fmt"
	"sync"
)

func main() {
	// go numbers()
	// go alphabets()
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func ()  {
		defer wg.Done()
		fmt.Println("Hello")
	}()

	go func ()  {
		defer wg.Done()
		fmt.Println("World")
	}()

	go func ()  {
		defer wg.Done()
		fmt.Println("Adam")
	}()

	wg.Wait()

}

func numbers() {
	for i:=1 ; i<=5; i++{
		fmt.Println(i)
	}
}

func alphabets() {
	for i:='a' ; i<='e'; i++{
		fmt.Println(i)
	}
}