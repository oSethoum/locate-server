package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jackpal/gateway"
)

var c chan string
var ports = []string{"3000", "3001", "3003"}

func ping(url string) {
	res, err := http.Get(url)
	if err == nil {
		if res.StatusCode == 200 {
			c <- url
		}
		res.Body.Close()
	}
	c <- ""
}

func main() {
	ip, err := gateway.DiscoverGateway()
	if err != nil {
		println(err.Error())
	}
	ur := strings.Join(strings.Split(ip.String(), ".")[:3], ".")
	c = make(chan string)
	for i := 2; i < 254; i++ {
		for _, p := range ports {
			url := fmt.Sprintf("http://%s.%d:%s/ping", ur, i, p)
			go ping(strings.Split(url, "/ping")[0])
		}
	}

	if url := <-c; url != "" {
		fmt.Printf(url)
		return
	}
	print("Offline")
}
