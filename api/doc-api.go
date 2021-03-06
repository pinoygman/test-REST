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

package api

import (
	"encoding/json"
	//"errors"
	"github.com/pborman/uuid"
	"io"
	//"strings"
	"log"
	"fmt"
	"net/http"
	"time"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3"
	//"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"
	//"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	//"io/ioutil"
	//"strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.build.ge.com/predixsolutions/catalog-onboarding-api/utils"
	"github.build.ge.com/predixsolutions/catalog-onboarding-api/model"

)

const (
	DOCPATH = "./docs/"
	FILEID  = "pcs-fileId"
	LABEL   = "label"
)

var (
	_doc       *S3Api
)

type S3Api struct {
	S           *session.Session
	Svc         *s3.S3
	BucketName  string
}

func InitDocApi(accessKeyID, secretAccessKey, bucketName, endpoint string) {
	
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

	_doc=&S3Api{
		S:          s,
		Svc:        svc,
		BucketName: bucketName,
	}
}

func UploadDocHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	key:=r.Header.Get("ProfileId")
	//vars := mux.Vars(r)
	//key := vars["profileId"]
	
	_a:=&model.Application{ProfileId:key}


	//if strings.ToUpper(r.Header.Get("Content-Type")) == "MULTIPART/FORM-DATA" {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(FILEID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileName := handler.Filename
	//_ct:=handler.Header
	//fmt.Println(_ct["Content-Type"][0])

	uploader := s3manager.NewUploader(_doc.S)
	svc := uploader.S3.(*s3.S3) // in multiPartUpload we don't use the s.svc
	svc.Handlers.Sign.Clear()
	svc.Handlers.Sign.PushBack(utils.SignV2)

	_guid:=uuid.New()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: &_doc.BucketName,
		Key:    &_guid,
	})

	fmt.Println(result)

	if err != nil {
		str:=fmt.Sprintf("failed to upload data to %s with fileName %s", _doc.BucketName, fileName)
		ErrResponse(w,http.StatusInternalServerError,err,str)
		return 
	}

	pd:=&model.Document{
		Guid: _guid,
		Label: r.FormValue(LABEL),
		UploadId: result.UploadID,
		ContentType: handler.Header["Content-Type"][0],
		FileName: fileName,
		CreatedDate: time.Now(),
		CreatedBy: _a.ProfileId,
	}

	_,err2:=pd.Create()

	if err2!=nil {
		str:=fmt.Sprintf("failed to upload data to %s with fileName %s", _doc.BucketName,  fileName)
		ErrResponse(w,http.StatusInternalServerError,err2,str)
		return 
		
	}

	_str,_:=json.Marshal(*pd)

	w.WriteHeader(http.StatusOK)
	w.Write(_str)

}

func DeleteDocHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	fileName := mux.Vars(r)["docId"]

	params := &s3.DeleteObjectInput{
		Bucket: &_doc.BucketName,
		Key:    &fileName,
	}

	_, err := _doc.Svc.DeleteObject(params)
	if err != nil {
		str:=fmt.Sprintf("delete document %s error.",fileName)
		ErrResponse(w,http.StatusInternalServerError,err,str)
		return
	}

	_=model.DeleteDocById(fileName)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Successfully deleted the document# `+ fileName+`"}`))

}

func GetDocListHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	pid:=r.Header.Get("ProfileId")
	//vars := mux.Vars(r)
	//key := vars["profileId"]
	
	_a:=&model.Application{ProfileId:pid}


	_ref,_:=model.GetDocumentsByProfileId(_a.ProfileId)

	_str,_:=json.Marshal(_ref)

	w.WriteHeader(http.StatusOK)
	w.Write(_str)

}

func DownloadDocHttpHandler(w http.ResponseWriter, r *http.Request){
	
	fileName := mux.Vars(r)["docId"]

	input := &s3.GetObjectInput{
		Bucket: &_doc.BucketName,
		Key:    &fileName,
	}

	resp, err := _doc.Svc.GetObject(input)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	_ref:=&model.Document{Guid:fileName,}
	_,err3:=_ref.Load()

	if err3!=nil {
		ErrResponse(w,http.StatusInternalServerError,err,"data loading error.")
		return
	}
	//contentType := http.DetectContentType(resp.Body)
	
	w.Header().Set("Content-Disposition", "attachment; filename="+_ref.FileName)
	w.Header().Set("Content-Type", _ref.ContentType)
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)

}

func ErrResponse(w http.ResponseWriter, errcode uint16, err error, str string) {
	log.Println(str, "Error:", err)		
	fmt.Sprintf("err: %v, reason: %s", err, str)
	w.WriteHeader(int(errcode))
	w.Write([]byte(`{"err":"`+err.Error()+`","reason":"`+str+`"}`))
	//fmt.Fprint(w, )
	return
}


