package scripts

import "strings"

func BuildArrayFromString(strlist string, seperator string) [][]string {
	result := strings.Split(strlist, seperator)
	var arr [][]string
	for _, r := range result {
		arr = append(arr, strings.Split(r, ";"))
	}
	return arr
}
