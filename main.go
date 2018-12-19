package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var flagvar int

func main() {

	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		log.Infof("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		log.Infof(string(data))
	}

	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Infof("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		log.Infof(string(data))
	}

	flag.IntVar(&flagvar, "ip", 1234, "Ip Value")

	flag.Parse()

	log.Infof("ip: %v   ", flagvar)
}
