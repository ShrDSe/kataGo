package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		curr := roman[s[i]]
		if curr < prev {
			result -= curr
		} else {
			result += curr
		}
		prev = curr
	}

	return result
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i, v := range values {
		for num >= v {
			result += numerals[i]
			num -= v
		}
	}

	return result
}
func someMath(a int, b int, op byte) string {
	switch op {
	case '+':
		return strconv.Itoa(a + b)
	case '-':
		return strconv.Itoa(a - b)
	case '*':
		return strconv.Itoa(a * b)
	case '/':
		return strconv.Itoa(a / b)
	default:
		return "Invalid operator"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the equation (ex: 2 + 2): ")
	eq, _ := reader.ReadString('\n')
	equation := strings.Split(strings.TrimSpace(eq), " ")

	if len(equation) == 3 {
		a, err := strconv.Atoi(equation[0])
		if err == nil {
			b, err := strconv.Atoi(equation[2])
			if err == nil {
				fmt.Println(someMath(a, b, equation[1][0]))
				return
			}
			fmt.Println("Invalid input: both numbers should be integers")
		}
		if romanToInt(equation[0]) != 0 {
			a := romanToInt(equation[0])
			if romanToInt(equation[2]) != 0 {
				b := romanToInt(equation[2])
				if equation[1][0] == '-' && b <= a {
					fmt.Println("Invalid input: A should be greater than B (no negative numbers in roman)")
					return
				}
				res, _ := strconv.Atoi(someMath(a, b, equation[1][0]))
				fmt.Println(intToRoman(res))
				return
			}
			fmt.Println("Invalid input: both numbers should be roman")
			return
		}
	}
	fmt.Println("Invalid input: equation should be in the format of 'number operator number'")
	return
}
