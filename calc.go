package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var romans = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var arabic = []string{"0","1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var operations = []string{"+", "-", "*", "/"}

var romans_big = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
var arabic_big = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}


func rta(r rune) int {
    switch r {
    case 'I':
        return 1
    case 'V':
        return 5
    case 'X':
        return 10
    case 'L':
        return 50
    case 'C':
        return 100 
    default:
        return 0
    }
}

func rom_to_ar(s string) int {
    var result int
    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && rta(rune(s[i])) < rta(rune(s[i+1])) {
            result -= rta(rune(s[i]))
        } else {
            result += rta(rune(s[i]))
        }
    }
    return result
}

func ar_to_rom(n int) string {
    rom_n := ""
    for i := 12; i >= 0; i-- {
        for n >= arabic_big[i] {
            rom_n += romans_big[i]
            n -= arabic_big[i]
        }
    }
    return rom_n
}

func ans(op, a, b string) {
    if contains(romans, a) || contains(romans, b) {
        if contains(arabic, a) || contains(arabic, b) {
            panic("Используются одновременно разные системы счисления")
        }
        numA := rom_to_ar(a)
        numB := rom_to_ar(b)
        switch op {
        case "+", "*":
            result := calculate(op, numA, numB)
            fmt.Println(ar_to_rom(result))
        case "-":
            if numA > numB {
                result := calculate(op, numA, numB)
                fmt.Println(ar_to_rom(result))
            } else {
                panic("В римской системе нет отрицательных чисел и нуля")
            }
        case "/":
            if numA >= numB {
                result := calculate(op, numA, numB)
                fmt.Println(ar_to_rom(result))
            } else {
                panic("В римской системе нет отрицательных чисел и нуля")
            }
        default:
            panic("Некорректная операция")
        }
    } else if contains(arabic, a) && contains(arabic, b) {
        numA, _ := strconv.Atoi(a)
        numB, _ := strconv.Atoi(b)
        result := calculate(op, numA, numB)
        fmt.Println(result)
    } else {
        panic("Недопустимый формат чисел")
    }
}


func calculate(op string, a, b int) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b != 0 {
            return a / b
        } else {
            panic("Деление на ноль")
        }
    }
    return 0
}

func contains(set []string, x string) bool {
    for _, s := range set {
        if s == x {
            return true
        }
    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        input := scanner.Text()
        primer := strings.Fields(input)
        if len(primer) == 3 {
            op := primer[1]
            a := primer[0]
            b := primer[2]
            ans(op, a, b)
        } else {
            panic("Неверный формат ввода")
        }
    }
}
