package td

import (
	"encoding/json"
	"strconv"
)

func (t *TreasureData) ListDatabases() DatabaseList {
	r := GetRequest(t.Options, "/v3/database/list")
	dblist := DatabaseList{}
	err := json.Unmarshal([]byte(r), &dblist)
	if err != nil {
		panic(err)
	}
	return dblist
}

func (t *TreasureData) ListTables(db string) TableList {
	r := GetRequest(t.Options, "/v3/table/list/"+db)
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

func (t *TreasureData) ShowJobs(job_id int) string {
	return GetRequest(t.Options, "/v3/job/show/"+strconv.Itoa(job_id))
}

func (t *TreasureData) JobResult(job_id int) string {
	return GetRequest(t.Options, "/v3/job/result"+strconv.Itoa(job_id))
}

func (t *TreasureData) ListSchedules() string {
	return GetRequest(t.Options, "/v3/schedule/list")
}

func (t *TreasureData) ListResult() string {
	return GetRequest(t.Options, "/v3/result/list")
}

func (t *TreasureData) ServerStatus() string {
	return GetRequest(t.Options, "/v3/system/server_status")
}
