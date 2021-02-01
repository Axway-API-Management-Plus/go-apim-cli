# SwaggerSecurityDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type identifier for the device. Possible Values: HTTPBasicSecurityDevice, OAuthSecurityDevice, AWSRESTRequestSecurityDevice, AWSQueryStringRequestSecurityDevice, APIKeyOnlySecurityDevice, APIKeyAndSecretSecurityDevice, TwoWaySSLSecurityDevice | [optional] 
**TypeDisplayName** | **string** | Textual display name for the device | [optional] 
**Name** | **string** | Name of the device | [optional] 
**Order** | **int32** | Order of preference, zero being highest. Devices will attempt to authenticate the incoming request using this order of preference. | [optional] 
**Scopes** | **[]string** | The list of scopes defined for the security device. | [optional] 
**ScopeMatching** | **string** | Specifies how scopes will be matched. Possible values: [ Any, All ] | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


