# GOARG

## What is This?
+ This project is for those who use command line arguments in their own projects.
+ It's easy to use you can implement your project just 4 lines of code.

## What can you do?
+ In just one line, you can enter the name of your setting, the activity value, the usage information, and the error specific to it.
+ You can add the project title, examples of how it works. When added, they will all be displayed with auto usage writer.
+ This project does things for you that you don't want you to do. You just focus on your own project.

## How to Install?
+ Open your project and write the code below in the terminal. The project will be added.
+ `go get github.com/cetinboran/goarg`

## How to Use?
+ [!NOTE] First, let me explain the methods.
    + `Goarg := cla.Init()`: Initializes  the struct
    + `Goarg.SetTıtle(string)`: Addes the name of the project to the usage
    + `Goarg.SetExamples(string)`: Adds examples of how your project works to the usage.
    + `Goarg.SetUsage(string)`: You can set your own usage 
    + `Goarg.AddOption(args string, active bool, usage string, error []string)`: Addes a option to the goarg for later to use.
        + args: is the option name like -h --host
        + active: Indicates whether the setting takes input. if false it takes no input
        + usage: this is just a usage for the option. You can see on automatic usage.
        + error: you can add errors to the option for later to use.
    + `Goarg.AutomaticUsage()`: You can activate the automatic usage. This code should be after AddOption.
    + `Goarg.Start(string) []Input`: This is the last piece of code you need. This code returns us array of struct Input. Input has (for now) three variables.
        + Arguments: Here you can see which option is returning.
        + Value: The value of the option that returned.
        + Error: Custom errors that you add for option.
    + This is it.
+ Now a setup example for the project

```
package main

import (
	"fmt"

	"github.com/cetinboran/goarg/cla"
)
func main(){
    Goarg := cla.Inıt()
    a.SetTitle("cetinboran")
    a.SetExamples([]string{"go run main.go -h 127.0.0.1 -p 22", "go run main.go -p 192.168.1.*"})

    a.AddOption("--host", false, "Enter your host.", []string{"Please enter a ip not a domain."})
    a.AddOption("-p,--port", false, "Enter your Port.", []string{"Enter less than 65535"})
    a.AutomaticUsage()

    args := a.Start()
}
```
