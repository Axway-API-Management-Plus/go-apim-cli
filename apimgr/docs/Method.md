# Method

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The method identifier | [optional] 
**ApiId** | **string** | The API identifier to which this method belongs | [optional] 
**Path** | **string** | The API path | [optional] 
**Verb** | **string** | The HTTP verb | [optional] 
**Name** | **string** | The name of the operation | [optional] 
**Summary** | **string** | A short summary description of the operation | [optional] 
**Description** | **string** | A detailed description of the operation | [optional] 
**ReturnType** | **string** | The return type of the method, e.g. void, array, or a type found in models | [optional] 
**Parameters** | [**[]Parameter**](Parameter.md) | A list of accepted parameters | [optional] 
**ResponseCodes** | [**[]ResponseCode**](ResponseCode.md) | A list of possible response messages and their meanings | [optional] 
**Consumes** | **[]string** | The content types that the operation consumes | [optional] 
**Produces** | **[]string** | The content types that the operation produces | [optional] 
**Properties** | **map[string]string** | A list of properties associated with this API Method. The list of properties may vary, depending on the type of the parent API. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


