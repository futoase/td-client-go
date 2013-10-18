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
