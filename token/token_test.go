package token

import (
	"fmt"
	"testing"
)

const testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzE1NTQ5NzgsImlhdCI6MTYzMTUxMTc3OCwiaXNzIjoibGlua3NjYWxlIiwic3ViIjoidG9taXRvbnkxMjNAb3V0bG9vay5jb20iLCJhaWQiOjEsImNpZCI6Mywicm9sZXMiOlsiTGlua3NjYWxlLVRvcEFkbWluIiwiTGlua3NjYWxlLVNhbGVzQWNjb3VudE1hbmdlciIsIkRlbGVnYXRpb24tNjUtU2FsZXNBY2NvdW50TWFuZ2VyIl19.y6D3t-udRBT4IQdVPoXHjW2-7NEaJ_4r1PJfQbbSA_s"
func Test_ParseTokenWithClaimsUnverified(t *testing.T) {
	tokenInfo, ok := ParseTokenWithClaimsUnverified(testToken)
	if ok {
		fmt.Println("tokenInfo:", tokenInfo)
		fmt.Println("tokenInfo PID:", tokenInfo.PID)
	} else {
		t.Error("parse failed")
	}
}