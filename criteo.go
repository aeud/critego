package critego

import (
	"log"
	"time"
)

type Account struct {
	AdvertiserName string `xml:"advertiserName" json:"advertiserName"`
	Email          string `xml:"email" json:"email"`
	Currency       string `xml:"currency" json:"currency"`
	Timezone       string `xml:"timezone" json:"timezone"`
	Country        string `xml:"country" json:"country"`
}

func (a *Account) Jsonify() []byte {
	return Jsonify(a)
}

type BidInformation struct {
	BiddingStrategy string  `xml:"biddingStrategy" json:"biddingStrategy"`
	CpcBid          float32 `xml:"cpcBid>cpc" json:"cpcBid"`
}

type CategoryBid struct {
	CampaignCategoryUID int             `xml:"campaignCategoryUID" json:"campaignCategoryUID"`
	CampaignID          int             `xml:"campaignID" json:"campaignID"`
	CategoryID          int             `xml:"categoryID" json:"categoryID"`
	Selected            bool            `xml:"selected" json:"selected"`
	BidInformation      *BidInformation `xml:"bidInformation" json:"bidInformation"`
}

type Campaign struct {
	Account       *Account        `json:"account"`
	CampaignID    int             `xml:"campaignID" json:"campaignID"`
	CampaignName  string          `xml:"campaignName" json:"campaignName"`
	CampaignBid   *BidInformation `xml:"campaignBid" json:"campaignBid"`
	BudgetID      int             `xml:"budgetID" json:"budgetID"`
	RemainingDays int             `xml:"remainingDays" json:"remainingDays"`
	Status        string          `xml:"status" json:"status"`
	CategoryBids  []*CategoryBid  `xml:"categoryBids>categoryBid" json:"categoryBids"`
}

func (c *Campaign) Jsonify() []byte {
	return Jsonify(c)
}

type ReportJob struct {
	ReportSelector  string   `xml:"reportSelector"`
	ReportType      string   `xml:"reportType"`
	AggregationType string   `xml:"aggregationType"`
	StartDate       string   `xml:"startDate"`
	EndDate         string   `xml:"endDate"`
	SelectedColumns []string `xml:"selectedColumns"`
	IsResultGzipped bool     `xml:"isResultGzipped"`
}

type ReportJobResponse struct {
	JobID     int        `xml:"jobID"`
	JobStatus string     `xml:"jobStatus"`
	ReportJob *ReportJob `xml:"reportJob"`
}

func NewReportJob(reportType, aggregationType, startDate, endDate string) *ReportJob {
	r := new(ReportJob)
	r.ReportType = reportType
	r.AggregationType = aggregationType
	r.StartDate = startDate
	r.EndDate = endDate
	r.SelectedColumns = make([]string, 0)
	r.IsResultGzipped = true
	return r
}

type Category struct {
	Account          *Account `json:"account"`
	CategoryID       int      `xml:"categoryID" json:"categoryID"`
	CategoryName     string   `xml:"categoryName" json:"categoryName"`
	AvgPrice         float32  `xml:"avgPrice" json:"avgPrice"`
	NumberOfProducts int      `xml:"numberOfProducts" json:"numberOfProducts"`
	Selected         bool     `xml:"selected" json:"selected"`
}

func (c *Category) Jsonify() []byte {
	return Jsonify(c)
}

type CriteoClient struct {
	Username      string
	Password      string
	AuthToken     string
	AppToken      string
	ClientVersion string
	Account       *Account
}

func NewCriteoClient(appToken, username, password, clientVersion string) *CriteoClient {
	c := new(CriteoClient)
	c.AppToken = appToken
	c.Username = username
	c.Password = password
	c.ClientVersion = clientVersion
	return c
}

