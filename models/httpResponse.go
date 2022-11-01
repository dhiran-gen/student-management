package models

type HttpErrs struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errorMsg"`
}

type HttpResponse struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}
