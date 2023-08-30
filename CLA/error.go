package cla

import (
	"log"

	"github.com/cetinboran/goarg/errorHandler"
)

func OptionErrorInit(errors []string) *OptionError {
	return &OptionError{Values: errors}
}

func (oe *OptionError) GetErrors(errorId int) {
	// Eğer olmayan bir key için bu fonksiyonu çağırmaya çalışırlarsa uyarı versin.
	if oe == nil {
		log.Fatal(errorHandler.GetErrors("", 10))

	}
	for i, v := range oe.Values {
		if i+1 == errorId {
			log.Fatal(v)
		}
	}

}
