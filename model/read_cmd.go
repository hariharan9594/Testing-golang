package model

import (
	"fmt"
	"log"
	"os/exec"
)

func MonitorExim4OnServer() {
	result, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", string(result))
}
