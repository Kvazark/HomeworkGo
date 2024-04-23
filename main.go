package main

import (
	"fmt"
	"unicode"
)

func main() {
	//reverceNum()
	//rightAngledTriangle()
	//getHoursAndMinutes()
	//formatedStr()
	squaring()
}

func reverceNum() {
	var num int
	fmt.Scanf("%d", &num)
	var reverceNum = 0
	for num > 0 {
		lastDigit := num % 10
		reverceNum = reverceNum*10 + lastDigit
		num = num / 10
	}
	fmt.Println(reverceNum)
}

func rightAngledTriangle() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	fmt.Println("Стороны треугольника: ", a, b, c)
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Неккоректный ввод. Введите 3 положительных числа")
	} else {
		if c*c == (a*a + b*b) {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Нерямоугольный")
		}
	}

}

func getHoursAndMinutes() {
	var k int
	fmt.Scanln(&k)

	if k > 0 && k < 86399 {
		h := k / 3600
		m := (k % 3600) / 60
		fmt.Printf("It is %d hours %d minutes.", h, m)
	} else {
		fmt.Println("Неккоректный ввод. Введите число, которое должно быть в диапозоне 0 < ваше_число < 86399")
	}
}

func containsOnlyLatin(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) || !unicode.Is(unicode.Latin, char) {
			return false
		}
	}
	return true
}

func formatedStr() {
	var str string
	fmt.Scanf("%s", &str)

	var new_str string
	const maxChars = 1000
	containsOnlyLatin := true

	if len(str) > maxChars {
		containsOnlyLatin = false
	}

	for _, char := range str {
		if !unicode.IsLetter(char) || !unicode.Is(unicode.Latin, char) {
			containsOnlyLatin = false
		}
	}

	if containsOnlyLatin {
		for i, ch := range str {
			if i != len(str)-1 {
				chr := string(ch) + "*"
				new_str = new_str + chr
				continue
			}
			new_str = new_str + string(ch)
		}
		fmt.Printf(string(new_str))

	} else {
		fmt.Println("Неккоректный ввод. Введите строку, содержащую только литинские символы, количество которых не превышает 1000.")
	}

}

func squaring() {
	var input float64

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Неккоректный ввод. Введите целое число.")
		return
	}

	if input != float64(int(input)) {
		fmt.Println("Введенное число не является целым.")
	} else {
		num := int(input)
		array := []int{}

		for num > 0 {
			lastDigit := num % 10
			array = append(array, lastDigit*lastDigit)
			num = num / 10
		}

		for _, value := range array {
			fmt.Printf("%d", value)
		}
	}

}
