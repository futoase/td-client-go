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

	t := client.ListTables(d.Databases[0].Name)
	for i := 0; i < len(t.Tables); i++ {
		fmt.Printf("id: %v, name: %v, estimated_storate_size: %v, counter_updated_at: %v, last_log_timestamp: %v, expire_days: %v, type: %v, count: %v, created_at: %v, updated_at: %v, schema: %v", t.Tables[i].Id, t.Tables[i].Name, t.Tables[i].Estimated_Storage_Size, t.Tables[i].Counter_Updated_At, t.Tables[i].Expire_Days, t.Tables[i].Type, t.Tables[i].Count, t.Tables[i].Created_At, t.Tables[i].Updated_At, t.Tables[i].Schema)
	}
}
