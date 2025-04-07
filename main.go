package main

import "fmt"

func main() {
	transactions := [3]int{1, 2, 3}
	banks := [2]string{"Tinkoff", "Alfa"}

	fmt.Println(transactions[1])
	banks[0] = "Sber"
	fmt.Println(banks)
}