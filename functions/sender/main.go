package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"github.com/subosito/twilio"
	"os"
	"fmt"
//	"errors"
)
type ApiKey struct{
	AccountSid string
	AuthToken string
}

type LambdaRequest struct{
	Session  interface{} `json:"session"`
	Request  Request `json:"request"`
	Version  string `json:"version"`
}
type Request struct{
	Type string `json:"type"` // IntentRequest, SessionEndedRequest, LaunchRequest
	RequestId string `json:"requestId"`
	Locale string `json:"locale"`
	TimeStamp string `json:"timestamp"`
	Intent Intent `json:"intent,omitempty"` // IntentRequest param
	Error Error `json:"error, omitempty"` // SessionEndedRequest param
	Reason string `json:"reason, omitempty"` // SessionEndedRequest param
}

type Intent struct{
	Name string `json:"name"`
	Slots map[string]Slot `json:"slots"`
}
type Slot struct{
	Name string `json:"name"`
	Value string `json:"value"`
}
type Error struct{
	Type string `json:"type"`
	Message string `json:"message"`
}
type ResponseBody struct {
	Version string `json:"version"` 
	Response Response `json:"response"`
}
type Response struct{
	OutputSpeech map[string]string `json:"outputSpeech"`
}
func NewResponse() *ResponseBody {
	re:=new(ResponseBody)
	re.Version ="1.0"
	re.Response.OutputSpeech=make(map[string]string)
	re.Response.OutputSpeech["type"]= "PlainText"
    return re
}
func (re *ResponseBody) SetOutputSpeech(output string) {
    re.Response.OutputSpeech["text"] = output
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		re:= NewResponse()

		//TODO store these numbers somewhere else
		people := make(map[string]string)
		people["Anna"]="+17204404892"
		people["Jack"]="+16302355560"

		sender:="+17203304463"

		// unwrap the incoming json data into a LambdaRequest object, Req
		var req LambdaRequest
		if err := json.Unmarshal(event, &req); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		// TODO handle SessionEndedRequest, LaunchRequest
		if req.Request.Type =="IntentRequest" {
			// Extract the intent data from the Request object
			fmt.Fprintln(os.Stderr, req.Request.Intent.Slots)
			message:=req.Request.Intent.Slots["Message"].Value
			recieverName:=req.Request.Intent.Slots["Reciever"].Value
			// look up reciever number, make sure exists
			if _, ok := people[recieverName]; !ok {
				re.SetOutputSpeech("That person is not in you contacts")
				return *re, nil
 			}
			reciever := people[recieverName]
			//TODO access lambda enviornment varibable for twilio api if possible
			var twilioKey ApiKey
			twilioKey.AccountSid = "ACCOUNT_SID"
			twilioKey.AuthToken  = "AUTH_TOKEN"
			//TODO error check
			// Initialize twilio client
			c := twilio.NewClient(twilioKey.AccountSid, twilioKey.AuthToken, nil)
			// Send Message
			params := twilio.MessageParams{
				Body: message,
			}
			// returns message, response, err
			_, _, err := c.Messages.Send(sender, reciever, params)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil, err
			}
			re.SetOutputSpeech("Sent")
			return *re, nil
		}else {
			re.SetOutputSpeech("Sorry that feature has not been implemented")
			return *re, nil
		}
	})
}
