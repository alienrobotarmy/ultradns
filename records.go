package ultradns

import (
	"errors"
	"net/http"
)

type RecordsResourceSets struct {
	OwnerName string   `json:"ownerName"`
	RRType    string   `json:"rrtype"`
	TTL       int      `json:"ttl"`
	Rdata     []string `json:"rdata"`
}
type Records struct {
	ZoneName string                `json:"zoneName"`
	RRSets   []RecordsResourceSets `json:"rrSets"`
}

func getRecords(OAuth OAuthResponse, zone string, recordType string) (Records, error) {
	var r Records
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://restapi.ultradns.com/v2/zones/"+
		zone+"/rrsets/"+recordType, nil)
	req.Header.Add("Authorization", "Bearer "+OAuth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return r, err
	}
	if resp.StatusCode != 200 {
		return r, errors.New("Non 200 Status Code")
	}
	err = DecodeResponse(resp, &r)
	if err != nil {
		return r, err
	}
	return r, err
}
