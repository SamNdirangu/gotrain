package channels

import (
	"fmt"
	"sync"
	"time"
)

//Channels can be used to send any type of data even structs!
///type dummyChannelData struct {
///	typeOfStruct    string
///	message         string
///	someDummyNumber int
///}

func Channels() {

	unbufferedChannels()
	bufferedChannels()

	example1()
	example2()
}

func unbufferedChannels() {
	// Unbuffred channels have no space to store data and act as portals connecting the sender and consumer
	// This ultimately means they have to be executed at the same time and in different go routines

	//create our channel
	dataChannel := make(chan string)

	//start a go routine on another thread
	go func() {
		//Send data to our channel
		dataChannel <- "Hello this is string sent to the channel"
	}()
	//Consume the data sent via the channel by our go routine which is executing on another thread on this function
	dataFromChannel := <-dataChannel
	fmt.Println(dataFromChannel)
}

func bufferedChannels() {
	// Buffered channels have space to store data and go around the portal behavior of unbuffered channels.
	// This technically means one can send and consume data from a buffered channel in one thread

	//create our channel, we add number of items to be buffered at the edn
	dataChannel := make(chan string, 2)

	//start a go routine on another thread
	//Send data to our channel
	dataChannel <- "Hello this is string sent to the channel via our main thread"
	dataChannel <- "Hello this is second string sent to the channel via our main thread"
	go func() {
		//Send data to our channel
		dataChannel <- "Hello this is string sent to the channel via go routine"
	}()
	//Consume the data sent via the channel by our go routine which is executing on another thread on this function
	dataFromChannel := <-dataChannel
	fmt.Println(dataFromChannel)
	//Consume again second string
	dataFromChannel = <-dataChannel
	fmt.Println(dataFromChannel)
	//Consume again string from goroutine
	dataFromChannel = <-dataChannel
	fmt.Println(dataFromChannel)

}

func example1() {
	dataChannel := make(chan int, 1)
	go func() {
		for i := 0; i < 20; i++ {
			dataChannel <- i
		}
		close(dataChannel)
		//By closing the channel when complete sending data, we prevent a deadlock happening
	}()
	// Without closing our channel the range function will still try to access data even after our go routine finishes
	// execution and stops sending data thus leading to deadlock, by closing the channel we notify the consumer on the main thread
	// that no more data is coming via the channel and hence stop waiting for data evading the deadlock
	for n := range dataChannel {
		fmt.Println("N is ", n)
	}
}

func example2() {
	//follow up of example 1 but simulating a go routine that may take time to execute
	dataChannel := make(chan int)

	go func() {
		waitgroup := sync.WaitGroup{}
		for i := 0; i < 20; i++ {
			//introduce some delay pretending to do some hardwork
			// we can run it directly but will take long we can create goroutines to do each of the hard work and work simultaneously

			// in its default form without a waitgroup, the loop may complete long before any of the go routines complete and send data to the channel
			// once the loop completes the channel is closed and the main function proceeds without consuming anydata from our channel
			// to prevent this is why we use wait group to pause execution at some point until our goroutines complete execution
			waitgroup.Add(1)
			go func(i int) {
				// A defer statement defers the execution of a function until the surrounding function returns.
				defer waitgroup.Done()
				// Waitgroup done will only be called once our go func completely executes are returns
				result := doHardWork(i)
				dataChannel <- result
			}(i)
		}
		waitgroup.Wait()
		close(dataChannel)
	}()
	for n := range dataChannel {
		fmt.Println("N is ", n)
	}

}

///Complementary functions
func doHardWork(i int) int {
	//simulate a slow response of 1 second
	time.Sleep(time.Second)
	//return some random number
	return i
}
