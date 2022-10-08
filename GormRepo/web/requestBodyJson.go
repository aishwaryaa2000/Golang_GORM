package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// UnmarshalJSON checks for empty body and then parses JSON into the target
func UnmarshalJSON(r *http.Request, target interface{}) error {

	if r.Body == nil{
		err := errors.New("request body is nil")
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return (err)
	}

	if len(body) == 0 {
		err := errors.New("request body is nil")
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}
	return nil
}