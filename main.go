package main

import (
	"fmt"
	//"test/hari/model"
	"net/http"
)

func main() {
	fmt.Println("go running in local...")
	//model.UpdateStatus()
	//url := "http://www.google.com"
	//model.CallApiInGo(url)
	//model.SendMail()
	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello This is hari"))
	})
	http.ListenAndServe(":8080", nil)
}
