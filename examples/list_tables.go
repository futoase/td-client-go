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
	d := client.ListDatabases()

	for i := 0; i < len(d.Databases); i++ {
		row := d.Databases[i]
		fmt.Printf("table: %v\n count: %v\n created_at: %v\n updated_at: %v\n", row.Name, row.Count, row.Created_At, row.Updated_At)
	}
}
