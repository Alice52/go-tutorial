package sdk

import (
	"log"
	"net/http"
	"os"
)

func init() {
	logFileLocation, _ := os.OpenFile("./test.log",
		os.O_CREATE|
			os.O_APPEND|
			os.O_RDWR,
		0744)
	log.SetOutput(logFileLocation)
}

func HttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		defer resp.Body.Close()
	}
}
