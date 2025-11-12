package main

import (
	"Poligon/Operator"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Ваш банковский счет")
	cashAccount := 1000
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Выберете команду" +
			"\n1 - Выдать наличные" +
			"\n2 - показать баланс" +
			"\n3 - провести онлайн оплату" +
			"\nexit - завершить работу")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			return
		}

		input := scanner.Text()

		err := Operator.ExecutionError()
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch input {

		case "1":
			fmt.Println("Введите сумму ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Ошибка ввода")
				return
			}
			userRequest, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("804 -ошибка ввода, некорректное число")
			}
			cashAccount, _ = Operator.GiveCash(cashAccount, userRequest)

			fmt.Println("Выдача наличных средств")
			break

		case "2":
			fmt.Println("Баланс составляет", cashAccount)
			break

		case "3":
			fmt.Println("Введите сумму оплаты")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Ошибка ввода")
				return
			}
			userRequest, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("804 -ошибка ввода, некорректное число")
			}
			cashAccount, _ = Operator.OnlinePayment(cashAccount, userRequest)
			fmt.Println("Оплата выполнена успешно")

			break

		case "exit":
			os.Exit(0)

		default:
			fmt.Println("ошибка 801 - не известная команда")
		}

	}

}

/*
Коды ошибок
 800 - ошибка проведения операции
 801 - ошибка, не известная команда
 802 - ошибка, недостаточно средств
 803 - ошибка ввода, сумма для выдачи не должна быть меньше 10
 804 - ошибка ввода, некорректное число
 805 -
 806 -
 807 -
 808 -
 809 -
 810 -
*/
