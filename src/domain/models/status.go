package models

type Status struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	Status  string `json:"status"`
	Date    string `json:"date"`
}
