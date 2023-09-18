package main

import (
	"cpmk/cpmkgo"
	"flag"
	"fmt"
)

var (
	language    *string
	projectName *string
)

func init() {
	language = flag.String("l", "cpp", "language (c, cpp)")
	projectName = flag.String("n", "cpp_project", "project name")
}

func main() {
	flag.Parse()

	err := cpmkgo.SetupProject(*language, *projectName)
	if err != nil {
		fmt.Println(err)
	}
}
