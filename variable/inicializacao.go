package variable

import "fmt"

var i, j int = 1, 2

func inicializacao() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
