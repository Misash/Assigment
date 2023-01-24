package main

import (
	// "context"
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "golang.org/x/tools/go/analysis/passes/nilfunc"
	// "github.com/jackc/pgx/v4"
)


type Data struct{
    Message struct{
        FeeRecipient string `json:"fee_recipient"`
        GasLimit string `json:"gas_limit"`
        Timestamp string `json:"timestamp"`
        PubKey string `json:"pubkey"`
    }
    Signature string `json:"signature"`
}

// type Message struct{
//     FeeRecipient string
//     GasLimit string
//     Timestamp string
//     PubKey string
// }

func getFeeRecipient(pubkey string) string{

	req, err := http.NewRequest("GET", "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net/relay/v1/data/validator_registration?pubkey="+pubkey, nil)
	if err != nil {
        panic(err)
        fmt.Println("1")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
        panic(err)
        fmt.Println("2")
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
        fmt.Println("3")
	}

    // data := map[string]interface{}{}
    //to jason
    data := Data {}
    json.Unmarshal(bodyText,&data)

    if data.Message.FeeRecipient == "" {
        return "**"
    }else{
        return data.Message.FeeRecipient 
    }

}



func main() {

	// conn, err := pgx.Connect(context.Background(), "postgres://xxx:yyy@157.90.93.245:5432")

	// if err != nil {
	// 	panic(err)
	// }

	// rows, err := conn.Query(context.Background(),
	// 	"select f_validator_key from t_oracle_validator_balances limit 5")

	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	var pubkeys string
	// 	err = rows.Scan(&pubkeys)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// fmt.Println(pubkeys)
	// }



    pbkey := "0x9e1b930f03cca8898e255b0ce7ca196b94482956c9f8eca9db425caa7897db44116470b1fb8f08c959a9e623e3003a8d"

    fmt.Println("_"+getFeeRecipient(pbkey)+"_")


	// fmt.Printf("%s\n", data)
    // fee := hex.DecodeString(data)
    // fmt.Println(data["message"].(map[string]interface{})["fee_recipient"])





    // fmt.Printf("\n%T\n",data.Signature)
    

}
