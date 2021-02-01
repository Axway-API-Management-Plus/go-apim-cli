# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the organization | [optional] [readonly] 
**Name** | **string** | The name of the organization | 
**Description** | **string** | The description of the organization | [optional] 
**Email** | **string** | The contact email address associated with the organization | [optional] 
**Image** | **string** | URI of the image to be used for this organization. To update the image, please refer to the API. | [optional] 
**Restricted** | **bool** | Indicates that the organization is restricted.  Users in a restricted organization cannot see other users, and users cannot register for the organization using tokens.  Default is &#39;false&#39;. | [optional] [default to false]
**VirtualHost** | **string** | The virtual host associated with the organization | [optional] 
**Phone** | **string** | Contact phone number of the organization | [optional] 
**Enabled** | **bool** | Flag to indicate if this organization is enabled or not | [optional] [default to false]
**Development** | **bool** | Flag to indicate if this organization is enabled or not for API development. | [optional] [default to false]
**Dn** | **string** |  | [optional] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the organization was created | [optional] [readonly] 
**StartTrialDate** | **int64** | Epoch/Unix time stamp when the trial starts | [optional] 
**EndTrialDate** | **int64** | Epoch/Unix time stamp when the trial expires | [optional] 
**TrialDuration** | **int32** | Length of the trial in days | [optional] 
**IsTrial** | **bool** | Indicates if this Org is a trial or not | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


