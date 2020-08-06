package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//func main() {
//c := make(chan int)
//quit := make(chan int)
//go func() {
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-c)
//	}
//	quit <- 0
//}()

//fibonacci(c, quit)
//concorrenciaEmGrupo()
//}

func timerETicker() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func concorrenciaEmGrupo() {
	alphabet := []string{"A", "B", "C", "D"}
	var wg sync.WaitGroup

	for index, letter := range alphabet {
		if wgDelta := index % 2; wgDelta == 0 {
			wg.Add(wgDelta)
		} else {
			wg.Add(wgDelta)
		}

		go regiaoCritica(letter, &wg)
	}

	fmt.Println("Esperando....")
	wg.Wait()
	fmt.Println("FIM WG.")
}

var mutex = &sync.Mutex{}

func regiaoCritica(regiaoName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		fmt.Printf("Concorrência %s - %d\n", regiaoName, i)
		time.Sleep(1 * time.Millisecond)
		mutex.Unlock()
	}
}
