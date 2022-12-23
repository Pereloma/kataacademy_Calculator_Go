package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if isArabic, _ := regexp.MatchString("^[1-90]+\\s*[-+/*]\\s*[1-90]+$", input); isArabic {
		arabicCalc(input)
	} else if isRoman, _ := regexp.MatchString("^([IVX])+\\s*[-+/*]\\s*([IVX])+$", input); isRoman {
		romanCalc(input)
	} else {
		panic("")
	}
}

func arabicCalc(input string) {
	if strings.Contains(input, "+") {
		sNum := strings.Split(input, "+")
		numOne, numTwo := arabicParse(sNum[0], sNum[1])
		fmt.Println(numOne + numTwo)
	} else if strings.Contains(input, "-") {
		sNum := strings.Split(input, "-")
		numOne, numTwo := arabicParse(sNum[0], sNum[1])
		fmt.Println(numOne - numTwo)
	} else if strings.Contains(input, "/") {
		sNum := strings.Split(input, "/")
		numOne, numTwo := arabicParse(sNum[0], sNum[1])
		fmt.Println(numOne / numTwo)
	} else if strings.Contains(input, "*") {
		sNum := strings.Split(input, "*")
		numOne, numTwo := arabicParse(sNum[0], sNum[1])
		fmt.Println(numOne * numTwo)
	}
}

func romanCalc(input string) {
	if strings.Contains(input, "+") {
		sNum := strings.Split(input, "+")
		numOne, numTwo := romanParse(sNum[0], sNum[1])
		fmt.Println(arabicToRoman(numOne + numTwo))
	} else if strings.Contains(input, "-") {
		sNum := strings.Split(input, "-")
		numOne, numTwo := romanParse(sNum[0], sNum[1])
		fmt.Println(arabicToRoman(numOne - numTwo))
	} else if strings.Contains(input, "/") {
		sNum := strings.Split(input, "/")
		numOne, numTwo := romanParse(sNum[0], sNum[1])
		fmt.Println(arabicToRoman(numOne / numTwo))
	} else if strings.Contains(input, "*") {
		sNum := strings.Split(input, "*")
		numOne, numTwo := romanParse(sNum[0], sNum[1])
		fmt.Println(arabicToRoman(numOne * numTwo))
	}
}

func arabicParse(sNumOne, sNumTwo string) (numOne, numTwo int) {
	var err error
	numOne, err = strconv.Atoi(strings.TrimSpace(sNumOne))
	if err != nil {
		panic(err)
	}
	numTwo, err = strconv.Atoi(strings.TrimSpace(sNumTwo))
	if err != nil {
		panic(err)
	}

	if numOne > 10 || numTwo > 10 || numOne < 1 || numTwo < 1 {
		panic("")
	}
	return
}

func romanParse(rNumOne, rNumTwo string) (numOne, numTwo int) {

	numOne = romanToArabic(strings.TrimSpace(rNumOne))
	numTwo = romanToArabic(strings.TrimSpace(rNumTwo))

	if numOne > 10 || numTwo > 10 || numOne < 1 || numTwo < 1 {
		panic("")
	}
	return
}

func romanToArabic(roman string) (num int) {
	var romanNum map[rune]int = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sNum := []rune(roman)
	if len(sNum) <= 0 {
		return 0
	}
	integerValue := 0
	prevNumber := romanNum[sNum[0]]
	for i := 1; i < len(sNum); i++ {
		ch := sNum[i]
		number := romanNum[ch]
		if number <= prevNumber {
			integerValue += prevNumber
		} else {
			integerValue -= prevNumber
		}
		prevNumber = number
	}
	integerValue += prevNumber

	if integerValue < 1 {
		panic("")
	}
	return integerValue
}

func arabicToRoman(arabic int) (numR string) {
	var romanNumKey = []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	var romanNumValue = []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}
	/*	var romanNum map[int]string = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}*/

	if arabic < 1 {
		panic("")
	}

	numR = ""

	for i := 0; i < len(romanNumKey); i++ {
		for arabic >= romanNumKey[i] {
			max := arabic / romanNumKey[i]
			arabic = arabic % romanNumKey[i]
			for j := 0; j < max; j++ {
				numR = numR + romanNumValue[i]
			}
		}
	}

	return numR
}
