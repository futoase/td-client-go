package td

type DatabaseList struct {
	Databases []Database
}

type Database struct {
	Name         string `json:"name"`
	Count        int    `json:"count"`
	Created_At   string `json:"created_at"`
	Updated_At   string `json:"updated_at"`
	Organization string `json:"organization"`
}

type TableList struct {
	Database string  `json:"database"`
	Tables   []Table `json:"tables"`
}

type Table struct {
	Id                     int    `json:"id"`
	Name                   string `json:"name"`
	Estimated_Storage_Size int    `json:"estimated_storage_size"`
	Counter_Updated_At     string `json:"counter_updated_at"`
	Last_Log_Timestamp     string `json:"last_log_timestamp"`
	Expire_Days            string `json:"expire_days"`
	Type                   string `json:"type"`
	Count                  int    `json:"count"`
	Created_At             string `json:"created_at"`
	Updated_At             string `json:"updated_at"`
	Schema                 string `json:"schema"`
}

type JobsList struct {
	Count int   `json:"count"`
	From  int   `json:"from"`
	To    int   `json:"to"`
	Jobs  []Job `json:"jobs"`
}

type Job struct {
	Status             string `json:"status"`
	Job_Id             string `json:"job_id"`
	Created_At         string `json:"created_at"`
	Updated_At         string `json:"updated_at"`
	Start_At           string `json:"start_at"`
	End_At             string `json:"end_at"`
	Query              string `json:"query"`
	Type               string `json:"type"`
	Priority           int    `json:"priority"`
	Retry_Limit        int    `json:"retry_limit"`
	Hive_Result_Schema string `json:"hive_result_schema"`
	Result             string `json:"result"`
	Url                string `json:"url"`
	User_Name          string `json:"user_name"`
	Organization       string `json:"organization"`
	Database           string `json:"database"`
}

type JobResults struct {
	JobResult []JobResult `json:"job_result"`
}

type JobResult struct {
	Status string `json:"status"`
	Count  string `json:"count"`
}

type Schedules struct {
	Schedules []Schedule `json:"schedules"`
}

type Schedule struct {
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	Type        string `json:"type"`
	Query       string `json:"query"`
	TimeZone    string `json:"timezone"`
	Delay       int    `json:"delay"`
	Database    string `json:"database"`
	User_Name   string `json:"user_name"`
	Priority    int    `json:"prioriby"`
	Retry_Limit int    `json:"retry_limit"`
	Result      string `json:"result"`
	Next_Time   string `json:"next_time"`
}

type ServerStatus struct {
	Status string `json:"status"`
}

type Results struct {
	Result []Result `json:"result"`
}

type Result struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	Organization string `json:"organization"`
}

type CreateTableResult struct {
	Table    string `json:"table"`
	Type     string `json:"type"`
	Database string `json:"database"`
	Message  string `json:"message"`
}

type SwapTableResult struct {
	Table1   string `json:"table1"`
	Table2   string `json:"table2"`
	Database string `json:"database"`
	Message  string `json:"message"`
}

type CreateDatabaseResult struct {
	Database string `json:"database"`
	Message  string `json:"message"`
}

type DeleteDatabaseResult struct {
	Database string `json:"database"`
	Message  string `json:"message"`
}

type KillJobResult struct {
	Job_Id        string `json:"job_id"`
	Former_Status string `json:"former_status"`
}
