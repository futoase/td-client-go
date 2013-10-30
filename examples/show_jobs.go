package main

import (
	"fmt"
	"github.com/futoase/td-client-go"
	"io/ioutil"
	"strconv"
)

func main() {
	r, err := ioutil.ReadFile("apikey.txt")
	if err != nil {
		panic(err)
	}
	options := td.NewOptions(string(r))
	client := td.NewTreasureData(options)

	j := client.ListJobs(nil)
	job_id, err := strconv.Atoi(j.Jobs[0].Job_Id)
	if err != nil {
		panic(err)
	}
	job := client.ShowJobs(job_id)
	fmt.Printf("status: %v, job_id: %v, created_at: %v,  updated_at: %v, start_at: %v, end_at: %v, query: %v, type: %v, priority: %v, retry_limit: %v, hive_result_schema: %v, url: %v, user_name: %v, organization: %v, database: %v\n", job.Status, job.Job_Id, job.Created_At, job.Updated_At, job.Start_At, job.End_At, job.Query, job.Type, job.Priority, job.Retry_Limit, job.Hive_Result_Schema, job.Result, job.Url, job.User_Name, job.Organization, string(job.Database))
}
