package td

import (
  "bytes"
  "log"
  "net/http"
)

type Options struct {
  Header string
  Host string
  ApiKey string
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

func NewTreasureData (o *Options) *TreasureData {
  t := new(TreasureData)
  t.Options = o
  return t
}

func (t *TreasureData) ListDatabases() string {
  return GetRequest(t.Options, "/v3/database/list")
}

func GetHostUrl (o *Options) string{
  return o.Protocol + "://" + o.Host
}

func GetRequest (o *Options, path string) string{
  client := &http.Client{}
  req, err := http.NewRequest("GET", GetHostUrl(o) + "/" + path, nil)
  req.Header.Add(o.Header, "TD1 " + o.ApiKey)
  resp, err := client.Do(req) 
  if err != nil {
    log.Fatal(err)
  }
  return ReadBody(resp)
}

func ReadBody(resp *http.Response) string {
  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  return buf.String()
}
