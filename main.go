package main

import (
	"fmt"
)

/*
	Функцмя определяет, написано римскими числами или арабскими и выдает значеения в формате int

и выдает число, в котором зашифровано, арабские числа или римскик 0-римские, 2- арабские, 1 - одно арабское и второе римское,
5 - одно или оба участника операции не соответствуют требованиям задачи
*/
func rimiliarab(chislo1 string, chislo2 string) (int, int, int) {
	var a, b, z int
	if chislo1 == "1" {
		a = 1
		z = 1
	} else if chislo1 == "2" {
		a = 2
		z = 1
	} else if chislo1 == "3" {
		a = 3
		z = 1
	} else if chislo1 == "4" {
		a = 4
		z = 1
	} else if chislo1 == "5" {
		a = 5
		z = 1
	} else if chislo1 == "6" {
		a = 6
		z = 1
	} else if chislo1 == "7" {
		a = 7
		z = 1
	} else if chislo1 == "8" {
		a = 8
		z = 1
	} else if chislo1 == "9" {
		a = 9
		z = 1
	} else if chislo1 == "10" {
		a = 10
		z = 1
	} else if chislo1 == "I" {
		a = 1
		z = 0
	} else if chislo1 == "II" {
		a = 2
		z = 0
	} else if chislo1 == "III" {
		a = 3
		z = 0
	} else if chislo1 == "IV" {
		a = 4
		z = 0
	} else if chislo1 == "V" {
		a = 5
		z = 0
	} else if chislo1 == "VI" {
		a = 6
		z = 0
	} else if chislo1 == "VII" {
		a = 7
		z = 0
	} else if chislo1 == "VIII" {
		a = 8
		z = 0
	} else if chislo1 == "IX" {
		a = 9
		z = 0
	} else if chislo1 == "X" {
		a = 10
		z = 0
	} else {
		z = 5
	}
	if chislo2 == "1" {
		b = 1
		z++
	} else if chislo2 == "2" {
		b = 2
		z++
	} else if chislo2 == "3" {
		b = 3
		z++
	} else if chislo2 == "4" {
		b = 4
		z++
	} else if chislo2 == "5" {
		b = 5
		z++
	} else if chislo2 == "6" {
		b = 6
		z++
	} else if chislo2 == "7" {
		b = 7
		z++
	} else if chislo2 == "8" {
		b = 8
		z++
	} else if chislo2 == "9" {
		b = 9
		z++
	} else if chislo2 == "10" {
		b = 10
		z++
	} else if chislo2 == "I" {
		b = 1
	} else if chislo2 == "II" {
		b = 2
	} else if chislo2 == "III" {
		b = 3
	} else if chislo2 == "IV" {
		b = 4
	} else if chislo2 == "V" {
		b = 5
	} else if chislo2 == "VI" {
		b = 6
	} else if chislo2 == "VII" {
		b = 7
	} else if chislo2 == "VIII" {
		b = 8
	} else if chislo2 == "IX" {
		b = 9
	} else if chislo2 == "X" {
		b = 10
	} else {
		z = 5
	}
	return a, b, z
}
func znaktoop(znak string) int {
	switch znak {
	case "+":
		return 1
	case "-":
		return 2
	case "*":
		return 3
	case "/":
		return 4
	default:
		return 5
	}
}
func finish(a int, b int, op int, x int) int {
	if op == 1 {
		return a + b
	} else if op == 2 {
		return a - b
	} else if op == 3 {
		return a * b
	} else { // если ор == 4
		return a / b
	}
}

