package ultradns

import (
	"encoding/json"
	"net/http"
)

func DecodeResponse(resp *http.Response, v interface{}) error {
	err := json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return err
	}
	return nil
}
