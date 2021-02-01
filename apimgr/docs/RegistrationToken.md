# RegistrationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The registration code | [optional] 
**OrganizationId** | **string** | Unique identifier for the organization who the registration code applies to | [optional] [readonly] 
**Expiry** | **int64** | Epoch/Unix time stamp when the registration code will expire | [optional] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the registration code was created | [optional] [readonly] 
**CreatedBy** | **string** | The unique identifier for user that created the registration code | [optional] [readonly] 
**UserQuota** | **int32** | The remaining number of users that can use the registration code for self registration to an organization | [optional] 
**MaxUsers** | **int32** | The total number of users that can use the registration code for self registration to an organization since the code has been created | [optional] 
**Enabled** | **bool** | Flag disables registration code so that it can no longer be used for registration | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


