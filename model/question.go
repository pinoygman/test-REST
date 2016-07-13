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
	/*DueDiligent = 1001
	Architecture = 1002
	Security = 1003
*/
)
type QuestionType struct {

	Id    uint64   `json:"_id"`
	Desc  string   `json:"desc"`
}

type Question struct {
	Guid            string                      `json:"_id"`
	Title           string                      `json:"title"`
	Desc            string                      `json:"description"`
	Type            uint64                       `json:"type"`  //question type
	AnswerOptions   map[string]interface{}      `json:"answerOptions"`
}

var (
	questionnaire map[string]*Question
	questionTypes map[uint64]string
)

func init(){
	questionnaire=make(map[string]*Question)
	questionTypes=map[uint64]string{
		1001:"DueDiligent",
		1002:"Architecture",
		1003:"Security",
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

func (q *Question) load(guid string) (*Question, error){
	
	return questionnaire[guid],nil
}

func (q *Question) Save() (*Question, error) {

	if q.Guid=="" {
		q.Guid=uuid.New()
	}

	questionnaire[q.Guid]=q
	return q,nil
}


func (q *Question) Del() (string,error){

	delete(questionnaire,q.Guid)
	return "delete",nil

}
