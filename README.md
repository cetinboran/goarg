# GOARG

## What is This?
+ This project is for those who use command line arguments in their own projects.
+ It's easy to use you can implement your project just 4 lines of code.

## What can you do?
+ In just one line, you can enter the name of your option, the activity value, the usage information, and the error specific to it.
+ You can add the project title, examples of how it works. When added, they will all be displayed with auto usage writer.
+ This project does things for you that you don't want you to do. You just focus on your own project.

## How to Install?
+ Open your project and write the code below in the terminal. The project will be added.
    + `go get github.com/cetinboran/goarg@v1.x.x`

## Structs
+ I will talk about the structs that should be known here.
    + Let's start with the **Goarg Struct** first
    + The "goarg" struct can accept a total of 8 values: Title, Usage Examples, Description, Options, Mods, ModeName, and main.
    + The values for Title, Usage Examples, and Description are set and used by you using the .Set functions, so there is no need to go into details about them.
    + However, you need to have an understanding of the other fields:
        + **Options**: All the options you add are stored in an array-like structure known as the Options struct. You can refer to where I previously explained this struct for a better understanding.
        + **Mods**: This field is a combination of the goarg structure you're describing and the ModeName. When making a goarg.AddMode call, the first string argument you provide is used here as a key. This field becomes useful when searching for a specific mode.
        + **ModeName**: This field is a value you set when calling the AddMode function. It holds significant importance as it helps you properly direct incoming inputs. It's returned to you after processing within the Input struct. With this field, you can easily use a switch case structure to direct incoming inputs to their respective modes.
    + Let me explain the **Option Struct** for you.
    + The Option struct accepts 4 values, and understanding each of them will significantly simplify your task:
        + **PlaceHolder**: When using AddOption, the first argument you add becomes the PlaceHolder. This allows you to enter the argument through the Command Line. You can separate multiple placeholders using commas. Additionally, they must begin with a '-' or '--' to be recognized.
        + **Active**: This value indicates whether the command line argument expects additional input. When set to false, it expects input. For instance, after typing '-p', you would need to provide another value separated by a space. When set to true, it only returns one value. You can use this as 'true' when appropriate.
        + **Usage**: Here, you can describe the purpose of your Option. In Automatic Usage, this description will be visible.
        + **Global**: This determines whether the option is global across all modes or specific to a single mode. This distinction can be extremely useful.
    +  Finally let me explain **Input Struct** for you.
    + The Input struct accepts 3 values.
        + Argument: This field specifies the source placeholder from which the input originates. Consider a scenario where you're working on a project and receive 4 or 5 different inputs. To prevent confusion, you can organize and categorize incoming inputs based on their associated Argument labels.
        + Value: This field represents the value sent by the user and is returned as an string. It's a versatile data type that can hold different types of values, depending on the input provided by the user.
        + ModeName: As I previously explained in the context of the goarg struct, this field indicates which mode's option the input corresponds to. It helps you associate the input with the relevant mode.
        When we write `args := Setup.Start()`, the variable args receives the incoming values as a slice of Input structs, i.e., []Input.

## How to Use?
+ First, let me explain the methods.
    + `Goarg := cla.Init()`: Initializes the struct
    + `Goarg.SetTitle(string)`: Addes the name of the project to the usage
    + `Goarg.SetExamples(string)`: Adds examples of how your project works to the usage.
    + `Goarg.SetUsage(string)`: You can set your own usage 
    + `Goarg.AddOption(args string, active bool, usage string, error []string)`: Addes a option to the goarg for later to use.
        + args: the option name like -h --host. You must separate the setting names with (,).
        + active: Indicates whether the setting takes input. if true it takes no input it's just a "1"
        + usage: this is just a usage for the option. You can see on automatic usage.
        + error: you can add errors to the option for later to use.
    + `Goarg.AutomaticUsage()`: You can activate the automatic usage. This code MUST be after AddOptions.
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
    Goarg := cla.Init()
    Goarg.SetTitle("cetinboran")
    Goarg.SetExamples([]string{"go run main.go -h 127.0.0.1 -p 22", "go run main.go -p 192.168.1.*"})

    Goarg.AddOption("-h,--host", false, "Enter your host.", []string{"Please enter a ip not a domain."})
    Goarg.AddOption("-p,--port", false, "Enter your Port.", []string{"Enter less than 65535"})
    Goarg.AddOption("--fast", true, "Makes it faster.", []string{"If you dont have enough ram. It can be slower than faster."})
    Goarg.AutomaticUsage()

    args := Goarg.Start()

    for _, arg := range args {
		fmt.Println(arg.Argument, arg.Value, arg.Error)
	}
}
```
+ If we write the following code to the terminal `go run .\main.go -h 127.0.0.1 --port 22 --fast`.
+ we can see the result below.
    + h 127.0.0.1 [Please enter a ip not a domain.]
    + port 22 [Enter less than 65535]
    + fast 1 [If you dont have enough ram. It can be slower than faster.]

+ Below you can see the auto-generated help. You can see this message with the code below.
    + `go run .\main.go --help`

```
    cetinboran
    Enter your host. [-h --host]     
    Enter your Port. [-p --port]        
    Makes it faster. [--fast]

    Examples:
    1. go run main.go -h 127.0.0.1 -p 22
    2. go run main.go -p 192.168.1.* 
