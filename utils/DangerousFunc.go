package utils

import (
	"fmt"
	"strings"
)

func BannedFunctions(programCode string, bannedFunctions string) bool {
	words := strings.Fields(programCode)
	banned := strings.Fields(bannedFunctions)

	for _, word := range words {
		for _, bannedWord := range banned {
			if strings.Contains(word, bannedWord) {
				fmt.Println(bannedWord)
				return true
			}
		}
	}
	return false
}
