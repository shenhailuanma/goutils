package gogithub

import (
	"testing"
)

func TestGogithub_TriggerRepositoryDispatch(t *testing.T) {

	InitDefault(configUsername, configPassword)

	err := TriggerRepositoryDispatch(configRepo, configOwner, configRepositoryDispatchEventType, "")
	if err != nil {
		t.Error("error:", err.Error())
	}
}