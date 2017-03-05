package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"log"
	"github.com/subosito/twilio"
)
type ApiKey struct{
	AccountSid string
	AuthToken string
}

type Message struct {
	Version string `json:"version"`
	Response map[string]string `json:"response"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		people := make(map[string]string)
		people["Anna"]="+17204404892"
		people["Jack"]="+16302355560"

		sender:="+17203304463"
		//TODO get the massage and sender, look up number in people
		reciever:="+17204404892"
		message:="hey"

		//TODO access lambda enviornment varibable for twilio api if possible
		var twilioKey ApiKey
		twilioKey.AccountSid = "ACCOUNT_SID"
		twilioKey.AuthToken  = "AUTH_TOKEN"

		// Initialize twilio client
		c := twilio.NewClient(twilioKey.AccountSid, twilioKey.AuthToken, nil)
		// Send Message
		params := twilio.MessageParams{
			Body: message,
		}
		s, response, err := c.Messages.Send(sender, reciever, params)
		if err != nil {
			log.Fatal(s, response, err)
		}

		var m Message
		var data = []byte(`{"version": "1.0", "response":{"":""}}`)
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}
		return m, nil
	})
}
