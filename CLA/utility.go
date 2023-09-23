package cla

import (
	"fmt"
	"os"
	"strings"

	errorHandler "github.com/cetinboran/goarg/errorHandler"
)

func InputInit(argumentName string, value string, modeName string) Input {
	return Input{Argument: argumentName, Value: value, ModeName: modeName}
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
func GetInputs(g *Goarg, args []string) ([]Input, map[string]*OptionError) {
	inputs := []Input{}

	for i, argValue := range args {
		for _, o := range g.options {
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
						newInput := InputInit(PlaceHolder, "1", g.modeName)
						inputs = append(inputs, newInput)
					} else {
						PlaceHolder := strings.ReplaceAll(argValue, "-", "")
						newInput := InputInit(PlaceHolder, args[i+1], g.modeName)
						inputs = append(inputs, newInput)
					}
				}
			}
		}
	}

	return inputs, g.errors
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

		for _, o := range g.options {

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
	if len(g.options) == 0 {
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

	theUsage.WriteString(fmt.Sprintf("\n=> %s\n\n", g.usage.Title))

	if g.usage.Description != "" {
		theUsage.WriteString("=> Description\n")
		theUsage.WriteString("  " + g.usage.Description + "\n")
	}

	maxSpace := 0
	for _, o := range g.options {
		if len(o.Usage) > maxSpace {
			maxSpace = len(o.Usage)
		}
	}

	writeOptions := func(options []*Option, global bool) {
		check := false
		var title string
		for _, o := range options {
			if o.Global == global {
				if check {
					theUsage.WriteString("\n" + title + ":\n")
					check = false
				}
				if o.title != "" {
					title = o.title
					check = true
				}
				theUsage.WriteString(fmt.Sprintf("  %-*s %v\n", maxSpace, o.Usage, o.PlaceHolder))
			}
		}
	}

	if g.usage.mainTitle != "" {
		theUsage.WriteString("\n" + g.usage.mainTitle + ":\n")
	}

	writeOptions(g.options, false)
	writeOptions(g.options, true)

	if len(g.usage.Examples) != 0 {
		theUsage.WriteString("\nExamples\n")
		for i, v := range g.usage.Examples {
			theUsage.WriteString(fmt.Sprintf("  %v. %v\n", i+1, v))
		}
	}

	if len(g.mods) != 0 {
		once := true

		for k, v := range g.mods {
			if once {
				theUsage.WriteString("\n=> Modes\n")
				once = false
			}

			modOptionsMaxSpace := 0
			for _, o := range v.options {
				if len(o.Usage) > modOptionsMaxSpace {
					modOptionsMaxSpace = len(o.Usage)
				}
			}

			theUsage.WriteString(fmt.Sprintf("\nMode: %s\n\n", k))
			theUsage.WriteString("Options:\n")
			writeOptions(v.options, false)

			if len(v.mods) != 0 {
				count := 1
				theUsage.WriteString(fmt.Sprintf("\n%s's Mods\n", v.modeName))
				for k2 := range v.mods {
					theUsage.WriteString(fmt.Sprintf("%v. %s\n", count, k2))
					count++
				}
			}
		}
	}

	return theUsage.String()
}

func Help(g *Goarg, args []string) {
	// Eğer arg boş ise usage ekrana yaz.
	if len(args) == 0 {
		fmt.Println(g.usage.Message)
		os.Exit(0)
	}

	// Eğer arg'ların içinde --help veya -h var ise usage ekrana yaz.
	if strings.Contains(strings.Join(args, " "), "--help") {
		fmt.Println(g.usage.Message)
		os.Exit(0)
	}

	// Usage yazınca programı bitiriyoruz ki devam edip invalid input demesin.
}

func CheckValidMode(firstArg string, g *Goarg, firstInput string) *Goarg {
	for k, m := range g.mods {
		if firstArg == k {
			return m
		}
	}

	fmt.Println(errorHandler.GetErrors(firstInput, 5))
	os.Exit(5)

	return nil
}

func CheckValidModeNames(g *Goarg, mode string) {
	for k := range g.mods {
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
		for _, o := range g.options {
			for _, v2 := range o.PlaceHolder {
				if v2 == v {
					fmt.Println(errorHandler.GetErrors(v+"\nthe name of the mod that uses this setting is: "+g.usage.Title, 8))
					os.Exit(8)
				}
			}
		}
	}

}

func CheckOptionNameIsBeingUsedInModes(g *Goarg, args string) {
	argsArr := strings.Split(args, ",")

	for _, v := range argsArr {
		for _, m := range g.mods {
			for _, o := range m.options {
				for _, v2 := range o.PlaceHolder {
					if v2 == v {
						fmt.Println(errorHandler.GetErrors(v+"\nthe name of the mod that uses this setting is: "+g.usage.Title, 8))
						os.Exit(8)
					}
				}
			}
		}
	}

}

func findGlobalOptionCount(g *Goarg) int {
	count := 0
	for _, o := range g.options {
		if o.Global {
			count++
		}
	}

	return count
}
