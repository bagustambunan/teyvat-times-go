package models

type JSON struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}
