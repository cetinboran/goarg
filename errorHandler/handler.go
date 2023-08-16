package errorHandler

import "fmt"

func GetErrors(value string, errorId int) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("not an option: %v", value)
	case 2:
		return fmt.Sprintf("missing input for: %v", value)
	}

	return ""
}
