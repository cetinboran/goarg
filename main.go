package main

import (
	"fmt"

	"github.com/cetinboran/goarg/cla"
)

func main() {
	a := cla.Init()

	a.SetUsage("Blabla bla blaaaa bla")
	a.AddOption("-h,--host", false, "Enter your host.", []string{"Invalid"})
	a.AddOption("-p,--port", false, "Enter your Port.", []string{"Invalid"})
	a.AddOption("--code", true, "Open With VSCode.", []string{"Invalid"})
	args := a.Start()

	fmt.Println(args)
}
