package main 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


type Data struct {
	Message struct {
		FeeRecipient string `json:"fee_recipient"`
		GasLimit     string `json:"gas_limit"`
		Timestamp    string `json:"timestamp"`
		PubKey       string `json:"pubkey"`
	}
	Signature string `json:"signature"`
}


func getData(pubkey string) Data {

	req, err := http.NewRequest("GET", "https://boost-relay-goerli.flashbots.net/relay/v1/data/validator_registration?pubkey="+pubkey, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// get data from Json
	data := Data{}
	json.Unmarshal(bodyText, &data)


	return data
}