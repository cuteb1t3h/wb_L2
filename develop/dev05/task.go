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
	after := flag.Int("A", 0, "Print +N lines after the match")
	before := flag.Int("B", 0, "Print +N lines before the match")
	context := flag.Int("C", 0, "Print ±N lines around the match")
	count := flag.Bool("c", false, "Print count of matching lines")
	//ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNumber := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	// Получаем паттерн
	pattern := flag.Arg(0)
	if pattern == "" {
		fmt.Println("Usage: grep [options] pattern file")
		flag.PrintDefaults()
		return
	}

	// Открываем файл
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Считываем и обрабатываем строки файла
	var (
		lines       []string
		matchCount  int
		printBefore bool
		printAfter  int
	)
	for lineNumbers := 1; scanner.Scan(); lineNumbers++ {
		line := scanner.Text()

		matched := false
		if *fixed {
			matched = strings.Contains(line, pattern)
		} else {
			matched = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}

		if matched != *invert {
			if *count {
				matchCount++
			} else {
				if *lineNumber {
					fmt.Printf("%d:", lineNumber)
				}
				fmt.Println(line)
			}

			if *context > 0 {
				printBefore = true
				printAfter = *context
			}
		} else if printBefore || printAfter > 0 {
			if *lineNumber {
				fmt.Printf("%d:", lineNumber)
			}
			fmt.Println(line)

			if printBefore {
				printAfter--
			}
		}

		if printBefore {
			lines = append(lines, line)
			if len(lines) > *before {
				lines = lines[1:]
			}
		}

		if matched {
			printBefore = false
			printAfter = *after
		}
	}

	if *count {
		fmt.Println("Match count:", matchCount)
	}
}
