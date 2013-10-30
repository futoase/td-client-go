package main

import (
	"fmt"
	"github.com/futoase/td-client-go"
	"io/ioutil"
	"net/url"
)

func main() {
	r, err := ioutil.ReadFile("apikey.txt")
	if err != nil {
		panic(err)
	}
	options := td.NewOptions(string(r))
	client := td.NewTreasureData(options)

	j := client.ListJobs(nil)

	for i := 0; i < len(j.Jobs); i++ {
		fmt.Printf("status: %v, job_id: %v, created_at: %v,  updated_at: %v, start_at: %v, end_at: %v, query: %v, type: %v, priority: %v, retry_limit: %v, hive_result_schema: %v, url: %v, user_name: %v, organization: %v, database: %v\n", j.Jobs[i].Status, j.Jobs[i].Job_Id, j.Jobs[i].Created_At, j.Jobs[i].Updated_At, j.Jobs[i].Start_At, j.Jobs[i].End_At, j.Jobs[i].Query, j.Jobs[i].Type, j.Jobs[i].Priority, j.Jobs[i].Retry_Limit, j.Jobs[i].Hive_Result_Schema, j.Jobs[i].Result, j.Jobs[i].Url, j.Jobs[i].User_Name, j.Jobs[i].Organization, string(j.Jobs[i].Database))
	}

	fmt.Printf("\n-----------------------------\n")
	params := url.Values{}
	params.Add("from", "3")
	params.Add("to", "4")
	j = client.ListJobs(&params)

	for i := 0; i < len(j.Jobs); i++ {
		fmt.Printf("status: %v, job_id: %v, created_at: %v,  updated_at: %v, start_at: %v, end_at: %v, query: %v, type: %v, priority: %v, retry_limit: %v, hive_result_schema: %v, url: %v, user_name: %v, organization: %v, database: %v\n", j.Jobs[i].Status, j.Jobs[i].Job_Id, j.Jobs[i].Created_At, j.Jobs[i].Updated_At, j.Jobs[i].Start_At, j.Jobs[i].End_At, j.Jobs[i].Query, j.Jobs[i].Type, j.Jobs[i].Priority, j.Jobs[i].Retry_Limit, j.Jobs[i].Hive_Result_Schema, j.Jobs[i].Result, j.Jobs[i].Url, j.Jobs[i].User_Name, j.Jobs[i].Organization, string(j.Jobs[i].Database))
	}

}
