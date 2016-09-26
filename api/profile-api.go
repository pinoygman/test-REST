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
	"fmt"
	"os"
	"errors"
)

const (
	TEMPPWD = "TEMPPWD"
)

func GetProfileHttpHandler(w http.ResponseWriter, r *http.Request){

	eml, pwd, ok:=r.BasicAuth()
	if !ok {
		err := errors.New(`Authorization header's not in correct format.`)
		ErrResponse(w,http.StatusInternalServerError,err,"wrong format.")
		return
	}
	//temp pwd verify. will validate against pwd grant_type authentication later
	if pwd==os.Getenv(TEMPPWD){
		pf:=&model.Profile{Email:eml}
		err:=pf.GetProfileByEmail()
		
		if err!=nil {
			err := errors.New(`Invalid profile info.`)
			ErrResponse(w,http.StatusInternalServerError,err,"Invalid field(s) or wrong format.")
			return
		}
		
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		fmt.Println(pf)
		_str,_:=json.Marshal(pf)

		w.Write(_str)

		return
	}

	err:=errors.New("faild Authentication ")

	ErrResponse(w,http.StatusUnauthorized,err,"Authentication failed.")
	return	
}
/*
func GetProfileByPwdHttpHandler(w http.ResponseWriter, r *http.Request){

	
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	fmt.Println(model.CurrentProfile)
	_str,_:=json.Marshal(model.CurrentProfile)

	w.Write(_str)
}
*/
