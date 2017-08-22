package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"sync"
	"time"
	"bytes"
	"strconv"
	"flag"
)

func main()  {

	flag.Parse()
	s := flag.Arg(0)
	url := flag.Arg(1)
	numMessage, _ := strconv.Atoi(s)
	var wg sync.WaitGroup
	wg.Add(11)
	fmt.Println("Start sending request")
	start := time.Now()
	req := (numMessage / 10) + 1
	for i:=0; i<=10; i++ {
		go MakeRequest(url, &wg, req)
	}
	wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func MakeRequest(url string, wg *sync.WaitGroup, numMessage int) {
	defer wg.Done()
	for k:=0; k<= numMessage; k++ {
		user := rand.Int63n(10000)
		brand := rand.Int63n(10000)
		reqBody := fmt.Sprintf(`{"records":[{"value":{"user":%d, "brand":%d}}]}`, user, brand)
		var jsonStr = []byte(reqBody)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Accept", "application/vnd.kafka.v2+json")
		req.Header.Set("Content-Type", "application/vnd.kafka.json.v2+json")
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("ERROR: %s", err)
			panic(err)
		}
		defer resp.Body.Close()
	}
}


