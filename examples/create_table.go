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

	result := client.CreateTable("foo", "baz", "player")
	fmt.Printf("database: %v", result.Database)
	fmt.Printf("table: %v", result.Table)
	fmt.Printf("message: %v", result.Message)
}
