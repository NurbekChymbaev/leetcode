package _058_length_of_last_word

import "strings"

func defangIPaddr(address string) string {

	return strings.Replace(address, ".", "[.]", -1)
}
