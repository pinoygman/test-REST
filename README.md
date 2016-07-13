# catalog-onboarding-backend
 Catalog onboarding publishing service backend

## Requirement
- Go compiler. GNU Compilers
- Makefile

## Basic Commands:

### Build
```
cd ./<root-repo>
make
```

##Cloud Foundry: (Binary Push)
```
cd ./<root-repo>
make deploy
```

## APIs Info

###Documents
```
HTTP 1.1 POST multipart/form-data
Request: /v1/api/file

File/FileName

Response:
{
  "FileId":"1245e-23423-wfwef2-wrfw2"
}
```

###Application
```
HTTP 1.1 POST multipart/form-data
Request: /v1/api/application
{
  _id: <guid>
	partnerId: <string>
	name: <string>
	answers: {
    		"questionId1":{
       			_qid: <string>
       			content: <any_object>
       			filesList: ["file_guid1","file_guid2",...n]
    		}
    		"questionId2":{
       			_qid: <string>
       			content: <any_object>
       			filesList: ["file_guid1","file_guid2",...n]
    		}
	}
	appstatus: <string>
}
```

###Question
```
HTTP 1.1 POST multipart/form-data
Request: /v1/api/question
{
   _id:<guid>
   title:<string>
   description:<string>
   type: <DueDiligent = 1001, Architecture = 1002, Security = 1003>
   answerOptions: {any_object}
}
```

###Orchestration
```
{
  questionId: <guid>,
  triggerValue: <(text,num,boolean)>,
  trueQuestionId: <guid>,
  falseQuestionId: <guid>
}
```

###QuestionLog
```
{
   questionIds: <array>,
   currentQuestionId: <guid>,
   applicationId: <guid>
}
```

###Authentication Header
```
{
  Authorization: Bearer <UAA Token>
  CO-App: Basic <UAA ClientId:Sec>
}
```

###Partner:
```
HTTP 1.1 GET
Request: /v1/api/Partner/<PartnerId, SFDCContactId>

Response: Application/json
{
  "PartnerInfo":
  {
    "ContactName":<text>,
    "ContactEmail":<text>,
    "ContactPhone":<text>,
    "ContactOrgName":<text>
  }
}
```

###Due Diligent Status:
```
HTTP 1.1 GET Application/json
Request: /v1/api/DDStatus/<AppcationId>

Response:
{
  "DDCurrentStep":<integer>,
  "DDRequiredInfo":[
    {
      "RequiredRefId:"12345",
      "RequiredRefDesc:"Credit Card Auth Info"
    }
  ]
}
```

###Due Diligent Requirement:
```
HTTP 1.1 POST Application/json
Request: /v1/api/DDRequirement/<ApplicationId>
{
  "DDRequiredRefId":<text>,
  "DDRequiredField":<text>,
  "DDRequiredDoc":[FileId1,FileId2]
}

Response:
{
  "Status":"Replied",
  "RequiredRefId:"67893",
  "ApplicationId":"123td-1231f-wef1-312ed"
}
```

###Architecture Status:
```
HTTP 1.1 GET Application/json
Request: /v1/api/AStatus/<AppcationId>

Response:
{
  "ACurrentStep"<text>,
  "ARequiredInfo":[
    {
      "RequiredRefId:"67893",
      "RequiredRefDesc:"Architecture Diagram"
    }
  ]
}
```

###Architecture Reuirement:
```
HTTP 1.1 POST Application/json
Request: /v1/api/ARequirement/<ApplicationId>
Body:
{
  "ARequiredQuestionId:<text>,
  "ARequiredField":<text>,
  "ARequiredDoc:[FileId1,FileId2]
} 

Response:
{
  "Status":"Replied",
  "RequiredQuestionId:"67893",
  "ApplicationId":"123td-1231f-wef1-312ed"
}
```
