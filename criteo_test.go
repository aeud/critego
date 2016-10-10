package critego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

const (
	startDate string = "2016-07-01"
	endDate   string = "2016-07-01"
)

var now time.Time

func getParameters() map[string]string {
	bs, err := ioutil.ReadFile("parameters.json")
	if err != nil {
		fmt.Println("Error when reading: ", err)
	}
	params := make(map[string]string)
	err = json.Unmarshal(bs, &params)
	if err != nil {
		fmt.Println("Error when decoding: ", err)
	}
	return params
}

func getClient() *CriteoClient {
	params := getParameters()
	return NewCriteoClient(params["AppToken"], params["Username"], params["Password"], params["ClientVersion"])
}

func TestLogin(t *testing.T) {
	client := getClient()
	fmt.Println(client.AuthToken)
	client.Login()
	fmt.Println(client.AuthToken)
}

func TestGetAccount(t *testing.T) {
	client := getClient()
	client.Login()
	account := client.Account
	fmt.Println(string(Jsonify(account)))
}

func TestGetCampaigns(t *testing.T) {
	client := getClient()
	client.Login()
	campaigns := client.GetCampaigns()
	if l := len(campaigns); l > 0 {
		fmt.Printf("Found %v campaigns\n", l)
		campaign := campaigns[0]
		fmt.Println(string(Jsonify(campaign)))
		fmt.Printf("Campaign %v: %v categories\n", campaign.CampaignName, len(campaign.CategoryBids))
		categories := campaign.CategoryBids
		for i := 0; i < len(categories); i++ {
			category := categories[i]
			fmt.Println(category)
		}
	} else {
		fmt.Println("No campaign")
	}
}

func TestGetCategories(t *testing.T) {
	client := getClient()
	client.Login()
	categories := client.GetCategories()
	if l := len(categories); l > 0 {
		fmt.Printf("Found %v categories\n", l)
		for i := 0; i < len(categories); i++ {
			category := categories[i]
			fmt.Println(string(category.Jsonify()))
		}
	} else {
		fmt.Println("No category")
	}
}

func TestGetReport(t *testing.T) {
	client := getClient()
	client.Login()
	r := NewReportJob("Category", "Daily", startDate, endDate)
	j := client.ScheduleReportJob(r)
	for status := client.GetJobStatus(j); status != "Completed"; time.Sleep(5 * time.Second) {
		fmt.Println(status)
	}
	url := client.GetReportDownloadUrl(j)
	fmt.Println(url)
	rows := client.GetReport(url)
	fmt.Printf("Report has %v rows\n\n", len(rows))
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		fmt.Println(string(row.Jsonify()))
	}
}
