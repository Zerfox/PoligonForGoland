package Errors

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Calculator() {

	var operandOne int
	var operandTwo int
	var operator string

	fmt.Println("Добро пожаловать в калькулятор")

	for {
		fmt.Println("\nДля ознакомления с правилами введите то что указанно в [] [0 help 0]")
		fmt.Println("Для выхода из приложения введите [] [0 exit 0]")

		_, err := fmt.Scan(&operandOne, &operator, &operandTwo)
		if err != nil {
			log.Fatal(err)
		}
		if operandOne < -1000 || operandTwo < -1000 {
			fmt.Println(errors.New("величина операнда не должна быть меньше -1000"))
			continue
		}
		if operandOne > 1000 || operandTwo > 1000 {
			fmt.Println(errors.New("величина операнда не должна быть больше 1000"))
			continue
		}

		switch operator {

		case "+":
			fmt.Println(operandOne + operandTwo)
			break
		case "-":
			fmt.Println(operandOne - operandTwo)
			break
		case "/":
			if operandTwo == 0 {
				fmt.Println(errors.New("на 0 делить нельзя"))
				continue
			}
			fmt.Println(operandOne / operandTwo)
			break
		case "*":
			fmt.Println(operandOne * operandTwo)
			break
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Запрещено  деление на 0")
			fmt.Println("Существует ограничение аргументов они не могут быть больше 1000 и меньше -1000 ")
			fmt.Println("Правила ввода:  вводить в таком формате 1 + 1 ")
			fmt.Println("То есть разделять операторы и операнд пробелами, Прим: [-1000 + 1000]")
			break

		default:
			fmt.Println("Не известная команда")
		}

	}

}
