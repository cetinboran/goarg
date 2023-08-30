package errorHandler

import "fmt"

func GetErrors(value string, errorId int) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("not an option: %v", value)
	case 2:
		return fmt.Sprintf("missing input for: %v", value)
	case 3:
		return fmt.Sprintf("no need input for: %v", value)
	case 4:
		return fmt.Sprintf("you don't need more than one of this option: %v", value)
	case 5:
		return fmt.Sprintf("not a valid mod: %v", value)
	case 6:
		return fmt.Sprintf("this mod name is already in use: %v", value)
	case 7:
		return fmt.Sprintf("option names must start with - or --: %v", value)
	case 8:
		return fmt.Sprintf("this option name is being used please choose another: %v", value)
	case 9:
		return fmt.Sprintf("This is not a option name: %v", value)

	}

	return ""
}
