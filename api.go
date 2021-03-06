package td

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
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

func (t *TreasureData) CreateDatabase(db_name string) CreateDatabaseResult {
	r := PostRequest(t.Options, "/v3/database/create/"+db_name, nil)
	result := CreateDatabaseResult{}
	err := json.Unmarshal([]byte(r), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *TreasureData) CreateTable(db_name string, table_name string, table_type string) CreateTableResult {
	r := PostRequest(t.Options, "/v3/table/create/"+db_name+"/"+table_name+"/"+table_type, nil)
	result := CreateTableResult{}
	err := json.Unmarshal([]byte(r), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *TreasureData) SwapTable(db_name string, table_1 string, table_2 string) SwapTableResult {
	r := PostRequest(t.Options, "/v3/table/swap/"+db_name+"/"+table_1+"/"+table_2, nil)
	result := SwapTableResult{}
	err := json.Unmarshal([]byte(r), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *TreasureData) DeleteDatabase(db_name string) DeleteDatabaseResult {
	r := PostRequest(t.Options, "/v3/database/delete/"+db_name, nil)
	result := DeleteDatabaseResult{}
	err := json.Unmarshal([]byte(r), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *TreasureData) ListJobs(params *url.Values) JobsList {
	if params == nil {
		params := &url.Values{}
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
	job := Job{}
	err := json.Unmarshal([]byte(r), &job)
	if err != nil {
		panic(err)
	}
	return job
}

func (t *TreasureData) KillJob(job_id int) KillJobResult {
	r := PostRequest(t.Options, "/v3/job/kill/"+strconv.Itoa(job_id), nil)
	killresult := KillJobResult{}
	err := json.Unmarshal([]byte(r), &killresult)
	if err != nil {
		panic(err)
	}
	return killresult
}

func (t *TreasureData) JobResult(job_id int) JobResults {
	params := &url.Values{}
	params.Add("format", "tsv")
	r := GetRequest(t.Options, "/v3/job/result/"+strconv.Itoa(job_id), params)
	b := bytes.NewBufferString(r)
	reader := csv.NewReader(b)
	reader.Comma = '\t'
	result, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	jobresult := JobResults{}
	for i := 0; i < len(result); i++ {
		jobresult.JobResult = append(jobresult.JobResult, JobResult{result[i][0], result[i][1]})
	}
	return jobresult
}

func (t *TreasureData) ListSchedules() Schedules {
	r := GetRequest(t.Options, "/v3/schedule/list", nil)
	schedules := Schedules{}
	err := json.Unmarshal([]byte(r), &schedules)
	if err != nil {
		panic(err)
	}
	return schedules
}

func (t *TreasureData) ListResult() Results {
	r := GetRequest(t.Options, "/v3/result/list", nil)
	results := Results{}
	err := json.Unmarshal([]byte(r), &results)
	if err != nil {
		panic(err)
	}
	return results
}

func (t *TreasureData) ServerStatus() ServerStatus {
	r := GetRequest(t.Options, "/v3/system/server_status", nil)
	status := ServerStatus{}
	err := json.Unmarshal([]byte(r), &status)
	if err != nil {
		panic(err)
	}
	return status
}
