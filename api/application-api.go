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

func GetApplicationsByPartnerIdHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["partnerId"]

	_ref,err:=model.GetApplicationsByPartnerId(key)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "get applications service error.")
		return
	}

	_str,err:=json.Marshal(_ref)
	
	w.WriteHeader(http.StatusOK)
	w.Write(_str)	
}

func GetApplicationHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["applicationId"]

	_ref, err:=model.InitApplication(key)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "get application service error.")
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

func UpsertApplicationHttpHandler(w http.ResponseWriter, r *http.Request) {

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

	_ref,_:=model.InitApplication(key)

	_,err:=_ref.Del()
	
	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		fmt.Fprint(w, "delete application error.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "application `+key+` has been deleted."`))
}
