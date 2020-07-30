package main

import "fmt"

func loopComFor() {
	for i := 0; i< 10; i++ {
		fmt.Println(i)
	}
}

func loopComWhile() {
	i := 0
	for i< 10 {
		fmt.Println(i)
		i++
	}
}

func loopInfinito() {
	for {
		fmt.Println("infinito")
	}
}

func loopComForEach() {
	valores := []byte{0,1,2,3}
	for valor, index := range valores {
		fmt.Printf("Index: %d, valor: %d", index, valor)
	}

	for valor, _ := range valores {
		fmt.Printf("valor: %d", valor)
	}

	for _, index := range valores {
		fmt.Printf("index: %d", index)
	}
}