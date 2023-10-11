package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	// Определение флагов
	numColumK := flag.Int("k", 1, "указание колонки для сортировки")
	sortByNum := flag.Bool("n", false, "сортировать по числовому значению")
	sortByReverse := flag.Bool("r", false, "сортировать в обратном порядке")
	sortByUniq := flag.Bool("u", false, "не выводить повторяющиеся строки")
	sortByMonth := flag.Bool("M", false, "сортировать по названию месяца")
	sortCheck := flag.Bool("c", false, "проверять отсортированы ли данные")
	flag.Parse()

	// Открытие файла
	file, err := os.Open("develop/dev03/file.txt")
	if err != nil {
		err = fmt.Errorf("[Ошибка:%s] при открытии файла", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Чтение файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		err = fmt.Errorf("[Ошибка:%s] при чтении файла", scanner.Err())
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nСписок студентов:")
	fmt.Printf("%-7s | %-11s | %-5s |\n", "Имя", "Дата рождения", "Ср. балл")
	fmt.Println("------------------------------------")

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			name := fields[0]
			birthdate := fields[1]
			average := fields[2]
			fmt.Printf("%-7s | %-13s | %-8s |\n", name, birthdate, average)
		}
	}

	fmt.Println("------------------------------------")
	// Функция для сравнения строк при сортировке
	compare := func(i, j int) bool {
		a := strings.Fields(lines[i])
		b := strings.Fields(lines[j])

		// Проверка наличия достаточного количества колонок
		if *numColumK > len(a) {
			err = fmt.Errorf("[Введенное число K:%d] в то время как столбцов всего: %d", *numColumK, len(a))
			fmt.Println(err)
			os.Exit(1)
		}

		// Извлечение значений из указанной колонки
		valA := a[*numColumK-1]
		valB := b[*numColumK-1]

		// Преобразование в числа, если требуется
		if *sortByNum {
			numA := 0.0
			numB := 0.0
			_, errA := fmt.Sscanf(valA, "%f", &numA)
			_, errB := fmt.Sscanf(valB, "%f", &numB)

			if errA != nil || errB != nil {
				err = fmt.Errorf("[Столбец под номером K:%d] не является числовым", *numColumK)
				fmt.Println(err)
				os.Exit(1)
			}
			valA = fmt.Sprintf("%f", numA)
			valB = fmt.Sprintf("%f", numB)
		}

		// Сортировка по месяцу
		if *sortByMonth {
			dateA, _ := time.Parse("02.01.2006", valA)
			dateB, _ := time.Parse("02.01.2006", valB)
			if *sortByReverse {
				return dateA.Month() > dateB.Month()
			}
			return dateA.Month() < dateB.Month()
		}

		// Обратный порядок строк
		if *sortByReverse {
			return valA > valB
		}
		return valA < valB
	}

	// Сортировка среза строк
	sort.SliceStable(lines, func(i, j int) bool {
		return compare(i, j)
	})

	// Удаление повторяющихся строк
	if *sortByUniq {
		uniqueMap := make(map[string]bool)
		var uniqueLines []string
		for _, line := range lines {
			if !uniqueMap[line] {
				uniqueMap[line] = true
				uniqueLines = append(uniqueLines, line)
			}
		}
		lines = uniqueLines
	}

	fmt.Println("\nОтсортированный список студентов:")
	fmt.Printf("%-7s | %-11s | %-5s |\n", "Имя", "Дата рождения", "Ср. балл")
	fmt.Println("------------------------------------")

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			name := fields[0]
			birthdate := fields[1]
			average := fields[2]
			fmt.Printf("%-7s | %-13s | %-8s |\n", name, birthdate, average)
		}
	}

	if *sortCheck {
		if *sortByNum {
			fmt.Println("\nСортировка по числовому значению выполнена")
		}
		if *sortByReverse {
			fmt.Println("\nСортировка в обратном порядке выполнена")
		}
		if *sortByUniq {
			fmt.Println("\nУдаление повторяющихся строк выполнено")
		}
		if *sortByMonth {
			fmt.Println("\nСортировка по месяцу выполнена")
		}
	}
}
