package main

import (
	"fmt"
	"github.com/futoase/td-client-go"
	"io/ioutil"
)

func main() {
	r, err := ioutil.ReadFile("apikey.txt")
	if err != nil {
		panic(err)
	}
	options := td.NewOptions(string(r))
	client := td.NewTreasureData(options)

	result := client.SwapTable("foo", "cat", "dog")
	fmt.Printf("database: %v", result.Database)
	fmt.Printf("table1: %v", result.Table1)
	fmt.Printf("table2: %v", result.Table2)
	fmt.Printf("message: %v", result.Message)
}
