package cla

import (
	"fmt"
	"os"
	"strings"

	"github.com/cetinboran/goarg/errorHandler"
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

func GetInputs(args []string) map[string]string {

	inputs := make(map[string]string)
	for i, v := range args {
		if 0 > i-1 {
			fmt.Println("Please use this func after CheckValidOptions func.")
		}

		if !strings.HasPrefix(v, "-") {
			inputs[v] = args[i-1]
		}
	}

	return inputs
}

func CheckValidOptions(g *Goarg, args []string) {

	// This check only for the options not for the actual inputs.
	onlyArgs := getOnlyOptionsFromArg(args)

	// eğer false ise option aktifliği input gereklidir. O yüzden alt tarafta inputların doğru gelip gelmediğine bakmak için burada map oluşturdum.
	mapOfArgs := make(map[string]bool)

	check := true
	var option string
	for _, v := range onlyArgs {
		for _, o := range g.Options {
			// fmt.Println(strings.Join(o.PlaceHolder, " "), v)
			if !strings.Contains(strings.Join(o.PlaceHolder, " "), v) {
				check = false
				option = v
			} else {
				mapOfArgs[v] = o.Active
				check = true
				break
			}
		}
	}

	if !check {
		fmt.Println(errorHandler.GetErrors(option, 1))
		os.Exit(1)
	}

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

func Help(g *Goarg, args []string) {
	// Eğer arg boş ise usage ekrana yaz.
	if len(args) == 0 {
		fmt.Println(g.Usage)
		os.Exit(0)
	}

	// Eğer arg'ların içinde --help veya -h var ise usage ekrana yaz.
	if strings.Contains(strings.Join(args, " "), "--help") || strings.Contains(strings.Join(args, " "), "-h") {
		fmt.Println(g.Usage)
		os.Exit(0)
	}

	// Usage yazınca programı bitiriyoruz ki devam edip invalid input demesin.
}
