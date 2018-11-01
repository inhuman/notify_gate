package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseRequest is used for parsing http.Request into given interface
func ParseRequest(r *http.Request, object interface{}) error {
	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, &object)

	if err != nil {
		return err
	}
	return nil
}
