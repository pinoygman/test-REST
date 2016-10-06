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
	"github.build.ge.com/predixsolutions/catalog-onboarding-api/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"strconv"
)

func GetQuestionTypesHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	_ref,err:=model.GetQuestionTypes()

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "get questions service error.")
		return
	}

	w.WriteHeader(http.StatusOK)

	_str,err:=json.Marshal(&_ref)
	w.Write(_str)
}

func DeleteQuestionTypeHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["questionTypeId"]
	
	err:=model.DeleteQuestionTypeById(key)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		fmt.Fprint(w, "delete question type error.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "question type `+key+` has been deleted."}`))
	return
}

func AddQuestionTypeHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	p:=map[string]string{}
	b, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(b, &p)

	for k, v:= range p {
		_,err:=model.AddQuestionType(k,v)

		if err != nil {
			fmt.Sprintf("err: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"err":`+err.Error()+`}`))
			fmt.Fprint(w, "Add question type error.")
			return
		}
	}	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "question type has been added."}`))
	return
}

func GetQuestionsHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	_ref,err:=model.GetQuestions()

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "get questions service error.")
		return
	}

	w.WriteHeader(http.StatusOK)

	_str,err:=json.Marshal(&_ref)
	w.Write(_str)	
}

func GetQuestionsByTypeIdHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["typeId"]
	
	u, _:=strconv.ParseUint(key,10,32)
	_ref,err:=model.GetQuestionsByType(u)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "get questions service error.")
		return
	}

	w.WriteHeader(http.StatusOK)

	_str,err:=json.Marshal(_ref)
	w.Write(_str)
	return 
}

func UpdateQuestionHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	p:=&model.Question{}
	b, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(b, p)

	s,err:=p.Save()
	
	if err!=nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "update questions service error.")
		return
	}
	
	w.WriteHeader(http.StatusOK)

	j, _ := json.Marshal(&s)
	w.Write(j)
	return 
}

func InsertQuestionHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	p:=&model.Question{}
	b, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(b, p)

	s,err:=p.Create()
	
	if err!=nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		
		fmt.Fprint(w, "update questions service error.")
		return
	}
	
	w.WriteHeader(http.StatusOK)

	j, _ := json.Marshal(&s)
	w.Write(j)
	return 
}

func DeleteQuestionHttpHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["questionId"]

	err:=model.DeleteQuestionById(key)

	if err != nil {
		fmt.Sprintf("err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"err":`+err.Error()+`}`))
		fmt.Fprint(w, "delete service error.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "question `+key+` has been deleted."}`))
	return
}
