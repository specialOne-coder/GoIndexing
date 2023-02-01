package main

import (
	"encoding/json"
	//"fmt"
	//"io/ioutil"

	// "go/token"

	//"fmt"
	// "io/ioutil"
	"net/http"
	// "sync"
	//"time"
)

func GetTokenDatas(tokenAddress string) (TokenInfo, error) {

	// Get token holders
	resh, err := http.Get(ethplorerAPI + "getTopTokenHolders/" + tokenAddress + "?apiKey=" + apiKey)
	checkErr(err)
	defer resh.Body.Close()
	var tokenHolders Holders
	err = json.NewDecoder(resh.Body).Decode(&tokenHolders)
	checkErr(err)
	//fmt.Println("Holders", tokenHolders)
	// Requête HTTP vers l'API de Kraken
	res, err := http.Get(ethplorerAPI + "getTokenInfo/" + tokenAddress + "?apiKey=" + apiKey)
	checkErr(err)
	defer res.Body.Close()

	// Décodage de la réponse JSON
	var tokenInfo TokenInfo
	err = json.NewDecoder(res.Body).Decode(&tokenInfo)
	checkErr(err)
	if tokenInfo.Address == "" {
		tokenInfo.Address_Type = "Wallet"
	} else {
		tokenInfo.Address_Type = "Contract"
	}
	for i := 0; i < len(tokenHolders.Holders); i++ {
		tokenInfo.TopHolder = append(tokenInfo.TopHolder, tokenHolders.Holders[i].Address)
	}
	//fmt.Println(tokenInfo)
	return tokenInfo, err
}

func GetTxs(tokenAddress string) Operations {

	// Requête HTTP vers l'API de Kraken
	res, err := http.Get(ethplorerAPI + "getAddressTransactions/" + tokenAddress + "?apiKey=" + apiKey + "&limit=10")
	checkErr(err)
	defer res.Body.Close()
	//fmt.Println("resultat ", res)
	//body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	// Décodage de la réponse JSON
	var operations Operations
	err = json.NewDecoder(res.Body).Decode(&operations)
	//json.Unmarshal(body, &operations)
	checkErr(err)
	//fmt.Println("Operations", operations)
	return operations
}
