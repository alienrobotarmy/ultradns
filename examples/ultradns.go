package main

import (
	_ "bufio"
	"errors"
	"flag"
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
	recordType := flag.String("r", "A", "Record Set Type")
	flag.Parse()

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
	// List all zones
	zones, err := ultradns.GetZones(OAuth)
	for _, v := range zones.List {
		// List all A records in each Zone
		fmt.Printf("%s\n", v.Properties.Name)
		records, _ := getRecords(OAuth, v.Properties.Name, *recordType)
		for _, r := range records.RRSets {
			for _, z := range r.Rdata {
				fmt.Printf("\t%v\n", z)
			}
		}
	}
}
