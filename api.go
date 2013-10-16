package td

import (
	"strconv"
)

func (t *TreasureData) ListDatabases() string {
	return GetRequest(t.Options, "/v3/database/list")
}

func (t *TreasureData) ListTables(db string) string {
	return GetRequest(t.Options, "/v3/table/list/"+db)
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
