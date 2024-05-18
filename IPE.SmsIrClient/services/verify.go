package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)


type  VerifySendRequest struct {
	ApiKey     string
	Mobile     string
	TemplateId string
	Parameters []VerifySendParameter
}

type VerifySendParameter struct {
	Name  string
	Value string
}

var url string = "https://api.sms.ir/v1/send/verify"

func(v  *VerifySendRequest) Send() (*http.Response, error){
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", v.ApiKey)

	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return res, nil
}