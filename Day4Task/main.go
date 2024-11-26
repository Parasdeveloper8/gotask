package main

import (
	"fmt"
	"time"
)

func sayHello(ch chan string /*wg *sync.WaitGroup*/) {
	//defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- "hello from channel"
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) //close the channel when done
}
func main() {
	ch := make(chan string)
	//var wg sync.WaitGroup
	//wg.Add(1)
	go sayHello(ch /*&wg*/)
	//wg.Wait()
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println("Goroutine is done")
}
