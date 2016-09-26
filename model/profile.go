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
	"fmt"
	//"log"

	//"github.com/nimajalali/go-force/force"
	//"github.com/nimajalali/go-force/sobjects"
)

//question type
const (
//	DueDiligent = 1001
//	Architecture = 1002
//	Security = 1003
)

type Profile struct {
	ProfileId         string                    `json:"_pid"`
	Name              string                    `json:"name"`
	Email             string                    `json:"email"` 
	SFDCId            string                    `json:"sfdcid"`
}

var (
	//questionnaires map[string]Questionnaire
	//CurrentProfile    * Profile 
)

func init(){
}

func (p *Profile) Init(profileId string) (error){
	
	db.Get(p, `SELECT _id, name, email, sfdcid FROM "pcs-profile" where _id=$1`,profileId)
	fmt.Printf("%#v\n", p)
	
	return  nil
}

func GetProfileByEmail(email string) (Profile, error) {

	_pf:=Profile{}

	db.Get(&_pf, `SELECT _id, name, email, sfdcid FROM "pcs-profile" where email=$1`,email)
	fmt.Printf("%#v\n", _pf)
	
	return _pf, nil	
}
