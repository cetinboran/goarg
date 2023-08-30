package cla

import (
	"fmt"
	"os"
	"strings"

	errorHandler "github.com/cetinboran/goarg/errorHandler"
)

func InputInit(argumentName string, value string, modeName string, errors *OptionError) Input {
	return Input{Argument: argumentName, Value: value, ModeName: modeName, Errors: errors}
}

// Returns only option names from user input.
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

// Sets the input array
func GetInputs(g *Goarg, args []string) []Input {
	inputs := []Input{}

	for i, argValue := range args {
		for _, o := range g.Options {
			for _, v2 := range o.PlaceHolder {
				if v2 == argValue {
					// Eğer active ise input beklemeyecek.
					if o.Active == true {
						// Eğer bir i + 1 args boyutundan taşmıyorsa bir sağa bakıyoruz.
						if len(args) > i+1 {
							// Eğer bir sağdaki girdi - ile başlamıyorsa option değildir o zaman inputtur
							// Hata yolluyoruz active'i true olanların inputa ihtiyacı yok diye.
							if !strings.Contains(args[i+1], "-") {
								fmt.Println(errorHandler.GetErrors(argValue, 3))
								os.Exit(3)
							}
						}

						// Input'u initledik ve array'e attık.
						PlaceHolder := strings.ReplaceAll(argValue, "-", "")
						newInput := InputInit(PlaceHolder, "1", g.ModeName, g.Errors[v2])
						inputs = append(inputs, newInput)
					} else {
						PlaceHolder := strings.ReplaceAll(argValue, "-", "")
						newInput := InputInit(PlaceHolder, args[i+1], g.ModeName, g.Errors[v2])
						inputs = append(inputs, newInput)
					}
				}
			}
		}
	}

	return inputs
}

// Checks Valid Options.
func CheckValidOptions(g *Goarg, args []string) {
	// This check only for the options not for the actual inputs.
	onlyArgs := getOnlyOptionsFromArg(args)

	// Bu alttaki girilen inputtaki arg isimlerinin hepsini alıyor ve aynı olan var ise uyarı atıyor.
	// Burada önemli olan açık ise -u ve --user aynı yere input yolluyor ancak bunları aynı olarak görmüyor. onu düzeltmelyiiz ama şuanlk bu çalışıyor.
	var usedArgs []string
	for _, v := range onlyArgs {
		if len(usedArgs) == 0 {
			usedArgs = append(usedArgs, v)
		} else {
			for _, v2 := range usedArgs {
				if v2 == v {
					fmt.Println(errorHandler.GetErrors(v, 4))
					os.Exit(4)
				} else {
					usedArgs = append(usedArgs, v)
				}
			}
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

	// Eğer option hiç yok ise goarg içinde böyle bir input yok diyecek
	if len(g.Options) == 0 {
		fmt.Println(errorHandler.GetErrors(args[0], 1))
		os.Exit(1)
	}

	// Şimdiik bunu ekledim. EĞER ANA SETUPUN İÇİNDE OPTİON YOK İSE BÜTÜN OPTİONLARA VALİD DİYOR.
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
	var theUsage strings.Builder

	theUsage.WriteString(fmt.Sprintf("\n%s\n", g.Usage.Title))
	theUsage.WriteString("----------------------------\n\n")

	if g.Usage.Description != "" {
		theUsage.WriteString("Description\n")
		theUsage.WriteString("----------------------------\n")
		theUsage.WriteString(g.Usage.Description + "\n\n")
	}

	maxSpace := calculateMaxSpace(g.Options)

	writeOptionGroup(&theUsage, g.Options, false, maxSpace)
	writeOptionGroup(&theUsage, g.Options, true, maxSpace)

	writeExamples(&theUsage, g.Usage.Examples)

	// Write all modes under a single "Modes" header
	writeModeGroup(&theUsage, "Modes", g.Mods)

	return theUsage.String()
}

func writeModeGroup(builder *strings.Builder, modeGroupName string, mods map[string]*Goarg) {
	builder.WriteString(fmt.Sprintf("\n%s\n", modeGroupName))
	builder.WriteString("-----\n")

	for k, v := range mods {
		builder.WriteString(fmt.Sprintf("-> %s\n", k))
		writeModGroups(builder, v.Mods)
	}
}

func writeModGroups(builder *strings.Builder, mods map[string]*Goarg) {
	count := 1
	for k2 := range mods {
		builder.WriteString(fmt.Sprintf("%v. %s\n", count, k2))
		count++
	}
}

func calculateMaxSpace(options []Option) int {
	maxSpace := 0
	for _, o := range options {
		if len(o.Usage) > maxSpace {
			maxSpace = len(o.Usage)
		}
	}
	return maxSpace
}

func writeOptionGroup(builder *strings.Builder, options []Option, global bool, maxSpace int) {
	for _, o := range options {
		if o.Global == global {
			builder.WriteString(fmt.Sprintf("%-*s %v\n", maxSpace, o.Usage, o.PlaceHolder))
		}
	}
}

func writeExamples(builder *strings.Builder, examples []string) {
	if len(examples) != 0 {
		builder.WriteString("\nExamples\n")
		for i, v := range examples {
			builder.WriteString(fmt.Sprintf("%v. %v\n", i+1, v))
		}
	}
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
