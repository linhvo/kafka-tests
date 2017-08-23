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
	wg.Add(15)
	fmt.Println("Start sending request")
	start := time.Now()
		req := (numMessage / 15) + 1
	sum := 0
	for i:=0; i<15; i++ {
		go MakeRequest(url, &wg, req)
		sum = sum + (req -1 )
	}
	wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf("Sum Messages: %d", sum)

}

func MakeRequest(url string, wg *sync.WaitGroup, numMessage int) {
	defer wg.Done()
	for k:=0; k<= numMessage; k++ {
		user := rand.Int63n(1000000)
		brand := rand.Int63n(1000000)
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


