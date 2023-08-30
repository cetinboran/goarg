package cla

import (
	"fmt"
	"log"
)

func OptionErrorInit(errors []string) *OptionError {
	return &OptionError{Values: errors}
}

func (oe *OptionError) GetErrors(errorId int, value string) {
	for i, v := range oe.Values {
		if i+1 == errorId {
			if v != "" {
				value := fmt.Sprintf(v+": %v\n", value)
				log.Fatal(value)
			} else {
				log.Fatal(v)
			}
		}
	}

}
