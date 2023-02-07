package main

import (
	"fmt"
	"time"
)

func mainChan() {
	messChanel1 := make(chan string)
	messChanel2 := make(chan string)

	go func() {
		for {
			messChanel1 <- "Write to chanel 1 200 ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {

			messChanel2 <- "Write to chanel 2 1 sec"
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case msg := <-messChanel1:
			fmt.Println(msg)
		case msg := <-messChanel2:
			fmt.Println(msg)
		}

		// go readChan(messChanel1)
		// go readChan(messChanel2)

	}

}

func readChan(mess chan string) {
	fmt.Println(<-mess)
}
