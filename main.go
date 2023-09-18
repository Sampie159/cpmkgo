package main

import (
	"cpmk/cpmkgo"
	"fmt"
)

func main() {
	err := cpmkgo.SetupProject("cpp", "test")
	if err != nil {
		fmt.Println(err)
	}
}
