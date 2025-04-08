package main

import (
	"fmt"
)

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив
// Вывести сумму баланса в консоль

func main() {
	tr := make([]string, 0, 2)
	//tr[0] = "str0"
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "str1")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "str2")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "str3")
	fmt.Println(len(tr), cap(tr))
	fmt.Println(tr)


	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}

	fmt.Printf("Ваш баланс: %.2f", getSum(transactions))
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Println("Введите транзакцию или q для выхода")
	fmt.Scan(&transaction)
	return transaction
}

func getSum(arr []float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum
}