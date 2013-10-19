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
