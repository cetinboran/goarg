package cla

import (
	"fmt"
	"os"
	"strings"

	errorHandler "github.com/cetinboran/goarg/errorHandler"
)

func getOnlyOptionsFromArg(args []string) []string {
	var onlyArgs string
	for _, v := range args {
		if strings.HasPrefix(v, "-") || strings.HasPrefix(v, "--") {
			onlyArgs += v + " "
		}
	}

	onlyArgs = strings.TrimSpace(onlyArgs)

	return strings.Split(onlyArgs, " ")
}

func GetInputs(g *Goarg, args []string) []Input {
	inputs := []Input{}

	for i, argValue := range args {
		for _, o := range g.Options {
			for _, v2 := range o.PlaceHolder {
				if v2 == argValue {
					if o.Active == true {
						if len(args) < i+1 {
							if !strings.Contains(args[i+1], "-") {
								fmt.Println(errorHandler.GetErrors(argValue, 3))
								os.Exit(3)
							}
						}
						newInput := InputInit()
						PlaceHolder := strings.ReplaceAll(argValue, "-", "")

						newInput.Argument = PlaceHolder
						newInput.Value = "1"
						newInput.ModeName = g.ModeName
						inputs = append(inputs, newInput)
						// inputs = append(inputs, Input{Argument: strings.ReplaceAll(argValue, "-", ""), Value: "1", ModeName: g.ModeName})
					} else {
						newInput := InputInit()
						PlaceHolder := strings.ReplaceAll(argValue, "-", "")

						newInput.Argument = PlaceHolder
						newInput.Value = args[i+1]
						newInput.ModeName = g.ModeName
						inputs = append(inputs, newInput)

						// inputs = append(inputs, Input{Argument: strings.ReplaceAll(argValue, "-", ""), Value: args[i+1], ModeName: g.ModeName})
					}
				}
			}
		}
	}

	return inputs
}

func InputInit() Input {
	return Input{}
}

func CheckValidOptions(g *Goarg, args []string) {
	// This check only for the options not for the actual inputs.
	onlyArgs := getOnlyOptionsFromArg(args)

	for _, v := range onlyArgs {
		if strings.Count(strings.Join(onlyArgs, " "), v) != 1 {
			fmt.Println(errorHandler.GetErrors(v, 4))
			os.Exit(4)
		}
	}

	// eğer false ise option aktifliği input gereklidir. O yüzden alt tarafta inputların doğru gelip gelmediğine bakmak için burada map oluşturdum.
	mapOfArgs := make(map[string]bool)

	check := true
	var option string
	for _, v := range onlyArgs {
		for _, o := range g.Options {
			// Contains yerine yeni for açtın çünkü contains sıkıntı çıkarıyor.
			for _, v2 := range o.PlaceHolder {
				if v != v2 {
					check = false
					option = v
				} else {
					mapOfArgs[v] = o.Active
					check = true
					break
				}
			}
			// Eğer check true ise doğru arg bulundu o zaman bu fordan da çık.
			if check {
				break
			}
		}
		// Bunu dışarda yaptın o yüzden ilk arg yanlış olsa bile sonraki arg doğru olunca invalid argları aldı
		if !check {
			fmt.Println(errorHandler.GetErrors(option, 1))
			os.Exit(1)
		}
	}

	// BURADA SORUN YOK GİBİ.
	// Eğer option active false ise ekstra bir input lazım yoksa hata versin.

	// Şimdiik bunu ekledim. EĞER ANA SETUPUN İÇİNDE OPTİON YOK İSE BÜTÜN OPTİONLARA VALİT DİYOR.
	if len(mapOfArgs) == 0 {
		fmt.Println(errorHandler.GetErrors(args[0], 2))
		os.Exit(2)
	}

	for i, v := range args {

		if !strings.HasPrefix(v, "-") {
			continue
		}

		needInput := !mapOfArgs[v]

		if !needInput {
			continue
		}

		// Eğer döngüde buraya kadar geldiyse girilen option'un input alması gerekiyordur.
		// Eğer args içinde bir sonraki eleman yok ise bir önceki option için missing input hatası atıyoruz.
		if len(args) <= i+1 {
			fmt.Println(errorHandler.GetErrors(args[i], 2))
			os.Exit(2)
		}

		if strings.HasPrefix(args[i+1], "-") || strings.HasPrefix(args[i+1], "--") {
			fmt.Println(errorHandler.GetErrors(args[i], 2))
			os.Exit(2)
		}
	}

}

