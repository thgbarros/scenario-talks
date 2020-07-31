package main

import "fmt"

var a [10]int

func exemploArray() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(cap(primes))
	fmt.Println(len(primes))

	a[5] = 510
	fmt.Println(a)

	fmt.Println(primes)
}
