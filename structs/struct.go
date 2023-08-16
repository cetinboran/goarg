package structs

import (
	"fmt"
	"os"
	"strings"
)

type Goarg struct {
	Usage   string
	Options []Option
}

type Option struct {
	Option []string // "-s,--settings"
	Active bool
	Usage  string
	Error  []string // 1 , 2, 3
}

// Eklemeler diğer tarafa gitsin diye * kullanıyorum.
func (g *Goarg) SetUsage(usage string) {
	g.Usage = usage
}

func (g *Goarg) AddOption(arg string, active bool, usage string, myError []string) {
	arg = strings.ReplaceAll(arg, " ", "")

	g.Options = append(g.Options, Option{strings.Split(arg, ","), active, usage, myError})
}

// Kullanıcıdan alıp parse ettiğin inputları string array olarak döndür
func (g *Goarg) Start() {
	// Bütün boşluklar silinip geliyor buraya boşluklarla uğraşmana gerek yok.
	args := os.Args[1:]

	fmt.Println(args)

}

// func (x Goarg) Init() Goarg {
// 	return Goarg{"", nil}
// }
