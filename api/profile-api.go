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
)

func GetProfileHttpHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	fmt.Println(model.CurrentProfile)
	_str,_:=json.Marshal(model.CurrentProfile)

	w.Write(_str)
}
