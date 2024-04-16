package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nShell exited.")
				return
			}
			fmt.Println("Error reading input:", err)
			continue
		}

		// Удаление символа новой строки из ввода
		input = strings.TrimSuffix(input, "\n")

		// Проверка на команду выхода
		if input == "\\quit" {
			fmt.Println("Shell exited.")
			return
		}

		// Обработка ввода пользователя
		handleInput(input)
	}
}

func handleInput(input string) {
	// Разбиение введенной строки на аргументы
	args := strings.Fields(input)

	if len(args) == 0 {
		return
	}

	// Проверка наличия команды и ее обработка
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Println("Usage: cd <directory>")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println("Error changing directory:", err)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":

	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running ps command:", err)
		}
	default:
		runCommand(input)
	}
}

func runCommand(input string) {
	cmd := exec.Command("bash", "-c", input)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}
