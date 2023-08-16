package main

import (
	"github.com/cetinboran/goarg/structs"
)

func main() {
	a := structs.Goarg{}

	a.SetUsage("Blabla bla blaaaa bla")
	a.AddOption("-h,--host", false, "Enter your host.", []string{"Invalid"})
	a.AddOption("-p,--port", false, "Enter your Port.", []string{"Invalid"})
	a.AddOption("--code", true, "Enter your Port.", []string{"Invalid"})
	a.Start()
}
