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
	//"github.com/cloudfoundry-community/go-cfenv"
	"fmt"
	"encoding/json"
	"time"
	"github.com/pborman/uuid"
	"gopkg.in/redis.v4"
	"github.com/jmoiron/sqlx"
	sqltypes "github.com/jmoiron/sqlx/types"
	"errors"
	"log"
	//"os"
	//"strings"
)

//question type
const (
	
)

type Application struct {
	Guid              string                  `json:"_id"`
	ProfileId         string                  `json:"profileId"`
	Name              string                  `json:"applicationName"`
	Answers           sqltypes.JSONText       `json:"answers"`
	Notification      sqltypes.JSONText       `json:"notification"`
	CreatedDate       time.Time               `json:"created_date"`
	ModifiedDate      time.Time               `json:"modified_date"`
	ModifiedBy        string                  `json:"modified_by"`
	Status            string                  `json:"applicationStatus"`
}

var (
	client    *redis.Client
	db        *sqlx.DB
)

func init(){
	
	//init redis client
	
}

func (a *Application) Save() (*Application, error) {

	if a.Guid=="" {
		a.Guid=uuid.New()
	}

	_ea:=make(map[string]*Application)
	
	_b, err := client.Get(a.ProfileId).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist in redis")
		//do nothing
		//_ea=make(map[string]*Application)
		
	} else if err != nil {
		return nil, err.(error)
	} else {

		//fmt.Println(_b)
		json.Unmarshal([]byte(_b),&_ea)

		//fmt.Println(_ea)
		if _ea==nil{
			_ea=make(map[string]*Application)
		}
	}

	//fmt.Println("key does not exist in redis")

	fmt.Printf("a: %+v\n",a)
	
	if _, ok := _ea[a.Guid]; ok {
		a.CreatedDate=_ea[a.Guid].CreatedDate
	} else {
		a.CreatedDate=time.Now()
	}
	
	a.ModifiedDate=time.Now()
	a.ModifiedBy=a.ProfileId

	_ea[a.Guid]=a

	fmt.Printf("_ea: %+v\n",_ea[a.Guid])
	
	b,_:=json.Marshal(_ea)
	err2 := client.Set(a.ProfileId, string(b), 0).Err()

	if err2!=nil{
		return nil, err2
	}

	return a, nil

}

func (a *Application) Submit() (*Application,error){
	
	fmt.Println(a)
	
	_ea:=make(map[string]*Application)
	_b, err := client.Get(a.ProfileId).Result()
	if err == redis.Nil {
		return nil, errors.New("draft not found.")
		
	} else if err != nil {
		return nil, err
	} 

	json.Unmarshal([]byte(_b),&_ea)

	if _ea==nil{
		return nil, errors.New("draft not found.")
	}

	if _ea[a.Guid]==nil {
		return nil, errors.New("draft not found.")
	}

	a.Save()
	
	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO "pcs-application-tbl" (_id, profileid, name, answers, status, notification, modifieddate, modifiedby) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,a.Guid,a.ProfileId,a.Name,a.Answers,a.Status,a.Notification,a.ModifiedDate,a.ModifiedBy)

	tx.Commit()

	a.DeleteDraft(a.Guid)
	
	return a, nil
}

func InitRedisClient(_host string,_port string, _pwd string) error {

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s",_host,_port),
		Password: _pwd,
		DB:       0,
	})

	if _, err := client.Ping().Result();err!=nil{
		return err
	}

	InitQuestionType()
	
	return nil
}

func InitPostgresSql(_conn string) error {

	op, err := sqlx.Connect("postgres",_conn)
	
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	db=op
	return nil

}

func (a *Application) GetApplications() ([]Application,error){
	_ap := []Application{}
	db.Select(&_ap, `SELECT _id as "guid", profileid, name, status, answers, notification,createddate, modifieddate, modifiedby FROM "pcs-application-tbl" where profileid=$1`,a.ProfileId)
	//created_date, last_modifed,
	return _ap,nil
}

func (a *Application) GetDrafts() ([]Application ,error){

	_ea:=make(map[string]*Application)
	_b, err := client.Get(a.ProfileId).Result()
	if err == redis.Nil {
		return nil, errors.New("draft not found 1.")
		
	} else if err != nil {
		return nil, err
	} 

	json.Unmarshal([]byte(_b),&_ea)

	if _ea==nil{
		return nil, errors.New("draft not found 2")
	}

	var _a  []Application
	for _, v := range _ea {
		_a=append(_a,*v)
	}
	return _a,nil
	
}

func (a *Application) DeleteDraft(guid string) (error){

	_ea:=make(map[string]*Application)
	_b, err := client.Get(a.ProfileId).Result()
	if err == redis.Nil {
		return errors.New("draft not found 1.")
		
	} else if err != nil {
		return err
	} 

	json.Unmarshal([]byte(_b),&_ea)

	delete(_ea,guid)
	
	_v,_:=json.Marshal(_ea)

	err2 := client.Set(a.ProfileId, string(_v), 0).Err()

	if err2!=nil{
		return err2
	}

	return nil
}


func (a *Application) DeleteApplication(guid string) (error){

	if err:=a.DeleteDraft(guid);err!=nil{
		return err
	}

	tx := db.MustBegin()

	tx.MustExec(`DELETE FROM "pcs-application-tbl" WHERE _id=$1 and profileid=$2`,guid,a.ProfileId)
	
	tx.Commit()

	return nil
}

