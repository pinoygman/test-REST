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
  Authorization: Bearer \<UAA Token\>
  PAE-App: Basic \<UAA ClientId:Sec\>
}
```

###Service
```
HTTP 1.1 POST Application/json
Request
{
  "GESponsorship":\<text\>
  "ServiceName":\<text\>
  "Purpose":\<text\>
  "ServiceHowTo":\<text\>
  "IntegrationProsAndCons":\<text\>
  "ServiceVSAltSolutions":\<text\>
  "ServiceCustListBenifitExamples":\<text\>
  "ServicePotentialCustomers":\<text\>
}
Response:
{
  "Status":"Submitted"
  :ApplicationId":"123td-1231f-wef1-312ed"
}
```
