package td

type Options struct {
	Header   string
	Host     string
	ApiKey   string
	Protocol string
}

type TreasureData struct {
	Options *Options
}

func NewOptions(apikey string) *Options {
	if len(apikey) == 0 {
		return nil
	}

	o := new(Options)
	o.ApiKey = apikey
	o.Header = "AUTHORIZATION"
	o.Host = "api.treasure-data.com"
	o.Protocol = "http"
	return o
}

func NewTreasureData(o *Options) *TreasureData {
	t := new(TreasureData)
	t.Options = o
	return t
}
