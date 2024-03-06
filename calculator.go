package main

import (
	"errors"
	"strings"
)

// Функция для выполнения вычислений
func calculate(operation Operation) (string, error) {
	switch operation.OpType {
	case Add:
		// Сложение строк
		return handleStringOverflow(operation.LeftStr + operation.RightStr), nil
	case Subtract:
		// Вычитание строки из строки
		return handleStringOverflow(strings.Replace(operation.LeftStr, operation.RightStr, "", -1)), nil
	case Multiply:
		// Умножение строки на число
		return handleStringOverflow(strings.Repeat(operation.LeftStr, operation.Number)), nil
	case Divide:
		// Деление строки на число
		if operation.Number == 0 {
			return "", errors.New("деление на ноль")
		}
		partLength := len(operation.LeftStr) / operation.Number
		if partLength == 0 {
			return "", nil
		}
		return handleStringOverflow(operation.LeftStr[:partLength]), nil
	default:
		return "", errors.New("неизвестная операция")
	}
}

// Функция для обработки переполнения строки
func handleStringOverflow(result string) string {
	const maxOutputLength = 40
	if len(result) > maxOutputLength {
		return result[:maxOutputLength] + "..."
	}
	return result
}
