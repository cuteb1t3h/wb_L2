package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Объявляем флаги
	fields := flag.String("f", "", "Select fields")
	delimiter := flag.String("d", "\t", "Use a different delimiter")
	separated := flag.Bool("s", false, "Only output lines containing delimiter")
	flag.Parse()

	// Получаем запрошенные колонки
	selectedFields := strings.Split(*fields, ",")
	selectedFieldIndexes := make(map[int]bool)
	for _, field := range selectedFields {
		index := parseFieldIndex(field)
		selectedFieldIndexes[index] = true
	}

	// Создаем сканер для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разбиваем строку на колонки с использованием разделителя
		columns := strings.Split(line, *delimiter)

		// Выбираем только указанные колонки и выводим их
		var selectedColumns []string
		for index, column := range columns {
			if selectedFieldIndexes[index+1] {
				selectedColumns = append(selectedColumns, column)
			}
		}
		fmt.Println(strings.Join(selectedColumns, *delimiter))
	}
}

// Функция для разбора индекса поля
func parseFieldIndex(field string) int {
	index := 0
	fmt.Sscanf(field, "%d", &index)
	return index
}
