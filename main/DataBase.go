package main

import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)


func insertDataMyDB(pubkey, fee_recipient, timestamp string) {

	host := "localhost"
	port := 5432
	user := "misash"
	password := "1234"
	dbname := "mydb"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	insert into t_mapping_fee__recipient (fee_recipient,pubkey,timestampt) 
	values ($1,$2,$3)`

	
	db.QueryRow(sqlStatement, fee_recipient, pubkey, timestamp)

}