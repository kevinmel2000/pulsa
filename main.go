package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/nyelonong/gommon/jsonapi"
)

const (
	TBANK_URL string = "http://hackathon.bri.co.id"
)

type Pulsa struct {
	Status   string `json:"status"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Sender   string `json:"sender"`
}

type Template struct {
	Amount string
}

func main() {
	http.HandleFunc("/sms", func(res http.ResponseWriter, req *http.Request) {
		api := jsonapi.New(res, "")

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			api.ErrorWriter(http.StatusBadRequest, err.Error())
			return
		}

		log.Println(string(body))

		var p Pulsa
		if err := json.Unmarshal(body, &p); err != nil {
			log.Println(err)
			api.ErrorWriter(http.StatusBadRequest, err.Error())
			return
		}

		log.Println(p, p.Amount)

		msg := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
        <soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
          <soap12:Body>
            <TransferTBank xmlns="http://hackathon.bri.co.id/">
              <nohandphonePengirim>081210039328</nohandphonePengirim>
              <nohandphonePenerima>08980900120</nohandphonePenerima>
              <pin>133164</pin>
              <nominal>%s</nominal>
            </TransferTBank>
          </soap12:Body>
        </soap12:Envelope>
        `, p.Amount)

		log.Println(msg)

		if err := payment(msg); err != nil {
			log.Println(err)
			return
		}

		api.SuccessWriter(p)
	})

	http.HandleFunc("/register/merchant", handle(REGISTER_MERCHANT))
	http.HandleFunc("/register/tbank", handle(REGISTER_TBANK1))
	http.HandleFunc("/topup", handle(TOPUP))
	http.HandleFunc("/transfer", handle(TRANSFER))

	log.Fatal(http.ListenAndServe(":31337", nil))
}

func handle(body string) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req, err := http.NewRequest("POST", TBANK_URL+"/BRIHackathon.asmx", strings.NewReader(body))
		req.Method = "POST"

		if err != nil {
			log.Println(err)
			return
		}

		req.Header.Add("Content-Type", "application/soap+xml; charset=utf-8")
		client := http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(b))
	}
}

func payment(body string) error {
	req, err := http.NewRequest("POST", TBANK_URL+"/BRIHackathon.asmx", strings.NewReader(body))
	if err != nil {
		log.Println(err)
		return err
	}

	req.Method = "POST"
	req.Header.Add("Content-Type", "application/soap+xml; charset=utf-8")
	client := http.Client{}

	log.Printf("%+v\n", req)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("%+v\n", resp)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(string(b))
	return nil
}
