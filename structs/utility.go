package structs

import (
	"fmt"
	"os"
	"strings"

	"github.com/cetinboran/goarg/errorHandler"
)

func GetOnlyOptionsFromArg(args []string) []string {
	var onlyArgs string
	for _, v := range args {
		if strings.HasPrefix(v, "-") || strings.HasPrefix(v, "--") {
			onlyArgs += v + " "
		}
	}

	onlyArgs = strings.TrimSpace(onlyArgs)

	return strings.Split(onlyArgs, " ")
}

func CheckValidOptions(g *Goarg, args []string) {

	// This check only for the options not for the actual inputs.
	onlyArgs := GetOnlyOptionsFromArg(args)

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

func arrContains(input string, arr []string) bool {
	for _, v := range arr {
		if v == input {
			return true
		}
	}

	return false
}
