package cla

import (
	"fmt"
	"os"
	"strings"
)

func Init() Goarg {
	return Goarg{Title: "Example Title"}
}

// Using *
func (g *Goarg) SetUsage(usage string) {
	g.Usage = usage
}

func (g *Goarg) SetTitle(title string) {
	g.Title = title
}

func (g *Goarg) SetExamples(examples []string) {
	g.Examples = examples
}

func (g *Goarg) AutomaticUsage() {
	var theUsage string
	theUsage += fmt.Sprintf("%v\n", g.Title)

	MaxSpace := 0
	for _, o := range g.Options {
		if len(o.Usage) > MaxSpace {
			MaxSpace = len(o.Usage)
		}
	}

	for _, o := range g.Options {
		theUsage += fmt.Sprintf("%-*s %v\n", MaxSpace, o.Usage, o.PlaceHolder)
	}

	// 0 Değil ise bir example vardır onu help'e ekleyelim.
	if len(g.Examples) != 0 {
		theUsage += fmt.Sprintf("\nExamples:\n")
		for i, v := range g.Examples {
			theUsage += fmt.Sprintf("%v. %v\n", i+1, v)
		}
	}

	g.Usage = theUsage
}

func (g *Goarg) AddOption(arg string, active bool, usage string, myError []string) {
	arg = strings.ReplaceAll(arg, " ", "")
	g.Options = append(g.Options, Option{strings.Split(arg, ","), active, usage, myError})
}

func (g *Goarg) AddMode(mode string) {

}

// Kullanıcıdan alıp parse ettiğin inputları string array olarak döndür
func (g *Goarg) Start() []Input {
	// Bütün boşluklar silinip geliyor buraya boşluklarla uğraşmana gerek yok.
	args := os.Args[1:] // All inputs

	Help(g, args)

	CheckValidOptions(g, args)

	return GetInputs(g, args)
}
