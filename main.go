package main

import "fmt"

func main() {
	transactions := [5]int{1, 2, 3, 4, 5}
	banks := [2]string{"Tinkoff", "Alfa"}

	fmt.Println(transactions[1])
	banks[0] = "Sber"
	fmt.Println(banks)

	partial := transactions[1:4] // partial := transactions[1:] // partial := transactions[:4]
	fmt.Println(partial)
}