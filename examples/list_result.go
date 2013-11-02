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

	results := client.ListResult()
	for i := 0; i < len(results.Result); i++ {
		rs := results.Result[i]
		fmt.Printf("name: %v, url: %v, organization: %v", rs.Name, rs.Url, rs.Organization)
	}
}
