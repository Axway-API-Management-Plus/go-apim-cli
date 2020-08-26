#/bin/sh

mkdir -p apimgr
openapi-generator generate -i resources/api-manager-V_1_3-swagger.json -g go --package-name apimgr

### Once the code is generated change the below files to compile the code.
# apimgr/model_schema_object.go:13:6: invalid recursive type SchemaObject 
# 	- make it an array 
#    -->  Items      []SchemaObject       

# apimgr/model_application_request.go
# 	Apis []string `json:"apis,omitempty"` 
#    Delete omitempty --> `json:"apis"`

# apimgr/model_authentication_profile.go 
# 	Parameters map[string]map[string]interface{} `json:"parameters,omitempty"` -->  
# 	Parameters map[string]interface{} `json:"parameters,omitempty"`