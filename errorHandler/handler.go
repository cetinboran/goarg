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
	}

	return ""
}
