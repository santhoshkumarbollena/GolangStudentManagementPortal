I set the client Id and Client Secret to "demo" and "demo"


To get Access Key:
http://localhost:9096/token?grant_type=client_credentials&client_id=demo&client_secret=demo&scope=all


The response contains Access Key Copy That key and paste it in below url

http://localhost:9096/FirstService?access_token=ACCESS_KEY_PASTE_HERE

request method:POST

body:
 {
    "requestId" :                   "stringrequestId", 
	"memberId":                     "string memberId",
	"memberIdType"    :              "string memberIdType",
	"referedToSpecialtyCategory":   "string referedToSpecialtyCategory",
	"referedToSpecialityCode" :      ["string referedToSpecialityCode","string2","string3"],
	"referedToSpecialityAreaOfBody": "string referedToSpecialityAreaOfBody",
	"providerIds"   :                ["string providerIds","string2","string3"],
	"searchFilterCriteria"   :       "string searchFilterCriteria",
	"callingApp"         :       "string callingApp",
	"callingAppType"     :           "string callingAppType"
   }

   we can create another service to generate random client id and client Secret

If we want to generate random client Id and client Secret the Refer SecureRestApiExample where API "/credentials"
Need to be copy and paste

The response contains client Id and Client Secret
copy and paste in URI used to get Access_key in respective positions