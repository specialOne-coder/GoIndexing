package main

import (
	//"net/http"
	"fmt"
	"net/http"
	"sync"
	// "net/http"
	// "sync"
	// "time"
)

func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Server is running on port 8686")
		http.HandleFunc("/download", downloadFile)
		http.HandleFunc("/getDatas", getDatas)
		http.HandleFunc("/getFiles", getFileNames)
		http.HandleFunc("/insertData", insertDataInDBandFile)
		http.ListenAndServe(":8686", nil)
	}()
	wg.Wait()
	fmt.Println("End API")
}
