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
	//"log"
	"net/smtp"
)

var (
	REV string
	conf *config.Config
)

const (
	SETTING = "./settings.json"
	DOCPATH = "./docs/"
	ROOTPATH string = "api"
)

func send(){

	fmt.Println("..send email.")
	// Set up authentication information.
	auth := smtp.PlainAuth("", "raasuser", "helloraas", "localhost")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	//to := []string{"chia.chang@ge.com","Subba.Vadrevu@ge.com","subrata.saha@ge.com","Javier.Carbajal.Ramirez@ge.com"}
	to := []string{"chia.chang@ge.com"}
	msg := []byte("To: PCS-Portal-Development@ge.com; \r\n" +
		"Subject: Chia test SMTP for PCS Onboarding Portal #3\r\n" +
		"\r\n" +
		"This is a test email sent from the Predix Select. #3\r\n")
	//err := smtp.SendMail("rssmtp-212359746.run.aws-usw02-pr.ice.predix.io:80", auth, "chia.chang@ge.com", to, msg)
	err := smtp.SendMail("localhost:31373", auth, "chia.chang@ge.com", to, msg)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}

}

func init(){
	
	s,err:=config.Init(SETTING)

	if err!=nil{
		p:=fmt.Sprintf("error retrieve the %s. %s",SETTING,err)
		fmt.Println(p)
	}

	conf=s
	
	_=api.InitServices()

}

func main() {

	send()
	
	if REV=="" {
		REV="v1"
	}
/*
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})
*/

	
	r := mux.NewRouter()

	//Profile
	r.HandleFunc(fmt.Sprintf("/%s/%s/profile",REV,ROOTPATH), api.GetProfileHttpHandler).Methods("POST")
	
	//Questions
	r.HandleFunc(fmt.Sprintf("/%s/%s/question",REV,ROOTPATH), api.InsertQuestionHttpHandler).Methods("POST")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question",REV,ROOTPATH), api.UpdateQuestionHttpHandler).Methods("PUT")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/list",REV,ROOTPATH), api.GetQuestionsHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/list/{typeId}",REV,ROOTPATH), api.GetQuestionsByTypeIdHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/question/{questionId}",REV,ROOTPATH), api.DeleteQuestionHttpHandler).Methods("DELETE")

	//QuestionTypes
	r.HandleFunc(fmt.Sprintf("/%s/%s/questiontype/list",REV,ROOTPATH), api.GetQuestionTypesHttpHandler).Methods("GET")
	
	//Applications
	r.HandleFunc(fmt.Sprintf("/%s/%s/application",REV,ROOTPATH), api.CreateApplicationHttpHandler).Methods("POST")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application",REV,ROOTPATH), api.SaveApplicationHttpHandler).Methods("PUT")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/submit",REV,ROOTPATH), api.SubmitApplicationHttpHandler).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/%s/%s/application/list/{profileId}",REV,ROOTPATH), api.GetApplicationsByProfileIdHttpHandler).Methods("GET")
	
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/{applicationId}",REV,ROOTPATH), api.DeleteApplicationHttpHandler).Methods("DELETE")

	//draft
	r.HandleFunc(fmt.Sprintf("/%s/%s/application/draft/list/{profileId}",REV,ROOTPATH), api.GetDraftsByProfileIdHttpHandler).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/%s/application/draft/{applicationId}",REV,ROOTPATH), api.DeleteDraftHttpHandler).Methods("DELETE")
	
	//assets
	r.PathPrefix(fmt.Sprintf("/%s/%s/",REV,ROOTPATH)).Handler(http.StripPrefix(fmt.Sprintf("/%s/%s/",REV,ROOTPATH), http.FileServer(http.Dir("./assets"))))
	
	//http.Handle(fmt.Sprintf("/%s/%s/",REV,ROOTPATH), r)
	
	//cfEnv, err := cfenv.Current()
	_, err := cfenv.Current()

	if err != nil {
		s:=fmt.Sprintf("err cloud foundry env. %s. running server as localhost:%s.", err,conf.Port)
		fmt.Println(s)
		http.ListenAndServe(":"+conf.Port,cors.Default().Handler(r))
		
		return	
	}

	http.ListenAndServe(":"+os.Getenv("PORT"),cors.Default().Handler(r))

	
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
