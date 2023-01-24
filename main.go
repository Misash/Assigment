package main

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v4"
)

func main() {

    conn, err := pgx.Connect(context.Background(), "postgres://xxx:yyy@157.90.93.245:5432")
	
    if err != nil {
        panic(err)
    }


    rows , err := conn.Query(context.Background(),
        // "select f_validator_key from t_oracle_validator_balances order by f_validator_key desc ").Scan(&depositAddress)
		"select f_validator_key from t_oracle_validator_balances limit 10")

	
    if err != nil {
        panic(err)
    }

	for rows.Next(){
		var depositAddress string
		err = rows.Scan(&depositAddress)
		if err != nil {
			panic(err)
		}
		fmt.Println(depositAddress)
	}


    // fmt.Println(rows)

}