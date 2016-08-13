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
	"github.build.ge.com/predixsolutions/catalog-onboarding-backend/sql"
)

//question type
const (
	
	Security = 1004
	Pricing = 1003
	Architecture = 1002
	Service = 1001

)


type Question struct {
	Guid            string                      `json:"_id"`
	Title           string                      `json:"title"`
	Name            string                      `json:"name"`
	Desc            string                      `json:"description"`
	Type            uint64                      `json:"type"`  //question type
	AnswerOptions   map[string]interface{}      `json:"answerOptions"`
}

var (
	questionnaire map[string]*Question
	questionTypes map[uint64]string
)


func (q *Question) load(guid string) (*Question, error){
	
	return questionnaire[guid],nil
}

func (q *Question) Save() (*Question, error) {

	if q.Guid=="" {
		q.Guid=uuid.New()
	}

	err:=sql.AddAQuestion(q.Guid,q.Title,q.Name,q.Desc,q.Type,q.AnswerOptions)
	if err!=nil{
		return nil,err 
	}
	
	questionnaire[q.Guid]=q
	return q,nil
}

func (q *Question) Del() (string,error){

	delete(questionnaire,q.Guid)
	return "delete",nil

}


func init(){
	questionnaire=make(map[string]*Question)
	questionTypes=map[uint64]string{
		Security:"Security",
		Pricing:"Pricing",
		Architecture:"Architecture",
		Service:"Service",
	}
}

func InitQuestion(guid string) (*Question, error) {
	op:=&Question{}

	if guid!="" {
		return op.load(guid)
	}

	guid=uuid.New()
	questionnaire[guid]=op

	return questionnaire[guid], nil
}

func GetQuestionsByType(typeId uint64) (map[string]*Question, error){

	//op:=make(map[string]*Question)
	return questionnaire,nil

}

func GetQuestionTypes()(map[uint64]string, error){
        return questionTypes,nil
}

func GetQuestions() (map[string]*Question, error){

	//op:=make(map[string]*Question)
	return questionnaire,nil

}
