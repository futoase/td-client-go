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

	jobs := client.ListJobs(nil)

	job_id, err := strconv.Atoi(jobs.Jobs[0].Job_Id)
	if err != nil {
		panic(err)
	}

	result := client.KillJob(job_id)
	fmt.Printf("job_id: %v", result.Job_Id)
	fmt.Printf("former_status: %v", result.Former_Status)
}
