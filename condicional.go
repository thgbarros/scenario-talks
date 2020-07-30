package main

import (
	"fmt"
	"math"
	"runtime"
)

func ifComum(x uint) {
	if x < 0 {
		fmt.Println("Há,Pegadinha do malandro!")
	}
}

func ifComVariavelDeEscopoLocal(x uint, y uint) {
	if xy := math.Pow(float64(x), float64(y)); xy < 20 {
		fmt.Print(xy)
	}
}

func switchEmGo() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case funcaoParaExemplificarOrdenacaoDoSwitch(os):
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func funcaoParaExemplificarOrdenacaoDoSwitch(so string) string{
	fmt.Println("na função é: " + so)
	return so
}

func switchSemCondicao() {
	os := runtime.GOOS
	switch {
	case os == "darwin":
		fmt.Println("OS X.")
	case os == funcaoParaExemplificarOrdenacaoDoSwitch(os):
		fmt.Println("(switchSemCondicao) Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	switchSemCondicao()
}

