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
	//"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	//"github.com/aws/aws-sdk-go/service/s3/s3manager"
	//"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"

	//"log"
	
	//"fmt"
	//"io"
	//"net/http"
)

//question type
const (
//	DueDiligent = 1001
//	Architecture = 1002
//	Security = 1003
)

type Document struct {
	Guid        string                    `json:"_id"`
	Label       string                    `json:"label"`
	s           *session.Session
	svc         *s3.S3
	bucketName  string
	//Title           string                    `json:"title"`
	//Desc            string                    `json:"description"`
	//Type            uint8                     `json:"type"`  //question type
	//AnswerOptions   []string                  `json:"answerOptions"`
	//Answer           map[string]interface{}    `json:"answer"`
	//FileList          []string                  `json:"filesList"`  //file guid
}

var (
	//questionnaires map[string]Questionnaire
)

func InitDoc(accessKeyID, secretAccessKey, bucketName, endpoint string) *Document {
	
	region := "us-east-1"
	
	disableSSL := true
	logLevel := aws.LogDebugWithRequestErrors
	awsConfig := aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Region:      &region,
		Endpoint:    &endpoint,
		DisableSSL:  &disableSSL,
		LogLevel:    &logLevel,
	}

	s := session.New(&awsConfig)

	svc := s3.New(s)

	svc.Handlers.Sign.Clear()
	svc.Handlers.Sign.PushBack(utils.SignV2)

	return &Document{
		Guid: uuid.New(),
		s:          s,
		svc:        svc,
		bucketName: bucketName,
	}

}



