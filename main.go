package main

import (
	"fmt"
	"test/hari/model"
)

func main() {
	fmt.Println("go running in local...")
	//model.UpdateStatus()
	//url := "http://www.google.com"
	//model.CallApiInGo(url)
	model.SendMail()
}
