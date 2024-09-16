package main

import (
	"fmt"
	"net/http"
	"time"
)

// Channels
// channel <- 5 : Send the value '5' into this channel
// myNumber <- channel : Wait for a value to be sent into the channel. When we get one, assing the value to 'myNumber'
// fmt.Println(<- channel) : Wait for a value to be sent into the channel. When we get one, log it out immediately.

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Channels allows you to communicate data between diferent
	// routines therefore they are tight to one single type of data.
	c := make(chan string)

	for _, link := range urls {
		go checkLink(link, c)
	}

	// for {
	// 	// Receiving message from a channel is blocking code
	// 	go checkLink(<-c, c)
	// }

	// Alternative syntax
	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
		// Never try to access same variable from a different go routine
		// Always pass by value and prevent error whan var changed
		// outside of the scope
	}

	// An exta message receiver is going to hang up the program
	// fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
