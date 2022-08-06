package model

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Wait struct {
	*sync.WaitGroup
}

func Dispaly(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			i = i + 1
			fmt.Printf("Ticker %v min completed..\n", i)
		}
	}
}

func (wg *Wait) Dispaly(url string) {
	defer wg.Done()
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("website is down")
	} else {
		fmt.Println("website is up")
	}
}

func UpdateStatus() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go Dispaly(wg)
	wg.Wait()
}

func UpdateExim4Status(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			i = i + 1
			fmt.Printf("Ticker %v sec completed..\n", i)
		}
	}
}

func MonitorServer() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go UpdateExim4Status(wg)
	wg.Wait()
}

func CallApiInGo(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("body: ", res.Body)
}
