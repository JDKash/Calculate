package main

import (
	"errors"
	"strconv"
	"strings"
)

// Определение констант
const (
	MaxStringLength = 10
	MaxNumber       = 10
)

// Определение типа операции
type OperationType int

const (
	Add      OperationType = iota // Сложение
	Subtract                      // Вычитание
	Multiply                      // Умножение
	Divide                        // Деление
)

// Определение структуры операции
type Operation struct {
	LeftStr  string        // Левый операнд (строка)
	RightStr string        // Правый операнд (строка, может быть пустым)
	Number   int           // Операнд (число, может быть 0)
	OpType   OperationType // Тип операции
}

// Парсим ввод
func parseInput(input string) (Operation, error) {
	// Удаление пробелов и проверка пустой строки
	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		return Operation{}, errors.New("пустая строка ввода")
	}

	// Разделение строки на элементы
	parts := strings.Fields(trimmedInput)

	// Проверка количества элементов
	if len(parts) != 3 {
		return Operation{}, errors.New("ввод должен содержать ровно три элемента")
	}

	// Проверка и обработка первого элемента (должен быть строкой)
	leftStr, err := validateAndProcessString(parts[0])
	if err != nil {
		return Operation{}, err
	}

	// Определение типа операции
	opType, err := determineOperationType(parts[1])
	if err != nil {
		return Operation{}, err
	}

	// Обработка второго элемента в зависимости от типа операции
	rightStr, number, err := processSecondOperand(parts[2], opType)
	if err != nil {
		return Operation{}, err
	}

	// Создание и возврат Operation
	return Operation{
		LeftStr:  leftStr,
		RightStr: rightStr,
		Number:   number,
		OpType:   opType,
	}, nil
}

// Функция для валидации и обработки строки, заключенной в кавычки
func validateAndProcessString(str string) (string, error) {
	// Проверка, что строка заключена в двойные кавычки
	if !strings.HasPrefix(str, "\"") || !strings.HasSuffix(str, "\"") {
		return "", errors.New("строка должна быть заключена в двойные кавычки")
	}

	// Удаление кавычек
	processedStr := str[1 : len(str)-1]

	// Проверка длины строки
	if len(processedStr) > MaxStringLength {
		return "", errors.New("длина строки превышает максимально допустимую")
	}

	return processedStr, nil
}

// Функция для определения типа операции
func determineOperationType(opSymbol string) (OperationType, error) {
	switch opSymbol {
	case "+":
		return Add, nil
	case "-":
		return Subtract, nil
	case "*":
		return Multiply, nil
	case "/":
		return Divide, nil
	default:
		return 0, errors.New("неподдерживаемая операция")
	}
}

// Функция для обработки второго операнда в зависимости от типа операции
func processSecondOperand(operand string, opType OperationType) (string, int, error) {
	switch opType {
	case Add, Subtract:
		// Для сложения и вычитания второй операнд должен быть строкой
		processedStr, err := validateAndProcessString(operand)
		if err != nil {
			return "", 0, err
		}
		return processedStr, 0, nil
	case Multiply, Divide:
		// Для умножения и деления второй операнд должен быть числом
		number, err := strconv.Atoi(operand)
		if err != nil {
			return "", 0, errors.New("не удалось преобразовать второй операнд в число")
		}
		// Проверка диапазона числа
		if number < 1 || number > MaxNumber {
			return "", 0, errors.New("число вне допустимого диапазона")
		}
		return "", number, nil
	default:
		// Для полноты
		return "", 0, errors.New("неизвестный тип операции")
	}
}
