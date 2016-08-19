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
	"io"
	"strings"
	"log"
	"fmt"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"
	"github.build.ge.com/predixsolutions/catalog-onboarding-backend/model"
	"github.com/aws/aws-sdk-go/service/s3"
	//"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"
	//"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	//"io/ioutil"
	//"strconv"
)

const (
	DOCPATH = "./docs/"
	FILEID  = "pcs-fileId"
)

var (
	_doc       *model.Document
)

func InitDocApi(accessKeyId, secretAccessKey, bucketName, endpoint string){
	
	_doc=model.InitDoc(accessKeyId, secretAccessKey, bucketName, endpoint)

}

func UploadDocHttpHandler(w http.ResponseWriter, r *http.Request){

	if strings.ToUpper(r.Header.Get("Content-Type")) == "MULTIPART/FORM-DATA" {

		fmt.Println("start multipart")
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile(FILEID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileName := handler.Filename

		uploader := s3manager.NewUploader(_doc.S)
		svc := uploader.S3.(*s3.S3) // in multiPartUpload we don't use the s.svc
		svc.Handlers.Sign.Clear()
		svc.Handlers.Sign.PushBack(utils.SignV2)

		result, err := uploader.Upload(&s3manager.UploadInput{
			Body:   file,
			Bucket: &_doc.BucketName,
			Key:    &fileName,
		})

		if err != nil {
			log.Println("Failed to upload.", "Error:", err)
		}

		log.Println("Successfully uploaded data.", "FileName:", fileName, "uploadID:", result.UploadID)
		fmt.Println("end multi end")
		http.Redirect(w, r, "/", http.StatusFound)
		return

	}
	
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(FILEID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := handler.Filename

	_, err = _doc.Svc.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: &_doc.BucketName,
		Key:    &fileName,
	})

	if err != nil {
		log.Println("failed to upload data to", "", _doc.BucketName, "/", fileName, "Error:", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)

}

func DeleteDocHttpHandler(w http.ResponseWriter, r *http.Request){
	fileName := mux.Vars(r)["docId"]

	params := &s3.DeleteObjectInput{
		Bucket: &_doc.BucketName,
		Key:    &fileName,
	}

	_, err := _doc.Svc.DeleteObject(params)
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
	/*
	w.Header().Set("Content-Type", "application/json")

	uid, err:=utils.RetrieveUpdateFile(r,DOCPATH); 

	
	if err != nil {
		fmt.Sprintln("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok","_docId":"`+uid+`"}`))
*/
}

func GetDocListHttpHandler(w http.ResponseWriter, r *http.Request){
	params := &s3.ListObjectsInput{
		Bucket: &_doc.BucketName, // Required
	}
	resp, err := _doc.Svc.ListObjects(params)
	if err != nil {
		log.Println(err)
	}
	var files []string
	for _, file := range resp.Contents {
		files = append(files, *file.Key)
	}
	fmt.Println("End Get All Objects")
	http.Redirect(w, r, "/", http.StatusOK)

	/*
	w.Header().Set("Content-Type", "application/json")

	uid, err:=utils.RetrieveUpdateFile(r,DOCPATH); 

	
	if err != nil {
		fmt.Sprintln("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok","_docId":"`+uid+`"}`))
*/
}

func UpdateDocHttpHandler(w http.ResponseWriter, r *http.Request){


}

func DownloadDocHttpHandler(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Begin Get Object")
	fileName := mux.Vars(r)["docId"]
	fmt.Println("Get Object", "docId", fileName)

	input := &s3.GetObjectInput{
		Bucket: &_doc.BucketName,
		Key:    &fileName,
	}

	resp, err := _doc.Svc.GetObject(input)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", *resp.ContentType)
	io.Copy(w, resp.Body)
	fmt.Println("End Get Object")

	/*
	w.Header().Set("Content-Type", "application/json")

	uid, err:=utils.RetrieveUpdateFile(r,DOCPATH); 

	
	if err != nil {
		fmt.Sprintln("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok","_docId":"`+uid+`"}`))
*/
}


