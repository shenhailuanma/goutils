package gogithub

import (
	"fmt"
	"testing"
)

func TestGogithub_PullRequests(t *testing.T) {

	InitDefault(configUsername, configPassword)

	requests, err := PullRequests(configRepo, configOwner)
	if err != nil {
		t.Error("error:", err.Error())
	}

	for _, one := range requests {
		fmt.Println("requst:", one)
	}
}

func TestGogithub_PullRequestsFilter(t *testing.T) {
	InitDefault(configUsername, configPassword)

	requests, err := PullRequestsFilter(configRepo, configOwner, "shenhailuanma:test", "master", "all")
	if err != nil {
		t.Error("error:", err.Error())
	}

	for _, one := range requests {
		fmt.Println("requst:", one)
	}
}

func TestGogithub_PullRequestListFiles(t *testing.T) {
	InitDefault(configUsername, configPassword)

	files, err := PullRequestListFiles(configRepo, configOwner, 1)
	if err != nil {
		t.Error("error:", err.Error())
	}

	for _, one := range files {
		fmt.Println("requst:", one)
	}
}

func TestGogithub_PullRequestUpdate(t *testing.T) {
	InitDefault(configUsername, configPassword)

	err := PullRequestUpdate(configRepo, configOwner, 1, "open")
	if err != nil {
		t.Error("error:", err.Error())
	}
}