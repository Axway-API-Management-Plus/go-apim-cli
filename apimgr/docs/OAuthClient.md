# OAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The client ID to be used in OAuth flows | [optional] 
**Cert** | **string** | The PEM encodeded certificate used in JWT flow | [optional] 
**Secret** | **string** | The client application secret to be used in OAuth flows | [optional] 
**Type** | **string** | OAuth defines two client types, based on their ability to authenticate securely with the authorization server. Possible values public or confidential | [optional] 
**Enabled** | **bool** | Flag disables the OAuth credentials so they can&#39;t be used in authentication | [optional] [default to false]
**RedirectUrls** | **[]string** | The URL where the server will redirect the to present authorization codes or access tokens depending on the OAuth flow being executed | [optional] 
**CorsOrigins** | **[]string** | The domains to allow access for browser-based clients | [optional] 
**CreatedBy** | **string** | The unique identifier for user that generated the OAuth credentials | [optional] [readonly] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the OAuth credentials was created | [optional] [readonly] 
**ApplicationId** | **string** | The application identifier associated with the OAuth credential | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


