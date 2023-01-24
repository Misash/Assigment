package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"net/http"
	"os"
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

func getFeeRecipient(pubkey string) string {

	req, err := http.NewRequest("GET", "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net/relay/v1/data/validator_registration?pubkey="+pubkey, nil)
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

	// no fee recipient found
	if data.Message.FeeRecipient == "" {
		return "Nil"
	}

	return data.Message.FeeRecipient
}

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://xxx:yyy@157.90.93.245:5432")

	if err != nil {
		panic(err)
	}

	rows, err := conn.Query(context.Background(),
		"select f_validator_key from t_oracle_validator_balances limit 100")

	if err != nil {
		panic(err)
	}

	// create txt
	file, err := os.Create("pubkey_FeeRecipient.txt")

    // write data 
	for rows.Next() {
		var pubkey string
		err = rows.Scan(&pubkey)
		if err != nil {
			panic(err)
		}

		fee_recipient := getFeeRecipient(pubkey)
		// fmt.Println(pubkey)
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString("\n" + pubkey + "\t" + fee_recipient)
			fmt.Println("Done")
		}

	}

	//close txt
	file.Close()

}
