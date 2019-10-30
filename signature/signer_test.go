package signature

import "testing"

func TestSigner(t *testing.T) {

	// appends params
	var data = map[string]string{}
	data["k"] = "key"
	data["s"] = "skey"
	data["m"] = "POST"
	data["t"] = "1569468833"


	data["param1"] = "1569468833"

	signstring,_ := Signer(data)

	t.Log(signstring)
}
