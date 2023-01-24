
-- sudo -u postgres psql
-- psql -h localhost -d mydb -U misash -p 5432

create database mydb;

\c mydb

create table  t_mapping_fee__recipient(
    fee_recipient varchar(50),
    pubkey  varchar(99),  
    timestampt varchar(32)
);


