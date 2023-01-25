
# Assignment 

## tasks

- [x] obtener f_validator_keys
- [x]  obtener fee_recipient
- [x]  obtener un txt con los pubkeys - fee_recipient
- [x]  Almacenar el mapping fee_recipient/pubkey/timestampt en una tabla nueva. t_mapping_fee__recipient


## Run locally

```sh
  $ go env GO111MODULE
  $ go mod init main.go
  $ go mod tidy 
  $ go run .
```

## Database 

```sh
  $ psql -h localhost -d mydb -U <username> -p 5432
  $ create database mydb;
  $ \c mydb

  $ create table  t_mapping_fee__recipient(
      fee_recipient varchar(50),
      pubkey  varchar(99),  
      timestampt varchar(32)
    );
```


![image](https://user-images.githubusercontent.com/70419764/214428878-6f2031f0-c0c1-489d-b01b-fc3d33aadc95.png)
