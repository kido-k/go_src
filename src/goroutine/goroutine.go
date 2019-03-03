package main

import (
	"fmt"
	"sync"
	"time"
)

func printOneSecNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d", i)
	}
}

func printNumbers1() {
	for i := 0; i < 10; i++ {
		// fmt.Printf("%d", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		// fmt.Printf("%c", i)
	}
}

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%c", i)
	}
	wg.Done()
}

func printNumbers3(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	w <- true
}

func printLetters3(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	w <- true
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

// func print2() {
// 	printNumbers2()
// 	printLetters2()
// }

// func goPrint2() {
// 	go printNumbers2()
// 	go printLetters2()
// }

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func callerA(c chan string) {
	// time.Sleep(1000 * time.Millisecond)
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	// time.Sleep(2000 * time.Millisecond)
	c <- "Hola Mondo!"
	close(c)
}

func main() {
	// printOneSecNumbers()

	// var wg sync.WaitGroup
	// wg.Add(2)
	// go printNumbers2(&wg)
	// go printLetters2(&wg)
	// wg.Wait()

	// w1, w2 := make(chan bool), make(chan bool)
	// go printNumbers3(w1)
	// go printLetters3(w2)
	// fmt.Println(<-w1)
	// fmt.Println(<-w2)

	// c := make(chan int, 2)
	// go thrower(c)
	// go catcher(c)
	// time.Sleep(100 * time.Millisecond)

	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(1 * time.Microsecond)
	// 	select {
	// 	case msg := <-a:
	// 		fmt.Printf("%s from A\n", msg)
	// 	case msg := <-b:
	// 		fmt.Printf("%s from B\n", msg)
	// 	default:
	// 		fmt.Println("Default")
	// 	}
	// }

	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}

	// fmt.Printf("%s from A\n", <-a)
	// fmt.Printf("%s from B\n", <-b)
}
