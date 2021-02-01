# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the user | [optional] 
**OrganizationId** | **string** | The unique identifier for the organization to which the user belongs | [optional] 
**Name** | **string** | The user&#39;s name | [optional] 
**Description** | **string** | A description of the user | [optional] 
**LoginName** | **string** | A unique login name for the user | [optional] 
**Email** | **string** | An email address for the user | [optional] 
**Phone** | **string** | The user&#39;s phone number | [optional] 
**Mobile** | **string** | The user&#39;s mobile number | [optional] 
**Role** | **string** | The user&#39;s role, one of: user, oadmin, or admin | [optional] 
**Image** | **string** | The user&#39;s photo. To update the image, please refer to the API. | [optional] 
**Enabled** | **bool** | Indicates whether or not the user account is enabled or not | [optional] [default to false]
**CreatedOn** | **int64** | Epoch/Unix time stamp when the organization was created | [optional] 
**State** | **string** | The current state of the account, one of: approved, pending | [optional] 
**Type** | **string** | Indicates the type of user. Possible values: internal, external | [optional] 
**AuthAttrs** | [**AuthenticatedUserAttributes**](AuthenticatedUserAttributes.md) |  | [optional] 
**Dn** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


