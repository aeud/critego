package critego

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ContentType string = "application/soap+xml"
	Endpoint    string = "https://advertising.criteo.com/API/v201305/AdvertiserService.asmx"
)

func HttpRequest(body []byte) []byte {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(Endpoint, ContentType, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error when posting: %v", err)
		return HttpRequest(body)
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error when converting: %v", err)
	}
	return bs
}

func HttpGetRequest(url string) []byte {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error when getting: %v", err)
	}
	defer resp.Body.Close()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatalf("Error when ungzipping: %v", err)
	}
	bs, err := ioutil.ReadAll(gz)
	if err != nil {
		log.Fatalf("Error when converting: %v", err)
	}
	return bs
}
