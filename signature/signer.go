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

	// sort in increasing order
	sort.Strings(keys)

	// connect sign base string. key and value connect with ':'
	// format:
	// key1:value1key2:value2key3:value3
	for _, key := range keys {
		stringToSignBuilder.WriteString(key + ":" + params[key])
	}

	return md5.Md5Bytes(stringToSignBuilder.Bytes()), nil
}
