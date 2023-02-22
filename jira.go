package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

const (
	subTaskBulkTmpl = "/jisub/sub-task-bulk.tmpl"
	subTaskTmpl     = "/jisub/sub-task.tmpl"
)

func NewJira(baseUrl string, auth *Auth) *Jira {

	client := &http.Client{}

	return &Jira{
		BaseUrl: baseUrl,
		Auth:    auth,
		Client:  client,
	}
}

func (j *Jira) execute(req *http.Request, result interface{}) error {

	req.Header.Set("Accept", "application/json")

	if j.Auth != nil && j.Auth.BearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+j.Auth.BearerToken)
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		if respText, err := io.ReadAll(req.Body); err == nil {
			return fmt.Errorf("response error %v %v", resp.StatusCode, respText)
		} else {
			return err
		}
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return err
		}
	}
	return nil
}

// Retrieves issue details by issue key
func (j *Jira) Issue(issueKey string) (*Issue, error) {

	url := j.BaseUrl + "/issue/" + issueKey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var issue Issue
	return &issue, j.execute(req, &issue)
}

// Creates sub-task for the issue
func (j *Jira) SubTask(parent Issue, prefix string, sp float32) (*Issue, error) {

	pwd, _ := os.Getwd()
	tmpl, err := template.ParseFiles(pwd + subTaskTmpl)
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer

	err = tmpl.Execute(&buff, struct {
		Parent      Issue
		Prefix      string
		StoryPoints float32
	}{
		Parent:      parent,
		Prefix:      prefix,
		StoryPoints: sp,
	})

	if err != nil {
		return nil, err
	}

	url := j.BaseUrl + "/issue"
	req, err := http.NewRequest("POST", url, &buff)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	var issue Issue
	return &issue, j.execute(req, &issue)
}

// Creates multiple sub-tasks for the parent issue based on work breakdown Eg. QA:2, BE:3, ect
func (j *Jira) SubTasks(parent Issue, sp map[string]float64) (*Issues, error) {

	if len(sp) == 0 {
		return &Issues{
			Issues: []*Issue{},
		}, nil
	}

	pwd, _ := os.Getwd()
	tmpl, err := template.ParseFiles(pwd + subTaskBulkTmpl)
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, struct {
		Parent      Issue
		StoryPoints map[string]float64
	}{
		Parent:      parent,
		StoryPoints: sp,
	})

	if err != nil {
		return nil, err
	}

	url := j.BaseUrl + "/issue/bulk"
	req, err := http.NewRequest("POST", url, &buff)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	var issues Issues
	return &issues, j.execute(req, &issues)
}
