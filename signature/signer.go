package signature

import (
	"bytes"
	"github.com/shenhailuanma/goutils/md5"
	"sort"
)


func Signer(params map[string]string) (string, error) {

	var keys = []string{}
	var stringToSignBuilder = bytes.Buffer{}

	// get params keys
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// genrate sign base string
	for _, key := range keys {
		stringToSignBuilder.WriteString(key + ":" + params[key])
	}

	//fmt.Println("base string:", stringToSignBuilder.String())

	return md5.Md5Bytes(stringToSignBuilder.Bytes()), nil
}
