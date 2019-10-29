package random

import "testing"

func Test_Random(t *testing.T)  {
	t.Log(RandomString(6))
	t.Log(RandomString(10))
	t.Log(RandomString(16))

	t.Log(RandomStringBase64(6))
	t.Log(RandomStringBase64(10))
	t.Log(RandomStringBase64(16))

	t.Log(RandomStringBase36(6))
	t.Log(RandomStringBase36(10))
	t.Log(RandomStringBase36(16))

	t.Log(RandomNumberString(6))
	t.Log(RandomNumberString(10))
	t.Log(RandomNumberString(16))
}
