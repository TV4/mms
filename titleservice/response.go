package titleservice

// Response is used for responses from the MMS TitleService API
type Response struct {
	StatusCode        int      `json:"StatusCode"`
	StatusDescription string   `json:"StatusDescription"`
	Errors            []string `json:"Errors"`
}
