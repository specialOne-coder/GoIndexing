package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	// Ouvrez le fichier en lecture
	fileName := r.URL.Query().Get("filename")
	file, err := os.Open("Archive/" + fileName)
	if err != nil {
		http.Error(w, "Impossible d'ouvrir le fichier", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Définissez l'en-tête du fichier
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", FRONT)
	// Copiez le contenu du fichier dans la réponse
	io.Copy(w, file)
}

func getDatas(w http.ResponseWriter, r *http.Request) {
	_, bytes, err := GetDBTxs()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONT)
	w.Write(bytes)
}

func insertDataInDBandFile(w http.ResponseWriter, r *http.Request) {
	tokenAddress := r.URL.Query().Get("tokenAddress")
	userAddress := r.URL.Query().Get("userAddress")
	fmt.Println("Request Body", tokenAddress)
	db, err := ConnectionToDB()
	checkErr(err)
	InsertIntoDBAndIntoFile(db, tokenAddress, userAddress)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONT)
}

func getFileNames(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get file names")
	tokenAddress := r.URL.Query().Get("tokenAddress")
	userAddress := r.URL.Query().Get("userAddress")
	files, err := ioutil.ReadDir("Archive/")
	checkErr(err)

	var fileNames []string
	for _, file := range files {
		if file.Name() == "txs_"+tokenAddress+"_"+userAddress+".json" {
			fileNames = append(fileNames, file.Name())
		}
	}

	json, err := json.Marshal(fileNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", FRONT)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
