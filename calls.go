package critego

import (
	"encoding/xml"
	"log"
)

const (
	XMLNS = "https://advertising.criteo.com/API/v201305"
)

type CriteoRequestCallBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Content interface{}
}

func NewCriteoRequestCallBody(c interface{}) *CriteoRequestCallBody {
	b := new(CriteoRequestCallBody)
	b.Content = c
	return b
}

type CriteoRequestCallHeaderContent struct {
	XMLName       xml.Name `xml:"apiHeader"`
	AuthToken     string   `xml:"authToken"`
	AppToken      string   `xml:"appToken"`
	ClientVersion string   `xml:"clientVersion"`
}

func NewCriteoRequestCallHeaderContent(c *CriteoClient) *CriteoRequestCallHeaderContent {
	content := new(CriteoRequestCallHeaderContent)
	content.AuthToken = c.AuthToken
	content.AppToken = c.AppToken
	content.ClientVersion = c.ClientVersion
	return content
}

type CriteoRequestCallHeader struct {
	XMLName xml.Name `xml:"soap:Header"`
	XMLNS   string   `xml:"xmlns,attr"`
	Content *CriteoRequestCallHeaderContent
}

func NewCriteoRequestCallHeader(c *CriteoClient) *CriteoRequestCallHeader {
	h := new(CriteoRequestCallHeader)
	h.XMLNS = XMLNS
	if c != nil {
		h.Content = NewCriteoRequestCallHeaderContent(c)
	}
	return h
}

type CriteoRequestCall struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	RefI    string   `xml:"xmlns:xsi,attr"`
	RefD    string   `xml:"xmlns:xsd,attr"`
	RefSoap string   `xml:"xmlns:soap,attr"`
	Header  *CriteoRequestCallHeader
	Body    *CriteoRequestCallBody
}

func NewCriteoRequestCall(b *CriteoRequestCallBody, h *CriteoRequestCallHeader) *CriteoRequestCall {
	r := new(CriteoRequestCall)
	r.RefI = "http://www.w3.org/2001/XMLSchema-instance"
	r.RefD = "http://www.w3.org/2001/XMLSchema"
	r.RefSoap = "http://www.w3.org/2003/05/soap-envelope"
	r.Body = b
	r.Header = h
	return r
}

type CriteoResponseBody struct {
	Content []byte `xml:",innerxml"`
}

type CriteoResponse struct {
	XMLName xml.Name            `xml:"Envelope"`
	Body    *CriteoResponseBody `xml:"Body"`
}

func NewCriteoResponse() *CriteoResponse {
	r := new(CriteoResponse)
	return r
}

func (r *CriteoRequestCall) Bytes() []byte {
	bs, err := xml.Marshal(r)
	if err != nil {
		log.Fatalf("Error when XML Marshal: %v", err)
	}
	return bs
}

func (r *CriteoRequestCall) Do() []byte {
	bs := r.Bytes()
	data := HttpRequest(bs)
	v := NewCriteoResponse()
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v.Body.Content
}

// Login

