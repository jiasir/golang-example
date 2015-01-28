// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// License: The MIT License
	
package main

import "fmt"
import "time"

// The function run in goroutine
// The done channel will be used to notify another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	
	done <- true	// Send a value to notify that we're done
}

func main() {
	// Start a worker goroutine
	done := make(chan bool, 1)
	go worker(done)
	
	<- done // Block until we receive a notification from the worker on the channel
}
