# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The parameter name. | [optional] 
**Type** | **string** | The parameter data type, e.g. *boolean*, *byte*, *date*, *double*, *float*, *integer*, *long*, *string*, or a type name found in [APIDefinition models](APIDefinition.html#models). | [optional] 
**Format** | **string** |  | [optional] 
**Description** | **string** |  | [optional] 
**Required** | **bool** | Indicates that the parameter is required | [optional] [default to false]
**AllowMultiple** | **bool** | Indicates that the parameter can be included multiple times (e.g. query or form) | [optional] [default to false]
**Items** | [**SchemaObject**](SchemaObject.md) |  | [optional] 
**DefaultValue** | **string** | Provides a default value for the parameter | [optional] 
**Schema** | [**SchemaObject**](SchemaObject.md) |  | [optional] 
**ParamType** | **string** | The parameter type, e.g. query, form, path, body, header | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


