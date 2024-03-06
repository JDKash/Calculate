package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("String Calculator")
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при вводе:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		// Разбор ввода
		operation, err := parseInput(input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		// Вычисление результата
		result, err := calculate(operation)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		// Вывод результата
		fmt.Println("Результат:", result)
	}
}
