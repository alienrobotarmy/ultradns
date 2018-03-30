package main

import (
	_ "bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"ultradns"
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

func getRecords(OAuth ultradns.OAuthResponse, zone string, recordType string) (Records, error) {
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
	err = ultradns.DecodeResponse(resp, &r)
	if err != nil {
		return r, err
	}
	return r, err
}
func main() {
	credsBuf, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	OAuth, err := ultradns.Authenticate(string(credsBuf))
	if err != nil {
		fmt.Printf("Cannot continue: %s\n", err)
		return
	}
	zones, err := ultradns.GetZones(OAuth)
	for _, v := range zones.List {
		fmt.Printf("%s\n", v)
	}
}
