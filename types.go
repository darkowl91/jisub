package main

import "net/http"

type Jira struct {
	BaseUrl string
	Client  *http.Client
	Auth    *Auth
}

type Auth struct {
	BearerToken string
}

type Issue struct {
	Id     string       `json:"id"`
	Key    string       `json:"key"`
	Self   string       `json:"self"`
	Fields *IssueFields `json:"fields"`
}

type Issues struct {
	Issues []*Issue `json:"issues"`
}

type IssueFields struct {
	Summary           string       `json:"summary"`
	SubTasks          []*Issue     `json:"subtasks"`
	Status            *IssueStatus `json:"status"`
	IssueType         *IssueType   `json:"issuetype"`
	Project           *JiraProject `json:"project"`
	CustomField_26606 string       `json:"customfield_26606"` // deal_size
	CustomField_10004 float32      `json:"customfield_10004"` // story_points
}

type IssueStatus struct {
	Name string `json:"name"`
}

type IssueType struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Subtask bool   `json:"subtask"`
}

type JiraProject struct {
	Id string `json:"id"`
}
