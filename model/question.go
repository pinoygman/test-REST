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
	"os"
	"strconv"
	"github.com/pborman/uuid"
	_ "encoding/json"
	_ "github.com/lib/pq"
	"gopkg.in/redis.v4"
	"encoding/json"
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
	Status          string                      `json:"status"`
	Type            uint64                      `json:"type"`  //question type
	Priority        uint64                      `json:"priority"`
        //AnswerOptions   map[string]interface{}      `json:"answerOptions"`
	AnswerOptions   sqltypes.JSONText           `json:"answerOptions"`
}

var (
	//questionnaire map[string]*Question
	//questionTypes map[string]string
)

func ResetQuestionType()(map[string]string){
	return map[string]string{
		strconv.Itoa(Security):"Security",
		strconv.Itoa(Pricing):"Pricing",
		strconv.Itoa(Architecture):"Architecture",
		strconv.Itoa(Service):"Service",
	}
}

func InitQuestionType()(map[string]string, error){

	_qt:=make(map[string]string)

	_rq:=os.Getenv("ARTIFACT")+"_questionType"
	
	_b, err := client.Get(_rq).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist in redis")
		//do nothing
		//_ea=make(map[string]*Application)
		_qt=ResetQuestionType()
	} else if err != nil {
		return nil, err
		//panic(fmt.Sprintf("%v", err))
	} else {

		json.Unmarshal([]byte(_b),&_qt)

		if _qt==nil{
			_qt=ResetQuestionType()
		}
	}

	b,_:=json.Marshal(_qt)
	err2 := client.Set(_rq, string(b), 0).Err()

	if err2!=nil {
		return nil, err2
	}

	return _qt, nil

}

func init(){


}
/*
func (q *Question) load(guid string) (*Question, error){
	
	return questionnaire[guid],nil
}
*/
func (q *Question) Create() (*Question, error) {

	if q.Guid=="" {
		q.Guid=uuid.New()
	}
	
	tx := db.MustBegin()

	fmt.Println(q)
	tx.MustExec(`INSERT INTO "pcs-question-tbl" (_id, title, name, description, type, "answerOptions", status, priority) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, q.Guid,q.Title,q.Name,q.Desc,q.Type,q.AnswerOptions, q.Status, q.Priority)
	
	tx.Commit()

	//questionnaire[q.Guid]=q

	return q,nil
}

func (q *Question) Save() (*Question, error) {
	
	tx := db.MustBegin()

	tx.MustExec(`UPDATE "pcs-question-tbl" SET title=$1, name=$2, description=$3, type=$4, "answerOptions"=$5, status=$7, priority=$8 WHERE _id=$6`,q.Title,q.Name,q.Desc,q.Type,q.AnswerOptions,q.Guid, q.Status, q.Priority)
	
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
/*
func InitQuestion(guid string) (*Question, error) {
	op:=&Question{}

	if guid!="" {
		return op.load(guid)
	}

	guid=uuid.New()
	questionnaire[guid]=op

	return questionnaire[guid], nil
}*/

func GetQuestionsByType(typeId uint64) ([]Question, error){

	fmt.Println(typeId)
	_qs := []Question{}
	db.Select(&_qs, `SELECT _id as "guid", title, name, description as "desc", status, type, priority, "answerOptions" as "answeroptions" FROM "pcs-question-tbl" WHERE type=$1`,typeId)

	//_ref1:=_qs[0]
	
	fmt.Println(_qs)

	return _qs,nil

}

func GetQuestions() ([]Question, error){
	
	_qs := []Question{}
	db.Select(&_qs, `SELECT _id as "guid", title, name, description as "desc", status, type, priority, "answerOptions" as "answeroptions" FROM "pcs-question-tbl"`)

	//_ref1:=_qs[0]
	
	//fmt.Println(_qs)
	/*
	for k, v range 	questionnaire {
		_qs=append(_qs,*v)
	}*/
	

	return _qs,nil

}

func GetQuestionTypes() (map[string]string, error) {

	_qt:=make(map[string]string)

	_rq:=os.Getenv("ARTIFACT")+"_questionType"
	
	_b, err := client.Get(_rq).Result()

	if err!=nil{
		return nil,err
	}
	
	json.Unmarshal([]byte(_b),&_qt)
	return _qt,nil
}

func AddQuestionType(key,value string) (map[string]string,error){

	_rq:=os.Getenv("ARTIFACT")+"_questionType"

	_qt,err:=GetQuestionTypes()

	if err!=nil {
		return nil, err
	}

	_qt[key]=value
	
	b,_:=json.Marshal(_qt)

	err2 := client.Set(_rq, string(b), 0).Err()

	if err2!=nil {
		return nil, err2
	}

	return _qt, nil
	
}

func DeleteQuestionTypeById(key string) (error){

	_rq:=os.Getenv("ARTIFACT")+"_questionType"

	_qt,err:=GetQuestionTypes()

	if err!=nil {
		return err
	}

	delete(_qt,key)
	
	b,_:=json.Marshal(_qt)
	err2 := client.Set(_rq, string(b), 0).Err()

	if err2!=nil {
		return err2
	}

	return nil
	
}
