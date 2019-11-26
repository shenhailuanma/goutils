package gogithub

import (
	"fmt"
	"testing"
)

var testUsername = "shenhailuanma"
var testPassword = "" // password or personal access token

func TestGogithub_PullRequests(t *testing.T) {

	InitDefault(testUsername, testPassword)

	requests, err := PullRequests("toolkit", "")
	if err != nil {
		t.Error("error:", err.Error())
	}

	for _, one := range requests {
		fmt.Println("requst:", one)
	}
}

func TestGogithub_PullRequestsFilter(t *testing.T) {
	InitDefault(testUsername, testPassword)

	requests, err := PullRequestsFilter("toolkit", "", "shenhailuanma:test", "master", "all")
	if err != nil {
		t.Error("error:", err.Error())
	}

	for _, one := range requests {
		fmt.Println("requst:", one)
	}
}

func TestGogithub_PullRequestListFiles(t *testing.T) {
	InitDefault(testUsername, testPassword)

	files, err := PullRequestListFiles("toolkit", "", 1)
	if err != nil {
		t.Error("error:", err.Error())
	}

	for _, one := range files {
		fmt.Println("requst:", one)
	}
}

func TestGogithub_PullRequestUpdate(t *testing.T) {
	InitDefault(testUsername, testPassword)

	err := PullRequestUpdate("toolkit", "", 1, "open")
	if err != nil {
		t.Error("error:", err.Error())
	}
}