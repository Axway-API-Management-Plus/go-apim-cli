# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An identifier | [optional] [readonly] 
**Title** | **string** | Schema title | [optional] 
**Description** | **string** | Description of the Schema | [optional] 
**Format** | **string** | The format ex: int32, int64, float, double, byte, binary, date, date-time or password | [optional] 
**Required** | **[]string** | Specifies if the parameter is required | [optional] 
**Properties** | [**map[string]SchemaObject**](SchemaObject.md) | Not used beacause our model does not support inline nested types | [optional] 
**Items** | [**SchemaObject**](SchemaObject.md) |  | [optional] 
**Example** | [**map[string]interface{}**](.md) | if the schema is an array specifies the items type | [optional] 
**Discriminator** | **string** |  | [optional] 
**HttpMethod** | **string** | The HTTP method | [optional] 
**Nickname** | **string** | The nickname of the operation | [optional] 
**Summary** | **string** | A short summary description of the operation | [optional] 
**Notes** | **string** | A detailed description of the operation | [optional] 
**ResponseClass** | **string** | The return type of the method, e.g. void, array, or a type found in models | [optional] 
**ErrorResponses** | [**[]ErrorResponse**](ErrorResponse.md) | A list of possible response messages and their meanings | [optional] 
**Consumes** | **[]string** | The content types that the operation consumes | [optional] 
**Produces** | **[]string** | The content types that the operation produces | [optional] 
**Authorizations** | [**map[string][]map[string]interface{}**](array.md) | Authorizations | [optional] 
**Tags** | [**map[string][]string**](array.md) | The list of tags associated with this API operation. Each tag can have multiple values | [optional] 
**SecurityProfile** | [**SwaggerSecurityProfile**](SwaggerSecurityProfile.md) |  | [optional] 
**DocumentationUrl** | **string** | The documentation URL for the operation | [optional] 
**Cors** | **bool** | Indicates that the API is CORS enabled | [optional] [default to false]
**Parameters** | [**[]Parameter**](Parameter.md) | A list of accepted parameters | [optional] 
**Ref** | **string** | A Reference to a definition on definitions object | [optional] 
**Default** | [**map[string]interface{}**](.md) | Default value for this schema if it is applicable | [optional] 
**Type** | **string** | The return type of the method, e.g. void, array, or a type found in models | [optional] 
**Enum** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


