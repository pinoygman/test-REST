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
	//"fmt"
	"net/http"
	//"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"fmt"
	//"io/ioutil"
	//"strconv"
)

const (
	DOCPATH = "./docs/"	
)

func UploadDocesHttpHandler(w http.ResponseWriter, r *http.Request){
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
