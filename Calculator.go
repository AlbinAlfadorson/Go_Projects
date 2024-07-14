package main

import (
    "fmt"
    "strconv"
    "strings"
)

// Словарь для преобразования римских чисел в арабские
var romanNumerals = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
    "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Срез для быстрого преобразования арабских чисел в римские
var arabicToRoman = []string{
    "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
    "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
    "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
    "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
    "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
    "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
    "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
    "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
    "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
    "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

func main() {
    var input string
    fmt.Print("Введите выражение: ")
    fmt.Scanln(&input) // Считываем ввод пользователя

    parts := strings.Fields(input) // Разбиваем ввод на части
    if len(parts) != 3 {
        panic("Неверный формат ввода") // Проверяем, что ввод состоит из трех частей
    }

    a, aIsRoman := parseNumber(parts[0]) // Парсим первое число
    b, bIsRoman := parseNumber(parts[2]) // Парсим второе число
    op := parts[1] // Получаем операцию

    if aIsRoman != bIsRoman {
        panic("Использование одновременно разных систем счисления не допускается")
    }

    var result int
    switch op { // Выполняем соответствующую операцию
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            panic("Деление на ноль")
        }
        result = a / b
    default:
        panic("Неподдерживаемая операция")
    }

    if aIsRoman {
        if result <= 0 {
            panic("Результат работы с римскими числами должен быть больше нуля")
        }
        fmt.Println(arabicToRoman[result]) // Выводим результат в римской системе
    } else {
        fmt.Println(result) // Выводим результат в арабских числах
    }
}

func parseNumber(s string) (int, bool) {
    if val, ok := romanNumerals[s]; ok {
        return val, true // Если это римское число, возвращаем его арабский эквивалент и true
    }
    
    num, err := strconv.Atoi(s)
    if err != nil || num < 1 || num > 10 {
        panic("Неверное число: " + s) // Если число не в диапазоне от 1 до 10, вызываем панику
    }
    return num, false // Возвращаем арабское число и false (не римское)
}
