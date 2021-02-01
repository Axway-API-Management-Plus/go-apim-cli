# ApiAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for approval decisions (includes pending approvals) | [optional] 
**ApiId** | **string** | Virtualised REST API unique id | [optional] 
**CreatedBy** | **string** | The unique identifier for user that requested access | [optional] [readonly] 
**State** | **string** | Pending or approved state | [optional] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the approval decision was created  | [optional] [readonly] 
**Enabled** | **bool** | Flag disables access to an API for organization or application | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


