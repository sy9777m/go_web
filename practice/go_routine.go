package practice

import (
	"fmt"
	"time"
)

func GoRoutine() {
	fmt.Println("main function start", time.Now())

	done := make(chan bool)

	go long(done)
	go short(done)

	<- done
	<- done
	fmt.Println("main function end", time.Now())
}

func long(done chan bool) {
	fmt.Println("long function start", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long function end", time.Now())
	done <- true
}

func short(done chan bool) {
	fmt.Println("short function start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short function end", time.Now())
	done <- true
}
