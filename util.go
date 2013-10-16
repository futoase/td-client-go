package td

import (
  "bytes"
  "log"
  "net/http"
)

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
