package cla

import (
	"os"
	"strings"
)

// Burada main true ise bu ana goarg'tır değil ise bir mode dur.
func Init() Goarg {
	// map kullanabilmek için initilaze etmek lazım make ile
	return Goarg{Title: "Example Title", Mode: make(map[string]*Goarg), Main: true}
}

func ModInit() Goarg {
	return Goarg{Title: "Mod Title", Mode: make(map[string]*Goarg), Main: false, ModeName: "ANAN"}
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
	if g.Main {
		g.Usage = CreateMainHelp(g)
	} else {
		g.Usage = CreateHelp(g)
	}
}

func (g *Goarg) AddOption(arg string, active bool, usage string, myError []string) {
	arg = strings.ReplaceAll(arg, " ", "")
	g.Options = append(g.Options, Option{strings.Split(arg, ","), active, usage, myError})
}

// Mode ekliyorum
// mesela go run main.go mode1 -h
// mesela go run main.go mode2 -t
// gibi birden fazla mod'un birden fazla arg'ı olucak.
func (g *Goarg) AddMode(mode string, m *Goarg) {
	m.ModeName = mode
	g.Mode[mode] = m
}

// git push origin v1.0.0 yaparak tagı paylaştım.
// Version tagı oluşturdum.

// Kullanıcıdan alıp parse ettiğin inputları string array olarak döndür
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

func startMode(args []string, m *Goarg) []Input {
	Help(m, args)

	CheckValidOptions(m, args)

	// Yukarıda yaptıım mode kısmını buraya eklersem belki kaliteli sonsuz bir mode olayı yapabilirim.

	return GetInputs(m, args)
}
