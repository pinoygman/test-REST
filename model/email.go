/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: chia.chang@ge.com
 */

package model

import (
	"fmt"
	"net/smtp"
	"strings"
	"encoding/base64"
	"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"
)

var (
	auth     smtp.Auth
	haraka   string
)

const (

)

type Email struct {
	From             string                    `json:"from"` 
	To               string                    `json:"to"`
        Cc               string                    `json:"cc"`
	Bcc              string                    `json:"bcc"`
        Subject          string                    `json:"subject"`
        Body             string                    `json:"body"` 
        Attachment       []string                  `json:"attachment"`
}

func InitEmail(identity, username, password, hrk, host string) (*Email){
	auth = smtp.PlainAuth(
		identity,
		username,
		password,
		host,
	)
	haraka=hrk
	return &Email{}
}

func (e *Email) Send() (*Email, error){

	header := make(map[string]string)
	header["From"] = e.From
	header["To"] = e.To
	//header["Cc"] = e.Cc
	header["Subject"] = utils.EncodeRFC2047(e.Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(e.Body))

	_tos:=[]string{}

	for _, elm := range strings.Split(e.To,";") {
		_tos=append(_tos,elm)
	}

	for _, elm := range strings.Split(e.Cc,";") {
		_tos=append(_tos,elm)
	}

	for _, elm := range strings.Split(e.Bcc,";") {
		_tos=append(_tos,elm)
	}

	err := smtp.SendMail(
		haraka,
		auth,
		e.From,
		_tos,
		[]byte(message),
	)
	
	if err != nil {
		return nil, err
	}

	return e, nil
}
