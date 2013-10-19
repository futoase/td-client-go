package td

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func (t *TreasureData) ListDatabases() DatabaseList {
	r := GetRequest(t.Options, "/v3/database/list", nil)
	dblist := DatabaseList{}
	err := json.Unmarshal([]byte(r), &dblist)
	if err != nil {
		panic(err)
	}
	return dblist
}

func (t *TreasureData) ListTables(db string) TableList {
	r := GetRequest(t.Options, "/v3/table/list/"+db, nil)
	tablelist := TableList{}
	err := json.Unmarshal([]byte(r), &tablelist)
	if err != nil {
		panic(err)
	}
	return tablelist
}

func (t *TreasureData) CreateDatabase(db_name string) bool {
	PostRequest(t.Options, "/v3/database/create/"+db_name, nil)
	return true
}

func (t *TreasureData) ListJobs(params *url.Values) JobsList {
	if params == nil {
		params := url.Values{}
		params.Add("from", "0")
	}
	r := GetRequest(t.Options, "/v3/job/list", params)
	jobslist := JobsList{}
	err := json.Unmarshal([]byte(r), &jobslist)
	if err != nil {
		panic(err)
	}
	return jobslist
}

func (t *TreasureData) ShowJobs(job_id int) Job {
	r := GetRequest(t.Options, "/v3/job/show/"+strconv.Itoa(job_id), nil)
	fmt.Printf(r)
	job := Job{}
	err := json.Unmarshal([]byte(r), &job)
	if err != nil {
		panic(err)
	}
	return job
}

func (t *TreasureData) JobResult(job_id int) string {
	return GetRequest(t.Options, "/v3/job/result"+strconv.Itoa(job_id), nil)
}

func (t *TreasureData) ListSchedules() string {
	return GetRequest(t.Options, "/v3/schedule/list", nil)
}

func (t *TreasureData) ListResult() string {
	return GetRequest(t.Options, "/v3/result/list", nil)
}

func (t *TreasureData) ServerStatus() string {
	return GetRequest(t.Options, "/v3/system/server_status", nil)
}
