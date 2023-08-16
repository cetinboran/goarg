package structs

import "strings"

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
	g.Options = append(g.Options, Option{strings.Split(arg, ","), active, usage, myError})
}

// func (x Goarg) Init() Goarg {
// 	return Goarg{"", nil}
// }
