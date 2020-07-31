package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func exemploMap() {
	m = make(map[string]Vertex)
	m["Scenario DEV"] = Vertex{
		-21.995005, -47.899460,
	}
	fmt.Println(m)
}
