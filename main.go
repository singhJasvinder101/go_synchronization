package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	
	done <- true // This must block if no further receiver is ready
	// if the channel is unbuffered, the sender will block if the receiver
	// is ready to receive the value. This is useful for synchronizing
	// goroutines, as it allows one goroutine to wait for another to finish
	// before continuing. In this case, the sender is the worker goroutine
	// and the receiver is the main goroutine. The main goroutine will block
	// until the worker goroutine sends a value on the done channel, indicating
	// that it has finished its work. This is a common pattern in Go for
	// synchronizing goroutines and ensuring that they complete before the
	// program exits.
	
	// If the channel is buffered, the sender will block until there is space	
	// in the channel to send the value. This is useful for controlling the
	// flow of data between goroutines, as it allows one goroutine to send
	// data to another without blocking, as long as there is space in the
	// channel. In this case, the sender is the worker goroutine and the
	// receiver is the main goroutine. The main goroutine will block until
	// the worker goroutine sends a value on the done channel, indicating
	// that it has finished its work. This is a common pattern in Go for
	// synchronizing goroutines and ensuring that they complete before the
	// program exits.
	
	// If the channel is closed, the sender will panic. This is useful for
	// indicating that the channel is no longer available for sending data,
	// and that any goroutines that are trying to send data on the channel
	// should stop. In this case, the sender is the worker goroutine and
	// the receiver is the main goroutine. The main goroutine will panic if
	// the worker goroutine tries to send a value on the done channel after
	// it has been closed. This is a common pattern in Go for indicating
	// that a channel is no longer available for sending data, and that any
	// goroutines that are trying to send data on the channel should stop.
	// This is a common pattern in Go for synchronizing goroutines and
	// ensuring that they complete before the program exits.

	// If the channel is nil, the sender will block forever. This is useful
	// for indicating that the channel is not available for sending data,
	// and that any goroutines that are trying to send data on the channel
	// should stop. In this case, the sender is the worker goroutine and
	// the receiver is the main goroutine. The main goroutine will block
	// forever if the worker goroutine tries to send a value on the done
	// channel. This is a common pattern in Go for indicating that a channel
	// is not available for sending data, and that any goroutines that are
	// trying to send data on the channel should stop. This is a common
	// pattern in Go for synchronizing goroutines and ensuring that they
	// complete before the program exits.

	fmt.Println("done")
}

func main() {
	done := make(chan bool) // unbuffered

	var wg sync.WaitGroup
	// wg.Add(2)
	wg.Add(3) 	// wait for > recievers

	for i := 0; i < 5; i++ {
		go func ()  {
			worker(done) 
			wg.Done()
		}()
	}

	<-done 
	<-done 

	// TWIST:
	// IF I use waitgroup for this channel to wait for 6 goroutines
	// then I will get error 
	
	wg.Wait() // wait for specified number of goroutines to finish


	// print 2	times "done" only becoz receiver nhi ha to 3,4,5th
	// senders will block


	// for i := 0; i < 10; i++ {
	// 	// launching 10 goroutines
	// 	// But they all share the same variable i (which is declared 
	// 	// outside the goroutine) so they may conflic or either use
	// 	// mutex (instead of blocking it queues the goroutines interfering for synchronization)

	// 	go fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	// calling an immediately-invoked function (i) inside the 
	// 	// goroutine. This function receives the current i as a 
	// 	// copy (n), passed by value. So each goroutine prints its
	// 	//  own copy of the loop variable i.


	// 	go func(n int) {
	// 		fmt.Println(n)
	// 	}(i)
	// }

	// ch := make(chan int)

	// select {
	// case <-ch: // blocked forever
	// 	fmt.Println("Received")
	// }



}