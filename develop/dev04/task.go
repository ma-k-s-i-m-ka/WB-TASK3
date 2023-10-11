package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Пример использования функции
	words := []string{"корсет", "костер", "сектор", "стокер", "пятак", "пятка", "тяпка", "цунами",
		"умница ", "монета", "отмена", "немота", "листок", "слиток", "столик", "кот", "ток", "кто"}
	anagrams := find(words)

	for key, value := range anagrams {
		fmt.Printf("Анаграмма из букв [%s]: %v\n", key, value)
	}
}

func find(words []string) map[string][]string {
	// Создание мапы для хранения анаграмм
	res := make(map[string][]string)
	for _, word := range words {
		// Приведение слова к нижнему регистру и сортировка его букв
		charSlice := strings.Split(strings.ToLower(word), "")
		sort.Strings(charSlice)
		sortedWord := strings.Join(charSlice, "")
		// Добавление слова в соответствующее множество анаграмм
		res[sortedWord] = append(res[sortedWord], word)
	}
	// Удаление множеств из одного элемента
	for key, value := range res {
		if len(value) == 1 {
			delete(res, key)
		}
	}
	return res
}
