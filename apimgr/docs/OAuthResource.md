# OAuthResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the oauth protected resource | 
**ApplicationId** | **string** | The unique identifier for the application who has access to this resource | 
**Uriprefix** | **string** | The uri prefix which this oauth protected resource relates to | [optional] 
**Scopes** | **[]string** | Set of scope strings that have been resolved from querying the OAuth Resource Service at the uri prefix | [optional] 
**Enabled** | **bool** | Flag to indicate if this oauth protected resource is enabled or not | [optional] [default to false]
**IsDefault** | **bool** |  | [optional] [default to false]
**Scope** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


