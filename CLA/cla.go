package cla

import (
	"os"
	"strings"
)

// Starts a main goarg.
func Init() Goarg {
	// map kullanabilmek için initilaze etmek lazım make ile
	return Goarg{Title: "Example Title", Mode: make(map[string]*Goarg), Main: true}
}

// Starts a mode goarg.
func ModInit() Goarg {
	return Goarg{Title: "Mod Title", Mode: make(map[string]*Goarg), Main: false, ModeName: "ANAN"}
}

// You can set your own usage.
func (g *Goarg) SetUsage(usage string) {
	g.Usage = usage
}

// Sets the title of usage
func (g *Goarg) SetTitle(title string) {
	g.Title = title
}

// Adds example for your usage.
func (g *Goarg) SetExamples(examples []string) {
	g.Examples = examples
}

// Adds automatic helper.
func (g *Goarg) AutomaticUsage() {
	if g.Main {
		g.Usage = CreateMainHelp(g)
	} else {
		g.Usage = CreateHelp(g)
	}
}

// Adds option for goarg.
func (g *Goarg) AddOption(arg string, active bool, usage string, myError []string) {
	arg = strings.ReplaceAll(arg, " ", "")
	g.Options = append(g.Options, Option{strings.Split(arg, ","), active, usage, myError})
}

// Adds option to the every mode.
func (g *Goarg) AddGlobalOption(arg string, active bool, usage string, myError []string) {
	arg = strings.ReplaceAll(arg, " ", "")
	g.Options = append(g.Options, Option{strings.Split(arg, ","), active, usage, myError})

	for _, g2 := range g.Mode {
		g2.Options = append(g2.Options, Option{strings.Split(arg, ","), active, usage, myError})
	}
}

// Adds mode to the main goarg.
func (g *Goarg) AddMode(mode string, m *Goarg) {
	m.ModeName = mode
	g.Mode[mode] = m
}

// git push origin v1.0.0 yaparak tagı paylaştım.
// Version tagı oluşturdum.

// Starts the code. Return []Input array.
func (g *Goarg) Start() []Input {
	// Bütün boşluklar silinip geliyor buraya boşluklarla uğraşmana gerek yok.
	args := os.Args[1:] // All inputs

	// Mode için burası biraz karışık oldu sonra bakarsın düzeltmek için.
	// Eğer ilk arg - içermiyorsa bu bir mod olmak zorundadır yoksa hatalı girilmiştir.
	// valid mod ise o modun option'ları çalışır değil ise hata vericektir.
	// args yok ise burayı atlasın yoksa args[0] a eriştiğimiz için hata alıyoruz.
	if len(args) > 0 {
		if !strings.Contains(args[0], "-") {
			mode := CheckValidMode(args[0], g, args[0])

			if mode != nil {
				return startMode(args[1:], mode)
			}
		}
	}

	Help(g, args)

	CheckValidOptions(g, args)

	return GetInputs(g, args)
}

// Same with the Start just use for mods.
func startMode(args []string, m *Goarg) []Input {
	Help(m, args)

	CheckValidOptions(m, args)

	// Yukarıda yaptıım mode kısmını buraya eklersem belki kaliteli sonsuz bir mode olayı yapabilirim.

	return GetInputs(m, args)
}
