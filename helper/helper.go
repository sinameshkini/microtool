package helper

import "encoding/json"

func MyAssertion(src interface{}, dst interface{}) (err error) {
	var js []byte

	js, err = json.Marshal(src)
	if err != nil {
		return
	}

	err = json.Unmarshal(js, dst)
	if err != nil {
		return
	}

	return
}

