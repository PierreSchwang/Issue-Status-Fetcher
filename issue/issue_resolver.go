package issue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Component struct {
	Title  string  `json:"title"`
	Labels []Label `json:"labels"`
}

func (c Component) IsOperational() bool {
	for _, v := range c.Labels {
		if v.Name == "operational" {
			return true
		}
	}
	return false
}

func (c Component) IsSubComponent() bool {
	for _, v := range c.Labels {
		if v.Name == "subcomponent" {
			return true
		}
	}
	return false
}

func (c Component) GetStatus() Label {
	for _, v := range c.Labels {
		if v.Name == "component" || v.Name == "issue status" || v.Name == "subcomponent" {
			continue
		}
		return v
	}
	return Label{
		Name:  "Invalid Status",
		Color: "FFFFFF",
	}
}

const COMPONENTS = "https://api.github.com/repos/%s/issues?state=all&labels=issue%sstatus,component"
const INCIDENTS = "https://api.github.com/repos/%s/issues?state=all&labels=issue status,component"

var httpClient *http.Client

func Components(repo string, githubAccessToken string) ([]Component, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(COMPONENTS, repo, "%20"), nil)
	if githubAccessToken != "" {
		req.Header.Set("Authorization", "token "+githubAccessToken)
	}
	data, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return nil, err
	}
	var components []Component
	err = json.Unmarshal(body, &components)
	if err != nil {
		log.Println(string(body))
		return nil, err
	}
	return components, nil
}
