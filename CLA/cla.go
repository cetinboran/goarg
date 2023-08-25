package cla

import (
	"os"
	"strings"
)

// Starts a main goarg.
func Init() Goarg {
	// map kullanabilmek için initilaze etmek lazım make ile
	Goarg := Goarg{Mods: make(map[string]*Goarg), Usage: &Usage{}}
	Goarg.Usage.Title = "Main Title"
	Goarg.ModeName = "Main"

	return Goarg
}

// Starts a mode goarg.
func ModInit() Goarg {
	Mod := Goarg{Mods: make(map[string]*Goarg), Usage: &Usage{}}
	Mod.Usage.Title = "Mod Title"
	Mod.ModeName = "Mode"
	return Mod
}

// Sets the All in the usage struct. You can use your own usage help.
func (g *Goarg) SetMessage(message string) {
	g.Usage.Message = message
}

// Sets the Title in the usage struct
func (g *Goarg) SetTitle(title string) {
	g.Usage.Title = title
}

// Sets the Examples in the usage struct
func (g *Goarg) SetExamples(examples []string) {
	g.Usage.Examples = examples
}

// Sets the description in the usage struct
func (g *Goarg) SetDescription(description string) {
	g.Usage.Description = description
}

// You can set all of the usage struct by once.
func (g *Goarg) SetUsage(title string, description string, examples []string) {
	g.Usage.Title = title
	g.Usage.Description = description
	g.Usage.Examples = examples
}

// Adds automatic helper.
func (g *Goarg) AutomaticUsage() {
	g.Usage.Message = CreateHelp(g)
}

// Adds option for goarg.
func (g *Goarg) AddOption(args string, active bool, usage string) {
	args = strings.ReplaceAll(args, " ", "")
	CheckOptionNames(args)
	CheckOptionNameIsBeingUsed(g, args)

	g.Options = append(g.Options, Option{strings.Split(args, ","), active, usage, false})
}

// Adds option to the every mode.
func (g *Goarg) AddGlobalOption(args string, active bool, usage string) {
	args = strings.ReplaceAll(args, " ", "")
	CheckOptionNames(args)

	// Burada hem modlarda hemde kendi içine böyle bir option name var mı diye bakmalı yoksa önceden olan bir option ile karışır
	// ve fazladan input döner.
	CheckOptionNameIsBeingUsed(g, args)
	CheckOptionNameIsBeingUsedInModes(g, args)

	/*
		Şimdilik burayı kapadım yani global option kendisi hariç bütün modlarına o optionu ekleyecek
		Bunu yapmamın nedeni eğer modun bir modu olursa ve aynı optionu kullanmak isterse hata veriyor karşm bu kullanılıyor diye
		çünkü en başka tanımlanmıştı o.
	*/
	// g.Options = append(g.Options, Option{strings.Split(args, ","), active, usage, true})

	for _, g2 := range g.Mods {
		// g2.AddOption(args, active, usage, myError)
		g2.Options = append(g2.Options, Option{strings.Split(args, ","), active, usage, true})
	}

}

// Adds mode to the main goarg.
func (g *Goarg) AddMode(mode string, m *Goarg) {
	CheckValidModeNames(g, mode)
	m.ModeName = mode
	g.Mods[mode] = m
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

	// Artık mode ların içinde de modlar olabiliyor.
	if len(args) > 0 {
		if !strings.Contains(args[0], "-") {
			mode := CheckValidMode(args[0], m, args[0])

			if mode != nil {
				return startMode(args[1:], mode)
			}
		}
	}

	Help(m, args)

	CheckValidOptions(m, args)

	// Yukarıda yaptıım mode kısmını buraya eklersem belki kaliteli sonsuz bir mode olayı yapabilirim.

	return GetInputs(m, args)
}
