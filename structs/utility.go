package structs

import (
	"fmt"
	"strings"
)

func GetAllOptionsAsString(g *Goarg) string {
	var option string
	for _, o := range g.Options {
		optionP := strings.Join(o.PlaceHolder, " ")
		option += optionP + " "
	}

	return option
}

func CheckValidOptions(g *Goarg, args string) {
	AllOptions := g.Options[0]
	fmt.Println(AllOptions)
}

func arrContains(input string, arr []string) bool {
	for _, v := range arr {
		if v == input {
			return true
		}
	}

	return false
}
