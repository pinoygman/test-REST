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

package main

import (
	"fmt"
	"net/http"
	"os"
	config "github.build.ge.com/predixsolutions/catalog-onboarding-backend/config"
	api "github.build.ge.com/predixsolutions/catalog-onboarding-backend/api"
	model "github.build.ge.com/predixsolutions/catalog-onboarding-backend/model"
	"github.com/gorilla/mux"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/rs/cors"
	"strings"
	//"log"
)

var (
	REV string
	conf *config.Config
)

const (
	SETTING = "./settings.json"
	DOCPATH = "./docs/"
	ROOTPATH string = "api"

	PXBLOB = "predix-blobstore"
	PXREDIS= "redis"
	PXPSQL = "postgres"
	
	BACCESSKEYID     = "access_key_id"
	BSECRETACCESSKEY = "secret_access_key"
	BHOST            = "host"
	BBOCKETNAME      = "bucket_name"

	RHOST = "host"
	RPORT = "port"
	RPWD  = "password"

	SDSN  = "dsn"
		
)

func init(){
	
	s,err:=config.Init(SETTING)

	if err!=nil{
		p:=fmt.Sprintf("error retrieve the %s. %s",SETTING,err)
		fmt.Println(p)
	}

	conf=s

}

func main() {

	if REV=="" {
		REV="v1"
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","POST","DELETE","PUT"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()

	//Profile
	r.HandleFunc(fmt.Sprintf("/%s/%s/profile",REV,ROOTPATH), api.GetProfileHttpHandler).Methods("POST")
	
	//Questions
	r.HandleFunc(fmt.Sprintf("/%s/%s/question",REV,ROOTPATH), api.Auth(api.InsertQuestionHttpHandler)).Methods("POST")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question",REV,ROOTPATH), api.Auth(api.UpdateQuestionHttpHandler)).Methods("PUT")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/list",REV,ROOTPATH), api.Auth(api.GetQuestionsHttpHandler)).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/list/{typeId}",REV,ROOTPATH), api.GetQuestionsByTypeIdHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/{questionId}",REV,ROOTPATH), api.DeleteQuestionHttpHandler).Methods("DELETE")

	//QuestionTypes
	r.HandleFunc(fmt.Sprintf("/%s/%s/questiontype/list",REV,ROOTPATH), api.GetQuestionTypesHttpHandler).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/%s/questiontype",REV,ROOTPATH), api.AddQuestionTypeHttpHandler).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/%s/%s/questiontype/{questionTypeId}",REV,ROOTPATH), api.DeleteQuestionTypeHttpHandler).Methods("DELETE")

	//Applications
	r.HandleFunc(fmt.Sprintf("/%s/%s/application",REV,ROOTPATH), api.Auth(api.CreateApplicationHttpHandler)).Methods("POST")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application",REV,ROOTPATH), api.Auth(api.SaveApplicationHttpHandler)).Methods("PUT")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/submit",REV,ROOTPATH), api.Auth(api.SubmitApplicationHttpHandler)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/%s/%s/application/list",REV,ROOTPATH), api.Auth(api.GetApplicationsByProfileIdHttpHandler)).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/{applicationId}",REV,ROOTPATH), api.Auth(api.DeleteApplicationHttpHandler)).Methods("DELETE")

	//draft
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/draft/list",REV,ROOTPATH), api.Auth(api.GetDraftsByProfileIdHttpHandler)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/%s/application/draft/{applicationId}",REV,ROOTPATH), api.Auth(api.DeleteDraftHttpHandler)).Methods("DELETE")
	
	//email
	r.HandleFunc(fmt.Sprintf("/%s/%s/email",REV,ROOTPATH), api.SendMailApplicationHttpHandler).Methods("POST")

	//document
	r.HandleFunc(fmt.Sprintf("/%s/%s/document",REV,ROOTPATH), api.UploadDocHttpHandler).Methods("POST")	

	r.HandleFunc(fmt.Sprintf("/%s/%s/document/{docId}/delete",REV,ROOTPATH), api.DeleteDocHttpHandler).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/%s/%s/document/{docId}/download",REV,ROOTPATH), api.DownloadDocHttpHandler).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/%s/document/list",REV,ROOTPATH), api.GetDocListHttpHandler).Methods("GET")
	
	//assets
	r.PathPrefix(fmt.Sprintf("/%s/%s/",REV,ROOTPATH)).Handler(http.StripPrefix(fmt.Sprintf("/%s/%s/",REV,ROOTPATH), http.FileServer(http.Dir("./assets"))))

	cfEnv, err := cfenv.Current()

	if err != nil {
		s:=fmt.Sprintf("err cloud foundry env. %s. running server as localhost:%s.", err,conf.Port)
		fmt.Println(s)

		//blobstore
		api.InitDocApi("cf75ae68-2e02-4344-ba37-b85777f176a5-1","4d9a9c03-35c0-4848-af32-42cd4b377bdd","bucket-cf75ae68-2e02-4344-ba37-b85777f176a5", "store.gecis.io")

		//redis
		if err:=model.InitRedisClient("localhost","7991","8f5a2bd2-09db-4b6b-b6e7-2d191b07b11a");err!=nil{
			panic(fmt.Sprintf("%v", err))
		}
		
		//sql
		if err:=model.InitPostgresSql("host=localhost port=7990 user=uc49c9583047d4173a217667509e17ddf password=fb46202694704a7d994dd8e906666e6c dbname=d13291d5f50c645f5b90d26b8a58e2f6b connect_timeout=5 sslmode=disable");err!=nil {
			panic(fmt.Sprintf("%v", err))
		}

		http.ListenAndServe(":"+conf.Port,c.Handler(r))
		
		return
	}

	for k, _:= range cfEnv.Services {
		if strings.Contains(k,PXBLOB) {

			_cred:=cfEnv.Services[k][0].Credentials
			api.InitDocApi(_cred[BACCESSKEYID].(string),
				_cred[BSECRETACCESSKEY].(string),
				_cred[BBOCKETNAME].(string),
				_cred[BHOST].(string),
			)
			
			
		}

		if strings.Contains(k,PXREDIS) {
			o:=cfEnv.Services[k][0].Credentials
			err:=model.InitRedisClient(o[RHOST].(string),fmt.Sprintf("%.0f",o[RPORT]),o[RPWD].(string))
			if err!=nil {
				panic(fmt.Sprintf("%v", err))
			}
		}

		if strings.Contains(k,PXPSQL) {
			o:=cfEnv.Services[k][0].Credentials
			err:=model.InitPostgresSql(o[SDSN].(string))
			if err!=nil {
				panic(fmt.Sprintf("%v", err))
			}
			
		}

	}

	http.ListenAndServe(":"+os.Getenv("PORT"),c.Handler(r))

}