func CreateHelp(g *Goarg) string {
	var theUsage string
	theUsage += fmt.Sprintf("%v\n", g.Usage.Title)
	theUsage += "----------------------------\n\n"

	if g.Usage.Description != "" {
		theUsage += "DESCRIPTION\n"
		theUsage += "----------------------------\n"

		theUsage += g.Usage.Description + "\n\n"
	}

	MaxSpace := 0
	for _, o := range g.Options {
		if len(o.Usage) > MaxSpace {
			MaxSpace = len(o.Usage)
		}
	}

	for _, o := range g.Options {
		if !o.Global {
			theUsage += fmt.Sprintf("%-*s %v\n", MaxSpace, o.Usage, o.PlaceHolder)
		}
	}

	once := true
	for _, o := range g.Options {
		if o.Global {
			if once {
				theUsage += "\nGlobal Option\n"
				theUsage += "----------------------------\n"
				once = false
			}
			theUsage += fmt.Sprintf("%-*s %v\n", MaxSpace, o.Usage, o.PlaceHolder)
		}
	}

	// 0 Değil ise bir example vardır onu help'e ekleyelim.
	if len(g.Usage.Examples) != 0 {
		theUsage += "\nExamples\n"
		// theUsage += "----------------------------\n"
		for i, v := range g.Usage.Examples {
			theUsage += fmt.Sprintf("%v. %v\n", i+1, v)

		}
	}

	if len(g.Mods) != 0 {

		once := true

		for k, v := range g.Mods {
			// Eğer bir option var ise bunu yapsın yoksa geçsin tabi global olanları saymamak lazım yoksa sıkıntı çıkıyor
			if once {
				// theUsage += "\nMods"
				theUsage += "\n-----"
				once = false
			}

			MaxSpace := 0
			for _, o := range v.Options {
				if len(o.Usage) > MaxSpace {
					MaxSpace = len(o.Usage)
				}
			}

			theUsage += "\nMode: " + k + "\n\n"
			theUsage += "Options:\n"
			for _, o := range v.Options {
				if !o.Global {
					theUsage += fmt.Sprintf("%-*s %v\n", MaxSpace, o.Usage, o.PlaceHolder)
				}
			}

			if len(v.Mods) != 0 {
				count := 1
				theUsage += "\n" + v.ModeName + "'s Mods\n"
				// theUsage += "\n-----\n"
				for k2 := range v.Mods {
					theUsage += fmt.Sprint(count) + ". " + k2 + "\n"
					count++
				}
			}
		}
	}

	return theUsage
}

func Help(g *Goarg, args []string) {
	// Eğer arg boş ise usage ekrana yaz.
	if len(args) == 0 {
		fmt.Println(g.Usage.Message)
		os.Exit(0)
	}

	// Eğer arg'ların içinde --help veya -h var ise usage ekrana yaz.
	if strings.Contains(strings.Join(args, " "), "--help") {
		fmt.Println(g.Usage.Message)
		os.Exit(0)
	}

	// Usage yazınca programı bitiriyoruz ki devam edip invalid input demesin.
}

func CheckValidMode(firstArg string, g *Goarg, firstInput string) *Goarg {
	for k, m := range g.Mods {
		if firstArg == k {
			return m
		}
	}

	fmt.Println(errorHandler.GetErrors(firstInput, 5))
	os.Exit(5)

	return nil
}

func CheckValidModeNames(g *Goarg, mode string) {
	for k := range g.Mods {
		if k == mode {
			fmt.Println(errorHandler.GetErrors(mode, 6))
			os.Exit(6)
		}
	}
}

func CheckOptionNames(args string) {
	argsArr := strings.Split(args, ",")

	for _, v := range argsArr {
		if !strings.HasPrefix(v, "-") {
			fmt.Println(errorHandler.GetErrors(v, 7))
			os.Exit(7)
		}
	}
}

func CheckOptionNameIsBeingUsed(g *Goarg, args string) {
	argsArr := strings.Split(args, ",")

	for _, v := range argsArr {
		for _, o := range g.Options {
			for _, v2 := range o.PlaceHolder {
				if v2 == v {
					fmt.Println(errorHandler.GetErrors(v+"\nthe name of the mod that uses this setting is: "+g.Usage.Title, 8))
					os.Exit(8)
				}
			}
		}
	}

}

func CheckOptionNameIsBeingUsedInModes(g *Goarg, args string) {
	argsArr := strings.Split(args, ",")

	for _, v := range argsArr {
		for _, m := range g.Mods {
			for _, o := range m.Options {
				for _, v2 := range o.PlaceHolder {
					if v2 == v {
						fmt.Println(errorHandler.GetErrors(v+"\nthe name of the mod that uses this setting is: "+g.Usage.Title, 8))
						os.Exit(8)
					}
				}
			}
		}
	}

}

func findGlobalOptionCount(g *Goarg) int {
	count := 0
	for _, o := range g.Options {
		if o.Global {
			count++
		}
	}

	return count
}
