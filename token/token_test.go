package token

import (
	"fmt"
	"testing"
)

const testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODkxOTI4MzUsImlhdCI6MTU4ODU4ODAzNSwiaXNzIjoiY2xvdWR3YXkiLCJzdWIiOiI0NDUwNjE4NzVAcXEuY29tIiwia2luIjozLCJhaWQiOjMwLCJwb3J0Ijo5MDAwLCJzbHQiOjF9.DRKoN_VRRN6V7buHTMzU9t56RdEFRHZVO_tSaKP-_ek"

func Test_ParseTokenWithClaimsUnverified(t *testing.T) {
	tokenInfo, ok := ParseTokenWithClaimsUnverified(testToken)
	if ok {
		fmt.Println("tokenInfo:", tokenInfo)
	} else {
		t.Error("parse failed")
	}
}