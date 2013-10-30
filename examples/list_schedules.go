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

	schedule := client.ListSchedules()

	for i := 0; i < len(schedule.Schedules); i++ {
		sc := schedule.Schedules[i]
		fmt.Printf("name: %v, cron: %v, type: %v, query: %v, timezone: %v, delay: %v, database: %v, user_name: %v, priority: %v, retry_limit: %v, result: %v, next_time: %v, ", sc.Name, sc.Cron, sc.Type, sc.Query, sc.TimeZone, sc.Delay, sc.Database, sc.User_Name, sc.Priority, sc.Retry_Limit, sc.Result, sc.Next_Time)
	}
}
