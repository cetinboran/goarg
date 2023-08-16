package main

import (
	"fmt"

	"github.com/cetinboran/goarg/structs"
)

func main() {
	a := structs.Goarg{}

	a.SetUsage("Blabla bla blaaaa bla")
	a.AddOption("-h,--host", false, "Enter your host.", []string{"Invalid"})

	fmt.Println(a.Options)
}
