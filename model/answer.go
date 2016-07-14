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

)

//question type
const (
//	DueDiligent = 1001
//	Architecture = 1002
//	Security = 1003
)

type Answer struct {
	QuestionId              string                    `json:"_qid"`
	//Title           string                    `json:"title"`
	//Desc            string                    `json:"description"`
	//Type            uint8                     `json:"type"`  //question type
	//AnswerOptions   []string                  `json:"answerOptions"`
	Answer           map[string]interface{}    `json:"answer"`
	FileList          []string                  `json:"filesList"`  //file guid
}

var (
	//questionnaires map[string]Questionnaire
)