type ReportRow struct {
	Account               *Account `json:"account"`
	CampaignID            int      `xml:"campaignID,attr" json:"campaignID"`
	DateTimePosix         string   `xml:"dateTimePosix,attr" json:"dateTimePosix"`
	DateTime              string   `xml:"dateTime,attr" json:"dateTime"`
	CategoryID            int      `xml:"categoryID,attr" json:"categoryID"`
	CategoryName          string   `xml:"categoryName,attr" json:"categoryName"`
	Click                 int      `xml:"click,attr" json:"click"`
	Impressions           int      `xml:"impressions,attr" json:"impressions"`
	CTR                   float32  `xml:"CTR,attr" json:"CTR"`
	Revcpc                float32  `xml:"revcpc,attr" json:"revcpc"`
	Ecpm                  float32  `xml:"ecpm,attr" json:"ecpm"`
	Cost                  float32  `xml:"cost,attr" json:"cost"`
	Sales                 float32  `xml:"sales,attr" json:"sales"`
	ConvRate              float32  `xml:"convRate,attr" json:"convRate"`
	OrderValue            float32  `xml:"orderValue,attr" json:"orderValue"`
	SalesPostView         int      `xml:"salesPostView,attr" json:"salesPostView"`
	ConvRatePostView      float32  `xml:"convRatePostView,attr" json:"convRatePostView"`
	OrderValuePostView    float32  `xml:"orderValuePostView,attr" json:"orderValuePostView"`
	OverallCompetitionWin float32  `xml:"overallCompetitionWin,attr" json:"overallCompetitionWin"`
	CostPerOrder          float32  `xml:"costPerOrder,attr" json:"costPerOrder"`
}

func (r *ReportRow) Jsonify() []byte {
	return Jsonify(r)
}

func (c *CriteoClient) Login() {
	c.AuthToken = NewClientLoginCall(c).Do().ClientLoginResult
	c.Account = c.GetAccount()
}

func (c *CriteoClient) GetAccount() *Account {
	return NewGetAccountCall().Do(c).Account
}

func (c *CriteoClient) GetCampaigns() []*Campaign {
	campaigns := NewGetCampaignsCall().Do(c).Campaigns
	for i := 0; i < len(campaigns); i++ {
		campaigns[i].Account = c.Account
	}
	return campaigns
}

func (c *CriteoClient) GetCategories() []*Category {
	categories := NewGetCategoriesCall().Do(c).Categories
	for i := 0; i < len(categories); i++ {
		categories[i].Account = c.Account
	}
	return categories
}

func (c *CriteoClient) ScheduleReportJob(r *ReportJob) *ReportJobResponse {
	return NewScheduleReportJobCall(r).Do(c).JobResponse
}

func (c *CriteoClient) GetJobStatus(j *ReportJobResponse) string {
	return NewGetJobStatusCall(j).Do(c).JobStatus
}

func (c *CriteoClient) GetReportDownloadUrl(j *ReportJobResponse) string {
	return NewGetReportDownloadUrlCall(j).Do(c).JobURL
}

func (c *CriteoClient) GetReport(url string) []*ReportRow {
	rows := NewReportCollection(url).Rows
	for i := 0; i < len(rows); i++ {
		rows[i].Account = c.Account
	}
	return rows
}

func (c *CriteoClient) ScheduleExecuteDownloadJob(r *ReportJob, timeout time.Duration, retry int, printLog bool) []*ReportRow {
	j := c.ScheduleReportJob(r)
	i := 0
	for status := c.GetJobStatus(j); status != "Completed" && i < retry; time.Sleep(timeout * time.Second) {
		if printLog {
			log.Printf("Waiting for job %v (%v)", j.JobID, status)
		}
		i++
	}
	if i >= retry {
		if printLog {
			log.Println("Retry")
		}
		return c.ScheduleExecuteDownloadJob(r, timeout, retry, printLog)
	}
	url := c.GetReportDownloadUrl(j)
	if printLog {
		log.Println(url)
	}
	rows := c.GetReport(url)
	return rows
}