```

## Errors
+ there are four error messages you may encounter
    + not an option -> If you didn't add such a option but you are using it, this message will appear.
    + missing input for ... -> If you have added the option with the false activity option but forgot to add the input, this message will appear.
    + no need input for ... ->  If you added the option with true activity option and add the input, this message will appear. ( You don't need the input )
    + you don't need more than one of this option -> If you try to use the same option more than once, this message will appear.

## What's new
+ 19.08.2023
    + now you can add mods to your project
    + it allows you to spread the options in your project more regularly.
    + Your project can be more organized.
    + First, let's explain the added methods.
        + `hello := cla.ModInit()`: Initializes the mode struct
        + `Goarg.AddMode("helloMod", &hello)`: This method adds a mod to the main goarg struct.
            + The first argument specifies how to call the mode, the other specifies the mod's struct
            + You must be the address of that struct when using the other argument.
    + The rest is the same. You can add settings, add examples, add title as in the main struct.
    + At the end just run the start of the main struct
    + Reminder: Automatic Usage code must be at the end.
    + With the code on the below. We access the mode you added with AddMode() and we can use it just like we did in main goarg.
    + `go run .\main.go helloMod`
    + After typing the first argument you give in the addMode function to the terminal, access to the functions of that mode is provided.

```
package main

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func main() {
	Goarg := cla.Init()

	Goarg.SetTitle("cetinboran")
	Goarg.SetExamples([]string{"go run main.go -h 127.0.0.1 -p 22", "go run main.go -p 192.168.1.*"})

	Goarg.AddOption("--host", false, "Enter your host.", []string{"Please enter a ip not a domain."})
	Goarg.AddOption("-p,--port", false, "Enter your Port.", []string{"Enter less than 65535"})
	Goarg.AddOption("--fast", true, "Makes it faster.", []string{"If you dont have enough ram. It can be slower than faster."})

	hello := cla.ModInit()
	Goarg.AddMode("helloMod", &hello)

	hello.SetTitle("Hello Mod")
	hello.SetExamples([]string{"Example 1", "Example 2"})
	hello.AddOption("--hello", true, "Says hello", []string{"Error msg 1"})
	hello.AddOption("-sc,--scream", false, "You can scream!", []string{"This is an error"})

	hello.AutomaticUsage()
	Goarg.AutomaticUsage()

	args := Goarg.Start()

	for _, arg := range args {
		fmt.Println(arg.Argument, arg.Value, arg.Error)
	}
}
```   
+ 20.08.2023
    + Now we can understand from which mode the input came from, from the return of the arguments.
    + args.ModeName carries this information.
    + This information is the first input value we entered in the AddMode section. Accordingly, it became easier to process incoming inputs.
    + AddGlobalOption func added. Now, when you define a option with this function in the main goarg, this option will come to all the modes below.
    + This code must be under AddMode() Function
+ 23.08.2023
    + Fixed 4 medium bugs.
    + `Goarg.SetDescription(string)`: Added setdescription function. You can now write a description about your goarg.It will be added directly thanks to automatic usage.
+ 25.08.2023
    + Now we can add mods inside mods. like 'config create'
    + A few minor bugs fixed






# Contact

[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='40'>](https://github.com/cetinboran)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='40'>](https://www.linkedin.com/in/cetinboran-mesum/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/instagram.svg' alt='instagram' height='40'>](https://www.instagram.com/2023an_m/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/twitter.svg' alt='twitter' height='40'>](https://twitter.com/2023anM)  
