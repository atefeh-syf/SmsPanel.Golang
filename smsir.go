package smsir

import "github.com/atefeh-syf/SmsPanel.Golang/IPE.SmsIrClient/services"



func VerifySend(key, mobile, temp string, parameters []services.VerifySendParameter) {
	verify := &services.VerifySendRequest{
		ApiKey: key,
		Mobile: mobile,
		TemplateId: temp,
		Parameters: parameters,
	}

	resp, err := verify.Send()
	if err != nil{
		// do something
	}
}