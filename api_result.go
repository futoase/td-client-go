package td

type DatabaseList struct {
	Databases []Database
}

type Database struct {
	Name         string
	Count        int
	Created_At   string
	Updated_At   string
	Organization string
}

type TableList struct {
	Database string
	Tables   []Table
}

type Table struct {
	Id                     int
	Name                   string
	Estimated_Storage_Size int
	Counter_Updated_At     string
	Last_Log_Timestamp     string
	Expire_Days            string
	Type                   string
	Count                  int
	Created_At             string
	Updated_At             string
	Schema                 string
}

type JobsList struct {
	Count int
	From  int
	To    int
	Jobs  []Job
}

type Job struct {
	Status             string
	Job_Id             string
	Created_At         string
	Updated_At         string
	Start_At           string
	End_At             string
	Query              string
	Type               string
	Priority           int
	Retry_Limit        int
	Hive_Result_Schema string
	Result             string
	Url                string
	User_Name          string
	Organization       string
	Database           string
}
