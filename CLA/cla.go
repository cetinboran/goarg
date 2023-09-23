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
	Goarg := Goarg{mods: make(map[string]*Goarg), usage: &Usage{}, errors: make(map[string]*OptionError)}
	Goarg.usage.Title = "Main Title"
	Goarg.modeName = "Main"

	return Goarg
}

// Starts a mode goarg.
func ModInit() Goarg {
	Mod := Goarg{mods: make(map[string]*Goarg), usage: &Usage{}, errors: make(map[string]*OptionError)}
	Mod.usage.Title = "Mod Title"
	Mod.modeName = "Mode"
	return Mod
}

// Sets the All in the usage struct. You can use your own usage help.
func (g *Goarg) SetMessage(message string) {
	g.usage.Message = message
}

// Sets the Title in the usage struct
func (g *Goarg) SetTitle(title string) {
	g.usage.Title = title
}

// Sets the Examples in the usage struct
func (g *Goarg) SetExamples(examples []string) {
	g.usage.Examples = examples
}

// Sets the description in the usage struct
func (g *Goarg) SetDescription(description string) {
	g.usage.Description = description
}

func (g *Goarg) AddOptionTitle(title string) {
	length := len(g.options)

	if length != 0 {
		lastOption := g.options[length-1]
		lastOption.title = title
	} else {
		g.usage.mainTitle = title
	}
}

// You can set all of the usage struct by once.
func (g *Goarg) SetUsage(title string, description string, examples []string) {
	g.usage.Title = title
	g.usage.Description = description
	g.usage.Examples = examples
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
		g.errors[v] = OptionErrorInit(Error)
	}
}

// Adds automatic helper.
func (g *Goarg) AutomaticUsage() {
	g.usage.Message = CreateHelp(g)
}

// Adds option for goarg.
func (g *Goarg) AddOption(args string, active bool, usage string) {
	args = strings.ReplaceAll(args, " ", "")
	CheckOptionNames(args)
	CheckOptionNameIsBeingUsed(g, args)

	g.options = append(g.options, &Option{strings.Split(args, ","), active, usage, false, ""})
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

	for _, g2 := range g.mods {
		g2.options = append(g2.options, &Option{strings.Split(args, ","), active, usage, true, ""})
	}

}

// Adds mode to the main goarg.
func (g *Goarg) AddMode(mode string, m *Goarg) {
	CheckValidModeNames(g, mode)
	m.modeName = mode
	g.mods[mode] = m
}

// Starts the code. Return []cla.Input array.
func (g *Goarg) Start() ([]Input, map[string]*OptionError) {
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
func startMode(args []string, m *Goarg) ([]Input, map[string]*OptionError) {

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
