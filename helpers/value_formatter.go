package helpers

import "strconv"

func String(cmd, value string) string {
	return appendString(cmd) + appendString(value)
}

func StringArray(cmd string, array ...string) string {
	var result string
	if len(array) != 0 {
		for _, arr := range array {
			result += appendString(cmd) + appendString(arr)
		}
	}

	return result
}

func Int(cmd string, value int) string {
	return appendString(cmd) + appendString(strconv.Itoa(value))
}
