package cla

import (
	"fmt"
	"os"
	"strings"

	"github.com/cetinboran/goarg/errorHandler"
)

// Starts a main goarg.
func Init() Goarg {
	// map kullanabilmek için initilaze etmek lazım make ile
	Goarg := Goarg{Mods: make(map[string]*Goarg), Usage: &Usage{}, Errors: make(map[string]*OptionError)}
	Goarg.Usage.Title = "Main Title"
	Goarg.ModeName = "Main"

	return Goarg
}

// Starts a mode goarg.
func ModInit() Goarg {
	Mod := Goarg{Mods: make(map[string]*Goarg), Usage: &Usage{}, Errors: make(map[string]*OptionError)}
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

// Add Errors Spesific for options.
func (g *Goarg) AddError(argument string, Error []string) {
	// Option eklerken ekledikleri argları buraya yazıyorlar onlara göre input structına erroları ekliyorum.
	argument = strings.TrimSpace(argument)
	argumentArr := strings.Split(argument, ",")

	for _, v := range argumentArr {
		if !strings.HasPrefix(v, "-") {
			fmt.Println(errorHandler.GetErrors(v, 9))
			os.Exit(9)
		}
		g.Errors[v] = OptionErrorInit(Error)
	}
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

	CheckOptionNameIsBeingUsed(g, args)
	CheckOptionNameIsBeingUsedInModes(g, args)

	/*
		Şimdilik burayı kapadım yani global option kendisi hariç bütün modlarına o optionu ekleyecek
		Bunu yapmamın nedeni eğer modun bir modu olursa ve aynı optionu kullanmak isterse hata veriyor karşm bu kullanılıyor diye
		çünkü en başka tanımlanmıştı o.
	*/
	// g.Options = append(g.Options, Option{strings.Split(args, ","), active, usage, true})

	for _, g2 := range g.Mods {
		g2.Options = append(g2.Options, Option{strings.Split(args, ","), active, usage, true})
	}

}

// Adds mode to the main goarg.
func (g *Goarg) AddMode(mode string, m *Goarg) {
	CheckValidModeNames(g, mode)
	m.ModeName = mode
	g.Mods[mode] = m
}

// Starts the code. Return []cla.Input array.
func (g *Goarg) Start() []Input {
	// All inputs
	args := os.Args[1:]

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

	return GetInputs(m, args)
}
