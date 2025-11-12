package Operator

import (
	"errors"
	"math/rand"
)

func ExecutionError() error {

	probability := rand.Float64() * 100.0

	if probability < 30 {
		return errors.New("#####\n#802 - Ошибка проведения операции\n#####\n")
	}
	return nil
}

func GiveCash(wallet, giveOut int) (int, error) {

	if giveOut == 0 {

		return wallet, errors.New("803 - ошибка ввода, сумма для выдачи не должна быть меньше 10 ")
	}
	if giveOut > wallet {
		return wallet, errors.New("802 - ошибка, недостаточно средств")
	}
	wallet = wallet - giveOut
	return wallet, nil
}

func OnlinePayment(wallet, giveOut int) (int, error) {

	if giveOut == 0 {
		return wallet, errors.New("802 - ошибка, недостаточно средств")
	}

	if giveOut > wallet {
		return wallet, errors.New(`802 - ошибка, недостаточно средств`)
	}
	wallet -= giveOut
	return wallet, nil
}
