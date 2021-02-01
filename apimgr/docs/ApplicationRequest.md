# ApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the application request | [optional] 
**Name** | **string** | The display name for the application | [optional] 
**Description** | **string** | Descriptive text for the application | [optional] 
**OrganizationId** | **string** | The organization identifier to which the application request belongs | [optional] 
**Phone** | **string** | Contact phone number of the application | [optional] 
**Email** | **string** | The contact email address associated with the application | [optional] 
**Image** | **string** | URI of the image to be used for this application, this field only indicates that the application has an image assigned to it. In order to retrieve the actual image use the following URL /api/portal/organizations/{uid of org}/image/ | [optional] 
**Apis** | **[]string** | A list of unqiue API identifiers to which the application wants to use. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


