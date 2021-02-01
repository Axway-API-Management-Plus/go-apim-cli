# ParamValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The parameter name. | [optional] 
**ParamType** | **string** | The type of parameter type.  Can be one of: *body*, *query*, *path*, *form*, or *header*. | [optional] 
**Type** | **string** | The parameter data type.  Can be one of: *string*, *integer*, etc. | [optional] 
**Value** | **string** | The parameter value.  Can be a regular value, or a selector, e.g.: ${params.path.id}. | [optional] 
**Required** | **bool** | Indicates whether or not the parameter is required for the backend API. | [optional] [default to false]
**Exclude** | **bool** | Indicates whether or not the parameter is excluded for the backend API. | [optional] [default to false]
**Additional** | **bool** | Indicates whether or not the parameter is an additional parameter (does not replace an existing parameter). | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


