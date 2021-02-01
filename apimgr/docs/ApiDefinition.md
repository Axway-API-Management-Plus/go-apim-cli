# ApiDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. | [optional] 
**Name** | **string** | The name of the API. | [optional] 
**Summary** | **string** | A summary of the API. | [optional] 
**Description** | **string** | A detailed description of the API. | [optional] 
**Version** | **string** | The API version. | [optional] 
**BasePath** | **string** | The base path is where the API service is hosted. | [optional] 
**ResourcePath** | **string** | The resource path is applied to **basePath** to provide the prefix for all API methods. | [optional] 
**Models** | **map[string]map[string]interface{}** | The models/schema the that the API | [optional] 
**Consumes** | **[]string** | The content types that the API consumes | [optional] 
**Produces** | **[]string** | The content types that the API produces | [optional] 
**Integral** | **bool** | Indicates that the API definition is integral to a frontend API; that the API was imported to define the frontend API. | [optional] [default to false]
**CreatedOn** | **int64** | Epoch/Unix time stamp when the organization was created. | [optional] 
**CreatedBy** | **string** | The identifier of the user that created the API. | [optional] 
**OrganizationId** | **string** | The [Organization](Organization.html) identifier. | [optional] 
**ServiceType** | **string** | Indicates the type of service being imported. Possible values: rest, wsdl. | [optional] 
**HasOriginalDefinition** | **bool** | Indicates whether or not an original definition is available | [optional] [default to false]
**ImportUrl** | **string** | Specifies the URL used to import the backend API definition. | [optional] 
**Properties** | **map[string]string** | A list of properties associated with this API. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


