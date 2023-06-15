package main

import (
	"fmt"
	"runtime"
)

func printChanMessage(message chan string) {
	fmt.Println(<-message)
}

func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan string)

	for _, name := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			data := fmt.Sprintf("Hello %s", who)
			message <- data
		}(name)
	}

	for i := 0; i < 3; i++ {
		printChanMessage(message)
	}
}
