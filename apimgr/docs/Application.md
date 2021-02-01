# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the application | [optional] [readonly] 
**Name** | **string** | The display name for the application | 
**Description** | **string** | Descriptive text for the application | [optional] 
**OrganizationId** | **string** | The organization identifier to which the application belongs | 
**Phone** | **string** | Contact phone number of the application | [optional] 
**Email** | **string** | The contact email address associated with the application | [optional] 
**CreatedBy** | **string** | The unique identifier for user that created the application | [optional] [readonly] 
**ManagedBy** | **[]string** | A list of unique identifier for users that manages the application | [optional] 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the application was created | [optional] [readonly] 
**Enabled** | **bool** | Flag to indicate if this application is enabled or not | [optional] [default to false]
**Image** | **string** | URI of the image to be used for this application, this field only indicates that the application has an image assigned to it. In order to retrieve the actual image use the following URL /api/portal/applications/{id}/image/ | [optional] 
**State** | **string** | Flag to indicate if an application has been approved by the API Manager admin or if delegated then the org admin | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


