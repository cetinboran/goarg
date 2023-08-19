package cla

import (
	"fmt"
	"os"
	"strings"

	errorHandler "github.com/cetinboran/goarg/errorHandler"
)

// Deneme
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
						inputs = append(inputs, Input{Argument: strings.ReplaceAll(argValue, "-", ""), Value: "1", Error: o.Error})
					} else {
						inputs = append(inputs, Input{Argument: strings.ReplaceAll(argValue, "-", ""), Value: args[i+1], Error: o.Error})
					}
				}
			}
		}
	}

	return inputs
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
	theUsage += fmt.Sprintf("%v\n", g.Title)
	theUsage += "****************************\n"

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

	return theUsage
}

func CreateMainHelp(g *Goarg) string {
	var help string
	count := 1

	help += CreateHelp(g)
	help += "\nModes\n****************************\n"

	for k := range g.Mode {
		help += fmt.Sprint(count) + ". " + k
		count++
	}

	return help
}

func Help(g *Goarg, args []string) {
	// Eğer arg boş ise usage ekrana yaz.
	if len(args) == 0 {
		fmt.Println(g.Usage)
		os.Exit(0)
	}

	// Eğer arg'ların içinde --help veya -h var ise usage ekrana yaz.
	if strings.Contains(strings.Join(args, " "), "--help") {
		fmt.Println(g.Usage)
		os.Exit(0)
	}

	// Usage yazınca programı bitiriyoruz ki devam edip invalid input demesin.
}

func CheckValidMode(firstArg string, g *Goarg, firstInput string) *Goarg {
	for k, m := range g.Mode {
		if firstArg == k {
			return m
		}
	}

	fmt.Println(errorHandler.GetErrors(firstInput, 5))
	os.Exit(5)

	return nil
}
