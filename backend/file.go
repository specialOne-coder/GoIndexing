package main

import (
	"encoding/json"
	"os"
)

// InsertIntoFile insère les données dans un fichier JSON dans le directory Archive
func InsertIntoFile(operations Operations, fileName string) {
	dir := "Archive"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir("Archive", 0755)
		checkErr(err)
	}
	file, err := os.Create("Archive/txs_" + fileName + ".json")
	checkErr(err)
	defer file.Close()
	// Encodage de l'objet en JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(operations); err != nil {
		panic(err)
	}
}
