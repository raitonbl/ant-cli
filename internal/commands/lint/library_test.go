package lint

import (
	"fmt"
)

func toString(each *Violation) string {

	if each == nil {
		return "nil"
	}

	return fmt.Sprintf("{\"path\":\"%s\" , \"message\":\"%s\"}", each.Path, each.Message)
}

func toText(array []Violation) string {

	if array == nil || len(array) == 0 {
		return "[]"
	}

	if len(array) == 1 {
		return fmt.Sprintf("[%s]", toString(&array[0]))
	}

	text := "["

	for index, value := range array {
		if index != 0 {
			text += ","
		}

		text += toString(&value)
	}

	text += "]"

	return text
}