func vivod(rez int, z int) {
	if z == 2 {
		fmt.Println(rez)
	} else if z == 0 {
		if rez < 0 {
			fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
		} else if rez == 0 {
			fmt.Println("Выдача паники, так как в римской системе нет нуля.")
		} else if rez == 1 {
			fmt.Println("I")
		} else if rez == 2 {
			fmt.Println("II")
		} else if rez == 3 {
			fmt.Println("III")
		} else if rez == 4 {
			fmt.Println("IV")
		} else if rez == 5 {
			fmt.Println("V")
		} else if rez == 6 {
			fmt.Println("VI")
		} else if rez == 7 {
			fmt.Println("VII")
		} else if rez == 8 {
			fmt.Println("VIII")
		} else if rez == 9 {
			fmt.Println("IX")
		} else if rez == 10 {
			fmt.Println("X")
		} else if rez == 11 {
			fmt.Println("XI")
		} else if rez == 12 {
			fmt.Println("XII")
		} else if rez == 13 {
			fmt.Println("XIII")
		} else if rez == 14 {
			fmt.Println("XIV")
		} else if rez == 15 {
			fmt.Println("XV")
		} else if rez == 16 {
			fmt.Println("XVI")
		} else if rez == 17 {
			fmt.Println("XVII")
		} else if rez == 18 {
			fmt.Println("XVIII")
		} else if rez == 19 {
			fmt.Println("XIX")
		} else if rez == 20 {
			fmt.Println("XX")
		} else if rez == 21 {
			fmt.Println("XXI")
		} else if rez == 24 {
			fmt.Println("XXIV")
		} else if rez == 25 {
			fmt.Println("XXV")
		} else if rez == 27 {
			fmt.Println("XXVII")
		} else if rez == 28 {
			fmt.Println("XXVIII")
		} else if rez == 30 {
			fmt.Println("XXX")
		} else if rez == 32 {
			fmt.Println("XXXII")
		} else if rez == 35 {
			fmt.Println("XXXV")
		} else if rez == 36 {
			fmt.Println("XXXVI")
		} else if rez == 40 {
			fmt.Println("XL")
		} else if rez == 42 {
			fmt.Println("XLII")
		} else if rez == 45 {
			fmt.Println("XLV")
		} else if rez == 48 {
			fmt.Println("XLVIII")
		} else if rez == 49 {
			fmt.Println("XLIX")
		} else if rez == 50 {
			fmt.Println("L")
		} else if rez == 54 {
			fmt.Println("LIV")
		} else if rez == 56 {
			fmt.Println("LVI")
		} else if rez == 60 {
			fmt.Println("LX")
		} else if rez == 63 {
			fmt.Println("LXIII")
		} else if rez == 64 {
			fmt.Println("LXIV")
		} else if rez == 72 {
			fmt.Println("LXXII")
		} else if rez == 70 {
			fmt.Println("LXX")
		} else if rez == 80 {
			fmt.Println("LXXX")
		} else if rez == 81 {
			fmt.Println("LXXXI")
		} else if rez == 90 {
			fmt.Println("XC")
		} else if rez == 100 {
			fmt.Println("C")
		}
	}
}

func main() {
	var chislo1, chislo2, znak, konets string
	// переменная konets существует для того, чтобы понять, есть ли ещё что-то после первых трёх переменных
	fmt.Scanln(&chislo1, &znak, &chislo2, &konets)
	//Проверяем на то, выглядит ли наша задача, как операции а + b, a - b, a * b, a / b. Если нет, то выдаем панику, если да, то переходим дальше
	if konets != "" || chislo2 == "" || !(znak == "+" || znak == "-" || znak == "*" || znak == "/") {
		fmt.Println("Паника, так как строка не является математической операцией, подходящей по формату")
	} else {
		var a, b, z, op int
		// z - переменная в котором зашифоровано римские числа или арабские. z=0 римские, z=1 одно число арабское и другое римское, z = 2 арабские, z= 5 говорит, что одно или оба члена операции не подходят под условия задачи
		a, b, z = rimiliarab(chislo1, chislo2)
		if z == 5 {
			fmt.Println("Паника, одно или оба члена операции не подходят под условия задачи")
		} else if z == 1 {
			fmt.Println("Паника, так как используются одновременно разные системы счисления.")
		} else {
			op = znaktoop(znak)
			var rez int
			rez = finish(a, b, op, z)
			vivod(rez, z)
		}
	}
}
