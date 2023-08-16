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

func GetInputs(g *Goarg, args []string) []Input {
	inputs := []Input{}

	errorsMap := GetCustomErrors(g, args) // Custom yazılmış errorları çekiyorum.

	for i, v := range args {
		// -- Var ise onu input'a ata ve değerini 1 yani true yap
		if strings.HasPrefix(v, "--") {
			inputs = append(inputs, Input{Argument: v, Value: "1", Error: errorsMap[v]})
		} else if !strings.HasPrefix(v, "-") && !strings.HasPrefix(v, "--") { // - ve -- yok ise bu bir inputtur ona göre değerleri ata
			inputs = append(inputs, Input{Argument: args[i-1], Value: v, Error: errorsMap[args[i-1]]})
		}
	}

	return inputs
}

func GetCustomErrors(g *Goarg, args []string) map[string][]string {
	errorMap := make(map[string][]string)
	for _, v := range getOnlyOptionsFromArg(args) {
		for _, o := range g.Options {
			if strings.Contains(strings.Join(o.PlaceHolder, " "), v) {
				errorMap[v] = o.Error
			}
		}
	}

	return errorMap
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
			// Contains yerine yeni for açtın çünkü contains sıkıntı çıkarıyor.
			for _, v2 := range o.PlaceHolder {
				fmt.Println(v2)
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
