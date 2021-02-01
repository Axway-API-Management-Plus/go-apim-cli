# ErrorResponse

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
**Code** | **int64** | The response code | [optional] 
**Reason** | **string** | The response reason | [optional] 
**Ref** | **string** | A Reference to a definition on definitions object | [optional] 
**Default** | [**map[string]interface{}**](.md) | Default value for this schema if it is applicable | [optional] 
**Type** | **string** | The type ex: array , boolean, integer , null , number, object, string | [optional] 
**Enum** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


