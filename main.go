package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)



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


	for rows.Next() {

		var pubkey string
		err = rows.Scan(&pubkey)
		if err != nil {
			panic(err)
		}
		data := getData(pubkey)
		
		if err != nil {
			fmt.Println(err)
		} else
		{
			//write to txt
			file.WriteString(pubkey + "\t" + data.Message.FeeRecipient + "\n")

			//insert to DB 
			// if data.Message.FeeRecipient != "" {
			// 	insertDataMyDB(pubkey,data.Message.FeeRecipient,data.Message.Timestamp)
			// 	fmt.Println("insert!!")
			// }

			fmt.Println("Done")
		}

	}

	//close txt
	file.Close()

}
