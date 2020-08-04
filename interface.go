package main

import "fmt"

type Luminaria interface {
	Ligar()
	Desligar()
}

type LuminariaNaoDim struct{}

func (LuminariaNaoDim) Ligar() {
	fmt.Println("LuminariaNaoDim - Ligada")
}

func (LuminariaNaoDim) Desligar() {
	fmt.Println("LuminariaNaoDim - Desligada")
}


type LuminariaDim struct {
	dim int
}

func (dimmer *LuminariaDim) Ligar() {
	dimmer.dim = 100
	fmt.Printf("LuminariaDim - Ligada %d%s \n", dimmer.dim,"%")
}

func (dimmer *LuminariaDim) Desligar() {
	dimmer.dim = 0
	fmt.Printf("LuminariaDim - Desligada %d%s \n", dimmer.dim, "%")
}

func (dimmer *LuminariaDim) SetIntensidade(intensidade int) {
	dimmer.dim = intensidade
	fmt.Printf("LuminariaDim - Intensidade em %d%s \n", dimmer.dim, "%")
}

//func main() {
//	var luminaria Luminaria
//	luminaria = LuminariaNaoDim{}
//	luminaria.Ligar()
//	luminaria.Desligar()
//
//	luminariaDim := &LuminariaDim{}
//	luminaria = luminariaDim
//	luminaria.Ligar()
//	luminaria.Desligar()
//
//	luminariaDim.SetIntensidade(50)
//
//}


func switchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Luminaria:
		fmt.Printf("%q é uma Luminária\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}




