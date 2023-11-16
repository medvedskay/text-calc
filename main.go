package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func arabicToRoman(num int) (string, error) {
	// Функция для конвертации арабских цифр в римские
	if num <= 0 {
		return "", fmt.Errorf("Вывод ошибки, так как в римской системе нет отрицательных чисел и нуля.")
	}

	romanSymbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	romanValues := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	romanNum := ""

	for i := 12; i >= 0; i-- {
		for num >= romanValues[i] {
			romanNum += romanSymbols[i]
			num -= romanValues[i]
		}
	}

	return romanNum, nil
}

func romanToArabic(romanNum string) int {
	// Функция для конвертации римских цифр в арабские
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	arabicNum := 0
	prevValue := 0

	for i := len(romanNum) - 1; i >= 0; i-- {
		symbol := string(romanNum[i])
		value := romanMap[symbol]

		if value < prevValue {
			arabicNum -= value
		} else {
			arabicNum += value
		}

		prevValue = value
	}

	return arabicNum
}

func isValidIntExpression(s string) (string, error) {
	// Проверка на длину выражения
	lenExpRe := regexp.MustCompile(`^([1-9]|10)\s*([+\-*/])\s*([1-9]|10)\.`)
	// Проверка на Арабские совместно с Римскими цифрами
	arabRomanRe := regexp.MustCompile(`^([1-9]|10)\s*([+\-*/])\s*(I{1,3}|I?V|VI{1,3}|I?X)$`)
	// Проверка на Арабские совместно с Римскими цифрами
	romanArabRe := regexp.MustCompile(`^(I{1,3}|I?V|VI{1,3}|I?X)\s*([+\-*/])\s*([1-9]|10)$`)

	// Проверка на цифры арабские
	arabRe := regexp.MustCompile(`^([1-9]|10)\s*([+\-*/])\s*([1-9]|10)$`)
	// Проверка на цифры римские
	romRe := regexp.MustCompile(`^(I{1,3}|I?V|VI{1,3}|I?X)\s*([+\-*/])\s*(I{1,3}|I?V|VI{1,3}|I?X)$`)

	switch {
	case lenExpRe.MatchString(s):
		return "", fmt.Errorf("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	case romanArabRe.MatchString(s):
		return "", fmt.Errorf("Вывод ошибки, так как используются одновременно разные системы счисления.")
	case arabRomanRe.MatchString(s):
		return "", fmt.Errorf("Вывод ошибки, так как используются одновременно разные системы счисления.")
	case arabRe.MatchString(s):
		// Извлечение чисел и оператора из введенной строки
		matches := arabRe.FindStringSubmatch(s)

		a, _ := strconv.Atoi(matches[1])
		operator := matches[2]
		b, _ := strconv.Atoi(matches[3])

		// Выполнение операции
		result, err := calculateArab(operator, a, b)
		return strconv.Itoa(result), err
	case romRe.MatchString(s):
		// Извлечение чисел и оператора из введенной строки
		matches := romRe.FindStringSubmatch(s)

		a := matches[1]
		operator := matches[2]
		b := matches[3]

		// Выполнение операции
		result, err := calculateRoman(operator, a, b)
		return result, err
	default:
		return "", fmt.Errorf("Вывод ошибки, так как строка не является математической операцией.")
	}
}

func calculateRoman(operator string, a, b string) (string, error) {
	// Функция для выполнения операции калькулятора
	a1 := romanToArabic(a)
	b1 := romanToArabic(b)

	result, err := calculateArab(operator, a1, b1)

	if err != nil {
		return "", err
	}

	return arabicToRoman(result)
}

func calculateArab(operator string, a, b int) (int, error) {
	// Функция для выполнения операции калькулятора
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Неверная операция: %s", operator)
	}
}

func main() {
	for {
		fmt.Print("Введите значение: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		result, err := isValidIntExpression(text)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		fmt.Println(result)
	}

}
