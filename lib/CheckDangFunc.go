package lib

import "strings"

// Функция для проверки содержания запрещенных функций в программе
func ContainsBannedFunctions(programCode string, bannedFunctions map[string]bool) bool {
	// Разбиваем программный код на отдельные слова
	words := strings.Fields(programCode)
	// Проверяем каждое слово на содержание запрещенных функций

	for _, word := range words {
		if strings.Contains(programCode, word) {
			return true
		}
	}
	return false
}
