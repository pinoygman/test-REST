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
	"github.com/gorilla/mux"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/rs/cors"
)

var (
	REV string
	conf *config.Config
)

const (
	SETTING = "./settings.json"
	ROOTPATH string = "api"
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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()

	//Questions
	r.HandleFunc(fmt.Sprintf("/%s/%s/question",REV,ROOTPATH), api.UpsertQuestionHttpHandler).Methods("POST")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question",REV,ROOTPATH), api.UpsertQuestionHttpHandler).Methods("PUT")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/list",REV,ROOTPATH), api.GetQuestionsHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/list/{typeId}",REV,ROOTPATH), api.GetQuestionsByTypeIdHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/{key}",REV,ROOTPATH), api.DeleteQuestionHttpHandler).Methods("DELETE")

	//Applications
	r.HandleFunc(fmt.Sprintf("/%s/%s/application",REV,ROOTPATH), api.UpsertApplicationHttpHandler).Methods("POST")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application",REV,ROOTPATH), api.UpsertApplicationHttpHandler).Methods("PUT")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/list/{partnerId}",REV,ROOTPATH), api.GetApplicationsByPartnerIdHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/{applicationId}",REV,ROOTPATH), api.GetApplicationHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/{applicationId}",REV,ROOTPATH), api.DeleteApplicationHttpHandler).Methods("DELETE")
	
	http.Handle(fmt.Sprintf("/%s/",REV), r)

	//cfEnv, err := cfenv.Current()
	_, err := cfenv.Current()

	if err != nil {
		s:=fmt.Sprintf("err cloud foundry env. %s. running server as localhost:%s.", err,conf.Port)
		fmt.Println(s)
		http.ListenAndServe(":"+conf.Port,c.Handler(r))
		return	
	}

	http.ListenAndServe(":"+os.Getenv("PORT"),c.Handler(r))

/*
	for k, _:= range cfEnv.Services {
			if strings.Contains(k,"redis") {
				op.Init(cfEnv.Services[k][0])	
				p, _ :=json.Marshal(op)
				fmt.Println("service.PAEService")
				fmt.Println(string(p))
				log.Fatal("Failed to start server, exiting...", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
				return
			}
		}
 		
	}
*/
}
