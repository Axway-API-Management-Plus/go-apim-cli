# ExternalClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the external client | [optional] [readonly] 
**ClientId** | **string** | The 3rd party client ID. This client ID is provided by a 3rd party OAuth service and is used to map an application to an external client. This value is unique, i.e. no other application can specify the same client ID. | [optional] 
**Enabled** | **bool** | Flag disables the external client so it can&#39;t be used in authentication | [optional] [default to false]
**CreatedBy** | **string** | The unique identifier of the user that created the mapping | [optional] [readonly] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the mapping was created | [optional] [readonly] 
**CorsOrigins** | **[]string** | The domains to allow access for browser-based clients | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


