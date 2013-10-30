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
	job := client.JobResult(job_id)
	fmt.Printf("status: %v, count: %v", job.JobResult[0].Status, job.JobResult[0].Count)

}
