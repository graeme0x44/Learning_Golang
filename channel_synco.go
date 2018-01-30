package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {

	// Create a channel
	done := make(chan bool, 1)
	// Pass the channel to the worker
	go worker(done)

	// Block until the worker returns done
	<-done

	// Launch new worker
	go worker(done)
	<-done
}
