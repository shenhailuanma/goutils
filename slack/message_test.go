package slack

import (
	"testing"
)

var token = ""

func TestSlackObj_PostTextMessage(t *testing.T) {

	InitDefault(token)

	err := PostTextMessage("test", "test message")
	if err != nil {
		t.Error("error:", err.Error())
	}
}