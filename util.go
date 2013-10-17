package td

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func GetHostUrl(o *Options) string {
	return o.Protocol + "://" + o.Host
}

func CreateRequest(o *Options, method string, url string, body io.Reader) string {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	req.Header.Add(o.Header, "TD1 "+o.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return ReadBody(resp)
}

func GetRequest(o *Options, path string) string {
	return CreateRequest(o, "GET", GetHostUrl(o)+path, nil)
}

func PostRequest(o *Options, path string, body io.Reader) string {
	return CreateRequest(o, "POST", GetHostUrl(o)+path, body)
}

func ReadBody(resp *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}
