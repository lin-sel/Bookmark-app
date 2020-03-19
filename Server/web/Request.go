package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UnmarshalJSON checks for empty body and then parses JSON into the target
func UnmarshalJSON(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return errors.New("There is problem while reading data")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// microLog.Logger.Errorf("%#v", err)
		return errors.New("Data can't handle")
	}

	if len(body) == 0 {
		return errors.New("Empty Data")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		// microLog.Logger.Errorf("%#v", err)
		// microLog.Logger.Printf("error decoding request: %v", err)
		// if e, ok := err.(*json.SyntaxError); ok {
		// 	log.Printf("syntax error at byte offset %d", e.Offset)
		// }
		// microLog.Logger.Printf("request: %q", body)
		fmt.Println(err.Error())
		return errors.New("Unable to Parse Data")
	}
	return nil
}