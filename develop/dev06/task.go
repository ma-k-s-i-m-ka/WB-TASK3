package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Options struct {
	Fields    []Field
	Delimiter string
	Separated bool
}

type Field struct {
	Start, End int
}

func main() {
	options := Options{}

	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	flagSet.Func("f", "select fields (columns)", options.ParseFields)
	flagSet.StringVar(&options.Delimiter, "d", "\t", "use a different separator")
	flagSet.BoolVar(&options.Separated, "s", false, "delimited lines only")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("failed to parse flags: %s", err)
		return
	}

	if len(options.Fields) == 0 {
		fmt.Println("Команда для запуска: develop/dev06/task.go -f ")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), options.Delimiter)

		if len(split) == 1 {
			if !options.Separated {
				fmt.Println(split[0])
			}

			continue
		}

		var items []string

		for _, field := range options.Fields {
			start := max(0, field.Start)
			end := min(len(split)-1, field.End)

			r := split[start : end+1]

			if len(r) == 0 {
				continue
			}

			for _, v := range r {
				items = append(items, v)
			}
		}

		fmt.Println(strings.Join(items, options.Delimiter))
	}
}
func MergeFields(fields []Field) []Field {
	if len(fields) <= 1 {
		return fields
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Start < fields[j].Start
	})

	merged := make([]Field, 0, len(fields))
	merged = append(merged, fields[0])

	for _, field := range fields[1:] {
		i := len(merged) - 1
		if field.Start-1 > merged[i].End {
			merged = append(merged, field)
		} else if field.End > merged[i].End {
			merged[i].End = field.End
		}
	}

	return merged
}

func (options *Options) ParseFields(s string) error {
	fields := strings.Split(s, ",")

	for _, field := range fields {
		var (
			before, after string
			found         bool
		)

		if before, found = strings.CutSuffix(field, "-"); found { // k-
			start, err := strconv.Atoi(before)
			if err != nil {
				return fmt.Errorf("failed to parse field: %s", err)
			}

			options.Fields = append(options.Fields, Field{Start: start - 1, End: math.MaxInt})
		} else if after, found = strings.CutPrefix(field, "-"); found { // -k
			end, err := strconv.Atoi(after)
			if err != nil {
				return fmt.Errorf("failed to parse field: %s", err)
			}

			options.Fields = append(options.Fields, Field{Start: math.MinInt, End: end - 1})
		} else if before, after, found = strings.Cut(field, "-"); found { // k-m
			start, err := strconv.Atoi(before)
			if err != nil {
				return fmt.Errorf("failed to parse field: %s", err)
			}

			end, err := strconv.Atoi(after)
			if err != nil {
				return fmt.Errorf("failed to parse field: %s", err)
			}

			options.Fields = append(options.Fields, Field{Start: start - 1, End: end - 1})
		} else { // k
			n, err := strconv.Atoi(field)
			if err != nil {
				return fmt.Errorf("failed to parse field: %s", err)
			}

			options.Fields = append(options.Fields, Field{Start: n - 1, End: n - 1})
		}
	}

	options.Fields = MergeFields(options.Fields)

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
func main() {
	// Определение флагов командной строки
	fieldsFlag := flag.String("f", "", "выбрать поля (колонки)")
	delimiterFlag := flag.String("d", "\t", "использовать другой разделитель")
	separatedFlag := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Разделитель по умолчанию - табуляция
	delimiter := *delimiterFlag

	// Считываем строки из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Если флаг -s включен и строка не содержит разделитель, пропускаем строку
		if *separatedFlag && !strings.Contains(line, delimiter) {
			continue
		}

		// Разбиваем строку на колонки с помощью разделителя
		columns := strings.Split(line, delimiter)

		// Если указан флаг -f, выбираем только указанные колонки
		if *fieldsFlag != "" {
			fieldIndexes := strings.Split(*fieldsFlag, ",")
			for _, indexStr := range fieldIndexes {
				index := parseIndex(indexStr)
				if index >= 0 && index < len(columns) {
					fmt.Print(columns[index])
				}
				if index < len(columns)-1 {
					fmt.Print(delimiter)
				}
			}
			fmt.Println()
		} else {
			// Если флаг -f не указан, выводим всю строку без изменений
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка при чтении STDIN:", err)
		os.Exit(1)
	}
}

// parseIndex преобразует строку в индекс колонки и возвращает его
func parseIndex(indexStr string) int {
	index, err := fmt.Sscanf(indexStr, "%d", &indexStr)
	if err != nil || index != 1 {
		// Если не удалось распарсить индекс или индекс не является целым числом,
		// возвращаем -1 для обозначения ошибки.
		return -1
	}
	// Индекс должен быть больше 0, но считаем с 0.
	return index - 1
}
*/
