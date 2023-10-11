package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Определение флагов
	//var str string
	sortAfter := flag.String("A", "", "печатать +N строк после совпадения")
	sortBefore := flag.String("B", "", "печатать +N строк до совпадения")
	sortContext := flag.String("C", "", "(A+B) печатать ±N строк вокруг совпадения")
	sortCount := flag.Bool("c", false, "количество строк")
	sortIgnoreReg := flag.Bool("i", false, "игнорировать регистр")
	sortIgnore := flag.String("v", "", "вместо совпадения, исключать")
	sortFixed := flag.String("F", "", "точное совпадение со строкой, не паттерн")
	sortLine := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()

	words := []string{"корсет", "КОСТЕР", "сектор", "стокер", "пятак", "телефон", "пятка", "ТЯПКА", "цунами",
		"умница ", "Монета", "ТелЕфон", "отмена", "Немота", "листок", "СЛИТОК", "столик", "кот", "ток", "кто",
		"Лестница", "Пол", "ТелЕфон", "стол", "купол", "парта", "тетрадь", "карандаш", "ноутбук", "мышь", "Карта", "телефон", "собака", "Кошка"}
	search := make([]string, 0)

	if *sortIgnoreReg {
		for i := range words {
			words[i] = strings.ToLower(words[i])
		}
	}

	if *sortFixed != "" {
		countNum := 0
		if *sortIgnoreReg {
			fmt.Println("Для поиска исключите флаг -i")
			return
		}
		for i := range words {
			if words[i] == *sortFixed {
				countNum = i
				search = append(search, words[i])
				if *sortLine {
					fmt.Printf("\nНомер найденного слова: %d", countNum)
				}
			}
		}
		if len(search) == 0 {
			fmt.Printf("\nСовпадения со словом %s не найдено", *sortFixed)
			return
		}
		fmt.Printf("\nСтроки имеющие точное совпадение со словом %s: %s", *sortFixed, search)
		if *sortCount {
			fmt.Printf("\nКоличество строк: %d", len(search))
		}
	}

	if *sortIgnore != "" {
		countNum := 0
		for i := range words {
			if words[i] != *sortIgnore {
				search = append(search, words[i])
			} else if *sortLine {
				countNum = i
				fmt.Printf("\nНомер слова исключеня: %d", countNum)
			}
		}
		fmt.Printf("\nСтроки после удаления слов %s: %s", *sortIgnore, search)
		if *sortCount {
			fmt.Printf("\nКоличество строк: %d", len(search))
		}
	}

	if *sortAfter != "" {
		countNum := -1
		for i := range words {
			if words[i] == *sortAfter {
				search = words[i+1:]
				countNum = i
				break
			}
		}
		if *sortLine {
			if countNum == -1 {
				fmt.Printf("\nСовпадения со словом %s не найдено", *sortAfter)
				return
			}
			fmt.Printf("\nНомер первого найденного слова: %d", countNum)
		}
		if len(search) == 0 {
			fmt.Printf("\nСовпадения со словом %s не найдено", *sortBefore)
			return
		}
		fmt.Printf("\nСтроки после совпадения со словом %s: %s", *sortAfter, search)
		if *sortCount {
			fmt.Printf("\nКоличество строк: %d", len(search))
		}
	}

	if *sortBefore != "" {
		countNum := -1
		for i := range words {
			if words[i] == *sortBefore {
				search = words[:i]
				countNum = i
				break
			}
		}
		if *sortLine {
			if countNum == -1 {
				fmt.Printf("\nСовпадения со словом %s не найдено", *sortAfter)
				return
			}
			fmt.Printf("\nНомер первого найденного слова: %d", countNum)
		}
		if len(search) == 0 {
			fmt.Printf("\nСовпадения со словом %s не найдено", *sortBefore)
			return
		}
		fmt.Printf("\nСтроки до совпадения со словом %s: %s", *sortBefore, search)
		if *sortCount {
			fmt.Printf("\nКоличество строк: %d", len(search))
		}
	}

	if *sortContext != "" {
		countNum := -1
		for i := range words {
			if words[i] == *sortContext {
				search = append(words[:i], words[i+1:]...)
				countNum = i
				break
			}
		}
		if *sortLine {
			if countNum == -1 {
				fmt.Printf("\nСовпадения со словом %s не найдено", *sortAfter)
				return
			}
			fmt.Printf("\nНомер первого найденного слова: %d", countNum)
		}
		if len(search) == 0 {
			fmt.Printf("\nСовпадения со словом %s не найдено", *sortContext)
			return
		}
		fmt.Printf("\nСтроки до и после совпадения со словом %s: %s", *sortContext, search)
		if *sortCount {
			fmt.Printf("\nКоличество строк: %d", len(search))
		}
	}
}