type ClientLoginCall struct {
	XMLName  xml.Name `xml:"clientLogin"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	Source   string   `xml:"source"`
	Ref      string   `xml:"xmlns,attr"`
}

func NewClientLoginCall(c *CriteoClient) *ClientLoginCall {
	clc := new(ClientLoginCall)
	clc.Username = c.Username
	clc.Password = c.Password
	clc.Source = c.ClientVersion
	clc.Ref = "https://advertising.criteo.com/API/v201305"
	return clc
}

func (call *ClientLoginCall) Do() *ClientLoginResponse {
	data := NewCriteoRequestCall(NewCriteoRequestCallBody(call), NewCriteoRequestCallHeader(nil)).Do()
	v := new(ClientLoginResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type ClientLoginResponse struct {
	XMLName           xml.Name `xml:"clientLoginResponse"`
	ClientLoginResult string   `xml:"clientLoginResult"`
}

// Get account

type GetAccountCall struct {
	XMLName xml.Name `xml:"getAccount"`
	XMLNS   string   `xml:"xmlns,attr"`
}

func NewGetAccountCall() *GetAccountCall {
	call := new(GetAccountCall)
	call.XMLNS = XMLNS
	return call
}

func (gac *GetAccountCall) Do(c *CriteoClient) *GetAccountResponse {
	b := NewCriteoRequestCallBody(gac)
	h := NewCriteoRequestCallHeader(c)
	call := NewCriteoRequestCall(b, h)
	data := call.Do()
	v := new(GetAccountResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type GetAccountResponse struct {
	XMLName xml.Name `xml:"getAccountResponse"`
	Account *Account `xml:"getAccountResult"`
}

// Get campaigns

type GetCampaignsCall struct {
	XMLName xml.Name `xml:"getCampaigns"`
	XMLNS   string   `xml:"xmlns,attr"`
}

func NewGetCampaignsCall() *GetCampaignsCall {
	call := new(GetCampaignsCall)
	call.XMLNS = XMLNS
	return call
}

func (call *GetCampaignsCall) Do(c *CriteoClient) *GetCampaignsResponse {
	data := NewCriteoRequestCall(NewCriteoRequestCallBody(call), NewCriteoRequestCallHeader(c)).Do()
	v := new(GetCampaignsResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type GetCampaignsResponse struct {
	XMLName   xml.Name    `xml:"getCampaignsResponse"`
	Campaigns []*Campaign `xml:"getCampaignsResult>campaign"`
}

// Get categories

type GetCategoriesCall struct {
	XMLName xml.Name `xml:"getCategories"`
	XMLNS   string   `xml:"xmlns,attr"`
}

func NewGetCategoriesCall() *GetCategoriesCall {
	call := new(GetCategoriesCall)
	call.XMLNS = XMLNS
	return call
}

func (call *GetCategoriesCall) Do(c *CriteoClient) *GetCategoriesResponse {
	data := NewCriteoRequestCall(NewCriteoRequestCallBody(call), NewCriteoRequestCallHeader(c)).Do()
	v := new(GetCategoriesResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type GetCategoriesResponse struct {
	XMLName    xml.Name    `xml:"getCategoriesResponse"`
	Categories []*Category `xml:"getCategoriesResult>category"`
}

// Schedule Report Job

type ScheduleReportJobCall struct {
	XMLName   xml.Name   `xml:"scheduleReportJob"`
	XMLNS     string     `xml:"xmlns,attr"`
	ReportJob *ReportJob `xml:"reportJob"`
}

func NewScheduleReportJobCall(r *ReportJob) *ScheduleReportJobCall {
	call := new(ScheduleReportJobCall)
	call.XMLNS = XMLNS
	call.ReportJob = r
	return call
}

func (call *ScheduleReportJobCall) Do(c *CriteoClient) *ScheduleReportJobResponse {
	data := NewCriteoRequestCall(NewCriteoRequestCallBody(call), NewCriteoRequestCallHeader(c)).Do()
	v := new(ScheduleReportJobResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type ScheduleReportJobResponse struct {
	XMLName     xml.Name           `xml:"scheduleReportJobResponse"`
	JobResponse *ReportJobResponse `xml:"jobResponse"`
}

// Get job status

type GetJobStatusCall struct {
	XMLName xml.Name `xml:"getJobStatus"`
	XMLNS   string   `xml:"xmlns,attr"`
	JobID   int      `xml:"jobID"`
}

func NewGetJobStatusCall(r *ReportJobResponse) *GetJobStatusCall {
	call := new(GetJobStatusCall)
	call.XMLNS = XMLNS
	call.JobID = r.JobID
	return call
}

func (call *GetJobStatusCall) Do(c *CriteoClient) *GetJobStatusResponse {
	data := NewCriteoRequestCall(NewCriteoRequestCallBody(call), NewCriteoRequestCallHeader(c)).Do()
	v := new(GetJobStatusResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type GetJobStatusResponse struct {
	XMLName   xml.Name `xml:"getJobStatusResponse"`
	JobStatus string   `xml:"getJobStatusResult"`
}

// Get report download url

type GetReportDownloadUrlCall struct {
	XMLName xml.Name `xml:"getReportDownloadUrl"`
	XMLNS   string   `xml:"xmlns,attr"`
	JobID   int      `xml:"jobID"`
}

func NewGetReportDownloadUrlCall(r *ReportJobResponse) *GetReportDownloadUrlCall {
	call := new(GetReportDownloadUrlCall)
	call.XMLNS = XMLNS
	call.JobID = r.JobID
	return call
}

func (call *GetReportDownloadUrlCall) Do(c *CriteoClient) *GetReportDownloadUrlResponse {
	data := NewCriteoRequestCall(NewCriteoRequestCallBody(call), NewCriteoRequestCallHeader(c)).Do()
	v := new(GetReportDownloadUrlResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type GetReportDownloadUrlResponse struct {
	XMLName xml.Name `xml:"getReportDownloadUrlResponse"`
	JobURL  string   `xml:"jobURL"`
}

// Get report from url

type GetReportCall struct {
	XMLName xml.Name `xml:"getReportDownloadUrl"`
	XMLNS   string   `xml:"xmlns,attr"`
	JobID   int      `xml:"jobID"`
}

func NewReportCollection(url string) *GetReportResponse {
	data := HttpGetRequest(url)
	v := new(GetReportResponse)
	if err := xml.Unmarshal(data, v); err != nil {
		log.Fatalf("Error when unmarshal: %v", err)
	}
	return v
}

type GetReportResponse struct {
	XMLName xml.Name     `xml:"report"`
	Rows    []*ReportRow `xml:"table>rows>row"`
}
