package main

import (
	"fmt"

	"github.com/cetinboran/goarg/cla"
)

func main() {
	a := cla.Init()
	a.SetTitle("Boran")
	a.SetExamples([]string{"go run main.go -h 127.0.0.1 -p 22", "go run main.go -p 192.168.1.*"})

	a.SetUsage("Blabla bla blaaaa bla")
	a.AddOption("-h,--host", false, "Enter your host.", []string{"Invalid"})
	a.AddOption("-p,--port", false, "Enter your Port.", []string{"Invalid"})
	a.AddOption("--code", true, "Open With VSCode.", []string{"Invalid"})
	args := a.Start()

	fmt.Println(args)

}
