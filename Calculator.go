package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var romanNumerals = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
    "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

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
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Введите выражение (или 'выход' для завершения): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if strings.ToLower(input) == "выход" {
            fmt.Println("Калькулятор завершает работу.")
            break
        }

        result, err := calculate(input)
        if err != nil {
            fmt.Println("Ошибка:", err)
            continue
        }

        fmt.Println("Результат:", result)
    }
}

func calculate(input string) (string, error) {
    parts := strings.Fields(input)
    if len(parts) != 3 {
        return "", fmt.Errorf("неверный формат ввода")
    }

    a, aIsRoman, err := parseNumber(parts[0])
    if err != nil {
        return "", err
    }

    b, bIsRoman, err := parseNumber(parts[2])
    if err != nil {
        return "", err
    }

    if aIsRoman != bIsRoman {
        return "", fmt.Errorf("использование одновременно разных систем счисления не допускается")
    }

    op := parts[1]
    var result int

    switch op {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            return "", fmt.Errorf("деление на ноль")
        }
        result = a / b
    default:
        return "", fmt.Errorf("неподдерживаемая операция")
    }

    if aIsRoman {
        if result <= 0 {
            return "", fmt.Errorf("результат работы с римскими числами должен быть больше нуля")
        }
        return arabicToRoman[result], nil
    }

    return strconv.Itoa(result), nil
}

func parseNumber(s string) (int, bool, error) {
    if val, ok := romanNumerals[s]; ok {
        return val, true, nil
    }

    num, err := strconv.Atoi(s)
    if err != nil || num < 1 || num > 10 {
        return 0, false, fmt.Errorf("неверное число: %s", s)
    }
    return num, false, nil
}
