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

package model

import (	
	"github.com/pborman/uuid"
	"errors"
)

//question type
const (
	
)

type Application struct {
	Guid            string              `json:"_id"`
	PartnerId       string              `json:"partnerId"`
	//Title           string              `json:"title"`
	//Desc            string              `json:"description"`
	//Type            uint8               `json:"type"`  //question type
	Answers         map[string]Answer   `json:"answers"`
}

var (
	//questionnaires map[string]Questionnaire
	applications map[string]interface{}
)

func init(){
	
	applications=make(map[string]interface{})

}

func GetApplicationsByPartnerId(pId string) (map[string]*Application,error){

	//op:=make(map[string]*Application)
	
	return applications[pId].(map[string]*Application), nil

}

func GetApplication(guid string) (*Application, error){

	for _, l := range applications {
		v:=l.(map[string]*Application)
		if a:=v[guid];a!=nil {
			return a,nil
		}
	}

	return nil, errors.New("application does not exist.")
	
}


func (a *Application) Save() (*Application, error) {

	/*
	if services==nil {
		sevrvices=make(map[string]PAEService)
		fmt.Println("service list initialised.")
	}

	if k=="" {
		p.Guid=uuid.New()
	} else {
		p.Guid=k
	}
	
	services[p.Guid]=*p

	b, _ := json.Marshal(services[p.Guid]);
	
	fmt.Println(string(b))
	
	err := client.Set(p.Guid, string(b), 0).Err()

	//validate
	e:=&PAEService{}
	val, _ := client.Get(p.Guid).Result()
	json.Unmarshal([]byte(val),e)
	
	//fmt.Println("err adding: ",err)
	return e, err
*/

	if a.PartnerId=="" {
		return nil, errors.New("partnerId's empty.")
	}


	var ap interface{}

	if ap=applications[a.PartnerId]; ap==nil {
		ap=make(map[string]*Application)
		applications[a.PartnerId]=ap
	}
	
	if a.Guid=="" {
		a.Guid=uuid.New()
	}	

	ref:=ap.(map[string]*Application)
	ref[a.Guid]=a
	
	return ref[a.Guid],nil
	
}

func (a *Application) Submit() (*Application,error){
	a.Save()
	return applications[a.PartnerId].(map[string]*Application)[a.Guid],nil
}

func (a *Application) Del() (string,error){

	delete(applications[a.PartnerId].(map[string]*Application),a.Guid)
	return "delete",nil
}
