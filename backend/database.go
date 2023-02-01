package main

import (
	"database/sql"
	"encoding/json"

	// "encoding/json"
	"fmt"
	// "strconv"

	//"strconv"
	// "time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Connection to the database
func ConnectionToDB() (dbase *sql.DB, err error) {
	connectionString := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	checkErr(err)
	fmt.Println("Successfully connected!")
	return db, nil
}

func InsertIntoDBAndIntoFile(dbase *sql.DB, tokenAddress string, userAddress string) {
	sqlStat1 := "CREATE TABLE IF NOT EXISTS public.contracts(address TEXT ,total_supply TEXT,name TEXT ,symbol TEXT ,decimals TEXT,rate TEXT,currency TEXT,diff TEXT, diff7d TEXT,diff30d TEXT,market_cap_usd TEXT,available_supply TEXT,volume24h TEXT,ts TEXT,public_tags TEXT[],owner TEXT NOT NULL,count_ops TEXT ,total_in TEXT,total_out TEXT ,transfers_count TEXT,eth_transfers_count TEXT,holders_count TEXT,issuances_count TEXT ,image TEXT,description TEXT,website TEXT,last_updated TEXT, address_type TEXT,top_holders TEXT, PRIMARY KEY (address)); "
	_, errors := dbase.Exec(sqlStat1)
	checkErr(errors)
	sqlStat2 := "CREATE TABLE IF NOT EXISTS public.transactions(timestamp TEXT NOT NULL,_from TEXT NOT NULL,_to TEXT NOT NULL,hash TEXT NOT NULL,value TEXT NOT NULL,input TEXT NOT NULL,success TEXT NOT NULL,token TEXT,PRIMARY KEY (hash));"
	_, errors = dbase.Exec(sqlStat2)
	checkErr(errors)
	token, err := GetTokenDatas(tokenAddress)
	checkErr(err)
	//fmt.Println("token", token)
	// insert token into database
	_, err = dbase.Exec("INSERT INTO contracts(address,total_supply,name,symbol,decimals,rate,currency,diff,diff7d,diff30d,market_cap_usd,available_supply,volume24h,ts,public_tags,owner,count_ops,total_in,total_out,transfers_count,eth_transfers_count,holders_count,issuances_count,image,description,website,last_updated,address_type,top_holders) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29) ON CONFLICT (address) DO NOTHING", token.Address, token.TotalSupply, token.Name, token.Symbol, token.Decimals, token.Price.Rate, token.Price.Currency, token.Price.Diff, token.Price.Diff7d, token.Price.Diff30d, token.Price.MarketCapUsd, token.Price.AvailableSupply, token.Price.Volume24h, token.Price.Ts, pq.Array(token.PublicTags), token.Owner, token.CountOps, token.TotalIn, token.TotalOut, token.TransfersCount, token.EthTransfersCount, token.HoldersCount, token.IssuancesCount, token.Image, token.Description, token.Website, token.LastUpdated, token.Address_Type, pq.Array(token.TopHolder))
	checkErr(err)
	// insert transactions into database
	// var ops Operations
	ops := GetTxs(tokenAddress)
	//fmt.Println("operations in db => ", ops)
	checkErr(err)
	for i, operation := range ops {
		ops[i].Token = tokenAddress
		_, err = dbase.Exec("INSERT INTO transactions(timestamp,_from,_to,hash,value,input,success,token) VALUES($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT (hash) DO NOTHING", operation.Timestamp, operation.From, operation.To, operation.Hash, operation.Value, operation.Input, operation.Success, tokenAddress)
		checkErr(err)
	}
	filename := tokenAddress + "_" + userAddress
	InsertIntoFile(ops, filename)
	fmt.Println("Successfull insert into db and file")
}

// Get the asset pairs from the database
func GetDBTxs() (datas Operations, bytes []byte, err error) {
	dbase, err := ConnectionToDB()
	checkErr(err)
	rows, err := dbase.Query("SELECT * FROM transactions")
	checkErr(err)
	defer rows.Close()
	var dbDatas Operations

	for rows.Next() {

		var timestamp int64
		var _from string
		var _to string
		var hash string
		var value float64
		var input string
		var success bool
		var token string

		var dbData Operation

		err = rows.Scan(&timestamp, &_from, &_to, &hash, &value, &input, &success, &token)
		checkErr(err)

		dbData = Operation{timestamp, _from, _to, hash, value, input, success, token}
		dbDatas = append(dbDatas, dbData)

		fmt.Println("dbData query Succesfful")

	}
	datasByte, _ := json.MarshalIndent(dbDatas, "", "\t")
	fmt.Println("dbData query Succesfful")

	return dbDatas, datasByte, nil
}
