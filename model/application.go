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
)

//question type
const (
	
)

type Application struct {
	Guid            string              `json:"_id"`
	//Title           string              `json:"title"`
	//Desc            string              `json:"description"`
	//Type            uint8               `json:"type"`  //question type
	Answers         map[string]Answer   `json:"answers"`
}

var (
	//questionnaires map[string]Questionnaire
	applications map[string]*Application
)

func init(){
	
	applications=make(map[string]*Application)

}

func InitApplication(guid string) (*Application, error){

	op:=&Application{}

	if guid!="" {
		return op.load(guid)
	}

	guid=uuid.New()
	
	applications[guid]=op

	return applications[guid], nil

}

func GetApplicationsByPartnerId(pId string) (map[string]*Application,error){

	//op:=make(map[string]*Application)
	
	return applications, nil

}

func (a *Application) load(guid string) (*Application, error){

	return applications[guid],nil
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
	if a.Guid=="" {
		a.Guid=uuid.New()
	}

	applications[a.Guid]=a
	return applications[a.Guid],nil
}

func (a *Application) Submit() (*Application,error){
	applications[a.Guid]=a
	return applications[a.Guid],nil
}

func (a *Application) Del() (string,error){

	delete(applications,a.Guid)
	return "delete",nil
}
