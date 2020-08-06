package main

import (
	"fmt"
	"scenario-talks/pack"
	"scenario-talks/pack/innerPack"
)

func main() {
	fmt.Printf("Hello world")

	pack.naoExportavel()
	pack.Exportavel()
	pack.OutroExportavel()
	innerPack.InnerExportavel()
}
