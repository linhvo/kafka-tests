package main

import (
	"fmt"
	"net/http"
	"math/rand"
	//"flag"
	"sync"
	//"time"
	"bytes"
	//"strconv"
	"io/ioutil"
)

func main()  {

	//flag.Parse()
	//s := flag.Arg(0)
	//url := flag.Arg(1)
	//numMessage, _ := strconv.Atoi(s)
	//var wg sync.WaitGroup
	//wg.Add(11)
	//fmt.Println("Start sending request")
	//start := time.Now()
	//req := (numMessage / 10) + 1
	//for i:=0; i<=10; i++ {
	//	go MakeRequest(url, &wg, req)
	//}
	//wg.Wait()
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	user := rand.Int63n(1000)
	brand := rand.Int63n(1000)
	reqBody := fmt.Sprintf(`{"records":[{"value":{"user":%d, "brand":%d}}]}`, user, brand)
	testBody := []byte(`{"records":[{"value":{"user":1, "brand": 2}}]}`)
	var jsonStr = []byte(testBody)
	req, err := http.NewRequest("POST", "http://localhost:8082/topics/jsontest", bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/vnd.kafka.v2+json")
	req.Header.Set("Content-Type", "application/vnd.kafka.json.v2+json")
	fmt.Print(reqBody)
	client := &http.Client{}
	resp, err := client.Do(req)


	if err != nil {
		fmt.Println("ERROR: %s", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}

func MakeRequest(url string, wg *sync.WaitGroup, numMessage int) {
	defer wg.Done()
	for k:=0; k<= numMessage; k++ {
		user := rand.Int63n(1000)
		brand := rand.Int63n(1000)
		reqBody := fmt.Sprintf(`{"records":[{"value":{"user":%d, "brand":%d}}]}`, user, brand)
		var jsonStr = []byte(reqBody)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Accept", "application/vnd.kafka.v2+json")
		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			panic(err)
		}

	}
}


