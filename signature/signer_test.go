package signature

import "testing"

func TestSigner(t *testing.T) {

	// appends params
	var data = map[string]string{}
	data["k"] = "stalker"
	data["s"] = "465424e7e962795c8f7d841f5fa13c5a"
	data["m"] = "POST"
	data["t"] = "1569468833"


	data["tt"] = "1569468833"

	signstring,_ := Signer(data)

	t.Log(signstring)
}
