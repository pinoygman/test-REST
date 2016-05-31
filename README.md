# catalog-onboarding-backend
 Catalog onboarding publishing service backend

## Requirement
- Go compiler. GNU Compilers

## Basic Commands:
```
cd ./src/github.build.ge.com/predixsolutions/catalog-onboarding-backend
go install
go build
```

## APIs Info

###Authentication Header
```
{
  Authorization: Bearer <UAA Token>
  CO-App: Basic <UAA ClientId:Sec>
}
```
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

###Service
```
HTTP 1.1 POST Application/json
Request: /v1/api/service
{
  "GESponsorship":<text>,
  "ServiceName":<text>,
  "Purpose":<text>,
  "ServiceHowTo":<text>,
  "IntegrationProsAndCons":<text>,
  "ServiceVSAltSolutions":<text>,
  "ServiceCustListBenifitExamples":<text>,
  "ServicePotentialCustomers":<text>
  "ServiceDocs":["1245e-23423-wfwef2-wrfw2",..]
}

Response:
{
  "Status":"Submitted"
  "ApplicationId":"123td-1231f-wef1-312ed"
}
```

###Architecture
```
HTTP 1.1 POST Application/json
Request: /v1/api/architecture
{
  "ArchUserTypesAndTasks":<text>,
  "ArchIoTUseCases":<text>,
  "ArchAPIDetails":<text>,
  "ArchPredixNative":true/false,
  "ArchExternalSourceList":<text>,
  "ArchLicenseSoftwareList":<text>,
  "ArchMultiTenancyModel":<text>,
  "ArchVersioning":<text>,
  "ArchitectureDocs":["1245e-23423-wfwef2-wrfw2",..]
}

Response:
{
  "Status":"Submitted"
  "ApplicationId":"123td-1231f-wef1-312ed"
}
```

###Pricing:
```
HTTP 1.1 POST Application/json
Request: /v1/api/Pricing
{
  "PriceMetrics":<text>,
  "PriceMeterFeatureList":<text>,
  "PriceSBNuregoReport":<text>,
  "PriceMeterUnit":<text>,
  "PriceUnitCurrentOrCumulative:":<text>,
  "PriceProposal":<text>,
  "PriceNameAndDesc":<text>,
  "PriceInMarket":<text>,
  "PricingDocs":["1245e-23423-wfwef2-wrfw2",..]
}

Response:
{
  "Status":"Submitted"
  "ApplicationId":"123td-1231f-wef1-312ed"
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
    "RequiredRefId:"12345"
    "RequiredRefDesc:"Credit Card Auth Info"
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
    "RequiredRefId:"67893"
    "RequiredRefDesc:"Architecture Diagram"
  ]
}
```

###Architecture Reuirement:
```
HTTP 1.1 POST Application/json
Request: /v1/api/ARequirement/<ApplicationId>
Body:
{
  "ARequiredRefId:<text>,
  "ARequiredField":<text>,
  "ARequiredDoc:[FileId1,FileId2]
} 

Response:
{
  "Status":"Replied",
  "RequiredRefId:"67893",
  "ApplicationId":"123td-1231f-wef1-312ed"
}
```
