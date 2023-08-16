package main

import (
	"fmt"

	"github.com/cetinboran/goarg/cla"
)

func main() {
	a := cla.Init()

	a.SetTitle("cetinboran")
	a.SetExamples([]string{"go run main.go -h 127.0.0.1 -p 22", "go run main.go -p 192.168.1.*"})
	//a.SetUsage("Blabla bla blaaaa bla")

	a.AddOption("-h,--host", false, "Enter your host.", []string{"Please enter a ip not a domain."})
	a.AddOption("-p,--port", false, "Enter your Port.", []string{"Enter less than 65535"})
	a.AddOption("--code", true, "Open With VSCode.", []string{"Make sure that vscode is added to the path."})
	a.AutomaticUsage() // This has to be in last

	args := a.Start()

	for _, i2 := range args {
		fmt.Println(i2.Argument, i2.Value, i2.Error)
	}

}
