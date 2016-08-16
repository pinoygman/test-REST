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
	//"os"
	"strconv"
	"github.com/pborman/uuid"
	_ "encoding/json"
	_ "github.com/lib/pq"
	//"database/sql"
	//"github.com/jmoiron/sqlx"
	sqltypes "github.com/jmoiron/sqlx/types"
	//"log"
	//"strings"
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
        //AnswerOptions   map[string]interface{}      `json:"answerOptions"`
	AnswerOptions   sqltypes.JSONText           `json:"answerOptions"`
}

var (
	questionnaire map[string]*Question
	questionTypes map[string]string
)

func init(){

	questionnaire=make(map[string]*Question)
	questionTypes=map[string]string{
		strconv.Itoa(Security):"Security",
		strconv.Itoa(Pricing):"Pricing",
		strconv.Itoa(Architecture):"Architecture",
		strconv.Itoa(Service):"Service",
	}
}

func (q *Question) load(guid string) (*Question, error){
	
	return questionnaire[guid],nil
}

func (q *Question) Create() (*Question, error) {

	if q.Guid=="" {
		q.Guid=uuid.New()
	}
	
	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO "pcs-question-tbl" (_id, title, name, description, type, "answerOptions") VALUES ($1, $2, $3, $4, $5, $6)`, q.Guid,q.Title,q.Name,q.Desc,q.Type,q.AnswerOptions)
	
	tx.Commit()

	questionnaire[q.Guid]=q

	return q,nil
}

func (q *Question) Save() (*Question, error) {
	
	tx := db.MustBegin()

	tx.MustExec(`UPDATE "pcs-question-tbl" SET title=$1, name=$2, description=$3, type=$4, "answerOptions"=$5 WHERE _id=$6`,q.Title,q.Name,q.Desc,q.Type,q.AnswerOptions,q.Guid)
	
	tx.Commit()

	return q,nil
}

func DeleteQuestionById(guid string) (error){

	tx := db.MustBegin()

	tx.MustExec(`DELETE FROM "pcs-question-tbl" WHERE _id=$1`,guid)
	
	tx.Commit()

	//delete(questionnaire,q.Guid)
	return nil

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

func GetQuestionsByType(typeId uint64) ([]Question, error){

	fmt.Println(typeId)
	_qs := []Question{}
	db.Select(&_qs, `SELECT _id as "guid", title, name, description as "desc", type,"answerOptions" as "answeroptions" FROM "pcs-question-tbl" WHERE type=$1`,typeId)

	//_ref1:=_qs[0]
	
	fmt.Println(_qs)

	return _qs,nil

}

func GetQuestionTypes() (map[string]string, error) {

        return questionTypes,nil
}

func GetQuestions() ([]Question, error){
	
	_qs := []Question{}
	db.Select(&_qs, `SELECT _id as "guid", title, name, description as "desc", type,"answerOptions" as "answeroptions" FROM "pcs-question-tbl"`)

	//_ref1:=_qs[0]
	
	//fmt.Println(_qs)

	return _qs,nil

}
