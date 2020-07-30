package main

func UmaFuncao(x int, y int) int {
	return x + y
}

func OutraFuncao(x, y int) int {
	return x + y
}

func MaisUmaFuncao(x, y int) (int, int, int) {
	return y, x, x+y
}

func FuncaoComRetornoNomeado(x, y int) (xx, yy int) {
	xx = y + x
	yy = y - x
	return
}