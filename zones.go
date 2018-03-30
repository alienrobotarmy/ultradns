package ultradns

import (
	_ "bufio"
	"errors"
	"fmt"
	"net/http"
)

type ZonesResultInfo struct {
	TotalCount    int `json:"totalCount"`
	Offset        int `json:"offset"`
	ReturnedCount int `json:"returnedCount"`
}
type ZonesListProperties struct {
	Name                 string `json:"name"`
	AccountName          string `json:"accountName"`
	Owner                string `json:"owner"`
	Type                 string `json:"type"`
	RecordCount          int    `json:"recordCount"`
	DnssecStatus         string `json:"dnssecStatus"`
	LastModifiedDateTime string `json:"lastModifiedDateTime"`
}
type ZonesListNameServers struct {
	Ok []string `json:"ok"`
}
type ZonesListRegistrarInfo struct {
	NameServers ZonesListNameServers `json:"nameServers"`
}
type ZonesList struct {
	Properties    ZonesListProperties    `json:"properties"`
	RegistrarInfo ZonesListRegistrarInfo `json:"registrarInfo"`
}
type Zones struct {
	ResultInfo ZonesResultInfo `json:"resultInfo"`
	List       []ZonesList     `json:"zones"`
}

func GetZones(OAuth OAuthResponse) (Zones, error) {
	client := &http.Client{}
	var z Zones
	//resp, err := client.Get("https://restapi.ultradns.com/v2/zones")

	req, err := http.NewRequest("GET", "https://restapi.ultradns.com/v2/zones", nil)
	req.Header.Add("Authorization", "Bearer "+OAuth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return z, err
	}
	if resp.StatusCode != 200 {
		return z, errors.New("Non 200 Status Code")
	}
	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
	err = DecodeResponse(resp, &z)
	if err != nil {
		return z, err
	}
	return z, nil
}
