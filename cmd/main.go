package main

import (
	fisolver "as-capital-crawler-fi-ms/internal/fi_solver"
	"fmt"
)

func main() {
	data, err := fisolver.GetData("mana11")
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Quote.Current)
	fmt.Println(data.Quote.Min)
	fmt.Println(data.Quote.Max)
}
