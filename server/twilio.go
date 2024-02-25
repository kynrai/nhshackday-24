package server

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Twilio struct {
	from   string
	client *twilio.RestClient
	to     string
}

func NewTwilio(sid, token, from, to string) *Twilio {
	return &Twilio{
		from: from,
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: sid,
			Password: token,
		}),
		to: to,
	}
}

func (t *Twilio) SendSMS(to, from, body string) (*twilioApi.ApiV2010Message, error) {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(t.from)
	params.SetBody(body)

	return t.client.Api.CreateMessage(params)
}
