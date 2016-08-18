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
	//"github.com/gorilla/mux"
	//"fmt"
	"io/ioutil"
	//"strconv"
	//"github.com/cloudfoundry-community/go-cfenv"
	"os"
)

func SendMailApplicationHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	//auth := smtp.PlainAuth("", "raasuser", "helloraas", "localhost")

	p:=model.InitEmail("",os.Getenv("HUSER"),os.Getenv("HPWD"),os.Getenv("HHOST"),os.Getenv("HSHOST"))
	
	b, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(b, p)

	if _, err:=p.Send();err!=nil {
		w.Write([]byte(`{"err": "`+err.Error()+`"}`))
		return
	}
	
	w.Write([]byte(`{"status": "email has been sent successfully."}`))

}
