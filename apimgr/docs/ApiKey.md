# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID which is used to identify the API Key. You include it in each request, so it&#39;s not a secret. | [optional] 
**ApplicationId** | **string** | The ID which is used to identify an application. You include it in each request, so it&#39;s not a secret. | [optional] 
**Secret** | **string** | Each  API Key ID has a Secret Key associated with it. This key is just a long string of characters that can be used to calculate the digital signature that can be included in requests. Your Secret Key is a secret do not distribute. | [optional] 
**Enabled** | **bool** | Flag disables the API key so can&#39;t be used in authentication | [optional] [default to false]
**CreatedBy** | **string** | The unique identifier for user that generated the API Key | [optional] [readonly] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the API key was created | [optional] [readonly] 
**DeletedOn** | **int64** | Epoch/Unix time stamp when the API key was deleted | [optional] [readonly] 
**CorsOrigins** | **[]string** | The domains to allow access for browser-based clients | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


