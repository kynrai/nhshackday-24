package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kynrai/nhshackday-24/model"
)

type IanClient struct {
	c *http.Client
}

func NewIanClient() *IanClient {
	return &IanClient{
		c: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (i *IanClient) Read() (*model.Data, error) {
	url := "https://cdr.code4health.org/rest/v1/composition/48badd83-8437-4165-a64b-a613e18c290a::7815d0e1-7902-453c-835a-f2a57c5dfe57::1?format=STRUCTURED&templateId=NHSHACK%20-%20MTX%20report"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic NzgxNWQwZTEtNzkwMi00NTNjLTgzNWEtZjJhNTdjNWRmZTU3OiQyYSQxMCQ2MTlraQ==")
	req.Header.Add("Cookie", "_2269c=http://10.0.0.190:8080")

	res, err := i.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var data model.Data
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
