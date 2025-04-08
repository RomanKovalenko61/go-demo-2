package main

import (
	"fmt"
)

func main() {
	transactions := []int{1, 2, 3, 4, 5, 6}
	temp := transactions
	transactions = append(transactions, 100)
	fmt.Println(transactions[1:])
	fmt.Println(temp)
}