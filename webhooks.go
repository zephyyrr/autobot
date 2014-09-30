package main

import (
	"net/url"
	"time"
)

type PingEvent struct {
	Zen     string
	Hook_id string
	Hook    Hook
}

type Hook struct {
	Id         int
	URL        url.URL
	Updated_at time.Time
	Created_at time.Time
	Name       string
	Events     []string
	Active     bool
	Config     struct {
		URL          url.URL
		Content_type string
	}
}

type DeploymentEvent struct {
	SHA         string
	Name        string
	Payload     string
	Environment string
	Description string
}

type PushEvent struct {
	Head    string
	Ref     string
	Size    int
	Commits []Commit
}

type Commit struct {
	SHA      string
	Message  string
	Author   Author
	URL      url.URL
	Distinct bool
}

type Author struct {
	Name  string
	Email string
}
