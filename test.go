package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReturnError() {
	if x := recover(); x != nil {
		fmt.Printf("Вывод ошибки, %+v\n", x)
	}
}

func CheckDigit(operand string) string {
	reg := regexp.MustCompile("^(I|V|X)+$")
	check := reg.FindAllString(operand, -1)
	if len(check) > 0 {
		return "Romanian"
	}
	reg = regexp.MustCompile("^[0-9]+$")
	check = reg.FindAllString(operand, -1)
	if len(check) > 0 {
		return "Arabic"
	}
	err := errors.New("так как используются некорректные операнды.")
	panic(err)
}

func ToArabic(operand string) int {
	if CheckDigit(operand) == "Romanian" {
		romanianMap := map[string]int{
			"N": 0,
			"I": 1,
			"V": 5,
			"X": 10,
		}
		value := 0
		operand = operand + "N"
		i := 0
		for operand[i:i+1] != "N" {
			if romanianMap[operand[i:i+1]] < romanianMap[operand[i+1:i+2]] {
				value = value - romanianMap[operand[i:i+1]]
			} else {
				value = value + romanianMap[operand[i:i+1]]
			}
			i++
		}

		return value
	} else {
		tmpInt, _ := strconv.Atoi(operand)
		return tmpInt
	}
}

func CheckError(exampleArray []string) {
	textError := ""
	if len(exampleArray) < 3 {
		textError = "так как строка не является математической операцией."
	}
	if len(exampleArray) > 3 {
		textError = "так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	}
	if CheckDigit(exampleArray[0]) != CheckDigit(exampleArray[2]) {
		textError = "так как используются одновременно разные системы счисления."
	}
	reg := regexp.MustCompile("^(\\/|\\*|\\-|\\+)$")
	if len(reg.FindAllString(exampleArray[1], -1)) != 1 {
		textError = "так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	}

	operand1 := ToArabic(exampleArray[0])
	operand2 := ToArabic(exampleArray[2])
	if operand1 > 10 || operand1 < 1 || operand2 > 10 || operand2 < 1 {
		textError = "так как калькулятор должен принимать на вход числа от 1 до 10 включительно, не более."
	}
	if len(textError) > 0 {
		err := errors.New(textError)
		panic(err)
	}

}

func Calculation(exampleArray []string) int {
	operand1 := ToArabic(exampleArray[0])
	operand2 := ToArabic(exampleArray[2])
	var result int
	switch exampleArray[1] {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "/":
		result = operand1 / operand2
	case "*":
		result = operand1 * operand2
	}

	return result
}

func ToRomanian(arabic int) string {
	romanianResult := ""
	romanianDigits := [3][10]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	}
	bitDepth := len(strconv.Itoa(arabic))
	for i := bitDepth; i > 0; i-- {
		tmpDigit := romanianDigits[i-1][arabic/int(math.Pow10(i-1))]
		arabic = arabic % int(math.Pow10(i-1))
		romanianResult = romanianResult + tmpDigit
	}
	return romanianResult
}

func PrintResult(result int, exampleArray []string) {
	if CheckDigit(exampleArray[0]) == "Romanian" {
		if result < 1 {
			err := errors.New("так как в римской системе нет отрицательных и нулевых чисел.")
			panic(err)
		}
		fmt.Println(ToRomanian(result))
	} else {
		fmt.Println(result)
	}
}

func main() {

	defer ReturnError()

	//Запрашиваем ввод примера
	reader := bufio.NewReader(os.Stdin)

	//text := "VIII * ix"
	fmt.Println("Введите пример")
	text, _ := reader.ReadString('\n')

	//Очищаем, преобразуем строку и разбираем на массив
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)
	spaceRegexp := regexp.MustCompile(" +")
	exampleArray := spaceRegexp.Split(text, -1)

	//Проверяем корректность примера
	CheckError(exampleArray)

	//Посчитаем
	result := Calculation(exampleArray)

	//Сообщим ответ
	PrintResult(result, exampleArray)

}
