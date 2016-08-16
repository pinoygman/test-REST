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
	"net/http"
	"github.build.ge.com/predixsolutions/catalog-onboarding-backend/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
)

func InitServices() error {
	if err:=model.InitRedis();err!=nil {
		return err
	}
	
	if err:=model.InitPostgresSql();err!=nil {
		return err
	}

	return nil

}

func GetApplicationsByProfileIdHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["profileId"]

	_ref,err:=model.GetApplicationsByProfileId(key)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "get applications list error.")
		return
	}

	_str,err:=json.Marshal(_ref)
	
	w.WriteHeader(http.StatusOK)
	w.Write(_str)	
}

func GetDraftsByProfileIdHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["profileId"]

	_ref,err:=model.GetDraftsByProfileId(key)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":"`+err.Error()+`"}`))
		
		fmt.Fprint(w, "get applications list error.")
		return
	}

	_str,err:=json.Marshal(_ref)
	
	w.WriteHeader(http.StatusOK)
	w.Write(_str)	
}

func SubmitApplicationHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	p:=&model.Application{}
	b, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(b, p)

	s,err:=p.Submit()

	if err!=nil {
		w.Write([]byte(`{"err": "`+err.Error()+`"}`))
		return
	}
	
	j, _ := json.Marshal(s)
	w.Write(j)

}

func CreateApplicationHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	p:=&model.Application{}
	b, _ := ioutil.ReadAll(r.Body)

	
	json.Unmarshal(b, p)

	p.Guid=""
	s,err:=p.Save()
	
	if err!=nil{
		w.Write([]byte(`{"err": "`+err.Error()+`"}`))
		return 
	}

	j, _ := json.Marshal(s)
	w.Write(j)
	
}

func SaveApplicationHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	p:=&model.Application{}
	b, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(b, p)

	s,err:=p.Save()
	
	if err!=nil{
		w.Write([]byte(`{"err": "`+err.Error()+`"}`))
		return 
	}

	j, _ := json.Marshal(s)
	w.Write(j)
	
}

func DeleteApplicationHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["applicationId"]

	err :=model.DeleteApplicationById(key)
	
	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":"`+err.Error()+`"}`))
		fmt.Fprint(w, "delete application error.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "application `+key+` has been deleted."}`))
}

func DeleteDraftHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["applicationId"]

	err :=model.DeleteDraftById(key)
	
	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":"`+err.Error()+`"}`))
		fmt.Fprint(w, "delete application error.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "application `+key+` has been deleted."}`))
}
