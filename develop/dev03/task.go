package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Объявляем флаги командной строки
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	k := flag.Int("k", 0, "Specify the column for sorting")
	n := flag.Bool("n", false, "Sort by numerical value")
	r := flag.Bool("r", false, "Sort in reverse order")
	u := flag.Bool("u", false, "Do not output repeated lines")
	flag.Parse()

	*inputFile = "example.txt"
	*outputFile = "res.txt"

	// Открываем входной файл для чтения
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Считываем строки из файла
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Функция для сравнения строк по числовому значению
	less := func(i, j int) bool {
		if *n {
			num1, err1 := strconv.Atoi(strings.Fields(lines[i])[*k-1])
			num2, err2 := strconv.Atoi(strings.Fields(lines[j])[*k-1])
			if err1 != nil || err2 != nil {
				return lines[i] < lines[j]
			}
			return num1 < num2
		}
		return lines[i] < lines[j]
	}

	// Сортируем строки
	if *r {
		sort.SliceStable(lines, func(i, j int) bool { return !less(i, j) })
	} else {
		sort.SliceStable(lines, less)
	}

	// Удаляем повторяющиеся строки, если необходимо
	if *u {
		var unique []string
		seen := make(map[string]bool)
		for _, line := range lines {
			if !seen[line] {
				unique = append(unique, line)
				seen[line] = true
			}
		}
		lines = unique
	}

	// Открываем выходной файл для записи
	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer output.Close()

	// Записываем отсортированные строки в выходной файл
	writer := bufio.NewWriter(output)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("File sorted successfully!")
}
