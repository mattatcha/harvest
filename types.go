package harvest

import (
	"net/http"
	"net/url"
)

type Project struct {
	ID       int
	Name     string
	Code     string
	Billable bool

	Client   string
	ClientID int `json:"client_id"`

	Tasks []Task
}

type Task struct {
	ID       int
	Name     string
	Billable bool
}

type Daily struct {
	Projects []Project
}

type HarvestClient struct {
	encodedAuth string
	baseURL     *url.URL
	client      *http.Client
}
