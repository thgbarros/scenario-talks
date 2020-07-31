package main

import "fmt"

func exemploSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[:len(primes)/2])
	fmt.Println(primes[len(primes)/2:])
	fmt.Println(primes)
}

func exemploMake() {
	a := make([]int, 5)
	fmt.Printf("%s len=%d cap=%d %v\n",
		a, len(a), cap(a), a)

	b := make([]int, 0, 5)
	fmt.Printf("%s len=%d cap=%d %v\n",
		b, len(b), cap(b), b)

	c := b[:2]
	fmt.Printf("%s len=%d cap=%d %v\n",
		c, len(c), cap(c), c)

	d := c[2:5]
	fmt.Printf("%s len=%d cap=%d %v\n",
		d, len(d), cap(d), d)
}

func exemploAppend() {
	b := make([]int, 0, 5)
	fmt.Printf("%s len=%d cap=%d %v\n",
		b, len(b), cap(b), b)

	b = append(b, 0, 1, 2, 3, 4)
	fmt.Printf("%s len=%d cap=%d %v\n",
		b, len(b), cap(b), b)

	b = append(b, 0, 1, 2, 3, 4)
	fmt.Printf("%s len=%d cap=%d %v\n",
		b, len(b), cap(b), b)
}
