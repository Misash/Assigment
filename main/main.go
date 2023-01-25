package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

// func storeUniqueFeeRecipients(fee_recipient string){

// }

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://xxx:yyy@157.90.93.245:5432")

	if err != nil {
		panic(err)
	}

	rows, err := conn.Query(context.Background(),
		"select distinct f_validator_key from t_oracle_validator_balances")

	if err != nil {
		panic(err)
	}

	// create txt
	file, err := os.Create("pubkey_FeeRecipient.txt")
	file2, err := os.Create("uniqueFeeRecipient.txt")

	keys := make(map[string]bool)
	// uniqueFeeRecipients := []string{}

	for rows.Next() {

		var pubkey string
		err = rows.Scan(&pubkey)
		if err != nil {
			panic(err)
		}

		// get data from Request
		data := getData(pubkey)

		if err != nil {
			fmt.Println(err)
		} else {

			feeRecipient := data.Message.FeeRecipient
			//write to txt
			file.WriteString(pubkey + "\t" + feeRecipient + "\n")

			if feeRecipient != "" {
				
				//store unique FeeRecipients
				if keys[feeRecipient] == false  {
					//Print unique Fee Recipient
					fmt.Println(feeRecipient)
					keys[feeRecipient] = true
					file2.WriteString(feeRecipient + "\n")
				}

				//insert to databse
				// insertDataMyDB(pubkey,data.Message.FeeRecipient,data.Message.Timestamp)
			}
			
		}

	}

	//close txt
	file.Close()
	file2.Close()

}
