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
	//"os"
	//"errors"
	//"flag"
	//"log"
	//"path/filepath"
	conf "github.build.ge.com/PredixSolutions/catalog-onboarding-backend/config"
	utils "github.build.ge.com/PredixSolutions/catalog-onboarding-backend/utils"
	"github.com/gorilla/mux"	
)

type ServicePlan struct {
	Name        string      `json:"name"`
	Id          string      `json:"id"`
	Description string      `json:"description"`
	Metadata    interface{} `json:"metadata, omitempty"`
	Free        bool        `json:"free, omitempty"`
}

type Service struct {
	Name           string   `json:"name"`
	Id             string   `json:"id"`
	Description    string   `json:"description"`
	Bindable       bool     `json:"bindable"`
	PlanUpdateable bool     `json:"plan_updateable, omitempty"`
	Tags           []string `json:"tags, omitempty"`
	Requires       []string `json:"requires, omitempty"`

	Metadata        interface{}   `json:"metadata, omitempty"`
	Plans           []ServicePlan `json:"plans"`
	DashboardClient interface{}   `json:"dashboard_client"`
}

type Catalog struct {
	Services []Service `json:"services"`
}

func init(){

}

func main() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(dir)
	defaultConfigPath := utils.GetPath([]string{"assets", "settings.json"})
	fmt.Printf("%+v\n", defaultConfigPath)
	_, err := conf.LoadConfig(defaultConfigPath)
	if err != nil {
		panic(fmt.Sprintf("Error creating server [%s]...", err.Error))
	}

	router := mux.NewRouter()
	router.HandleFunc("/v2/catalog", catalog).Methods("GET")
	
	fmt.Printf("This's catalog onboarding publishing service.\n")
}

func catalog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Service Broker Catalog...")

	var catalog Catalog/*
	catalogFileName := "catalog.json"

	if c.cloudName == utils.AWS {
		catalogFileName = "catalog.AWS.json"
	} else if c.cloudName == utils.SOFTLAYER || c.cloudName == utils.SL {
		catalogFileName = "catalog.SoftLayer.json"
	}

	err := utils.ReadAndUnmarshal(&catalog, conf.CatalogPath, catalogFileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	utils.WriteResponse(w, http.StatusOK, catalog)
}
