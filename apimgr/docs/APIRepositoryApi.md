# \APIRepositoryApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApirepoGet**](APIRepositoryApi.md#ApirepoGet) | **Get** /apirepo | Get the list of API
[**ApirepoIdDelete**](APIRepositoryApi.md#ApirepoIdDelete) | **Delete** /apirepo/{id} | Deletes an API.
[**ApirepoIdDownloadGet**](APIRepositoryApi.md#ApirepoIdDownloadGet) | **Get** /apirepo/{id}/download | Downloads an API by ID.
[**ApirepoIdGet**](APIRepositoryApi.md#ApirepoIdGet) | **Get** /apirepo/{id} | Get an API by ID
[**ApirepoIdMethodsGet**](APIRepositoryApi.md#ApirepoIdMethodsGet) | **Get** /apirepo/{id}/methods | Queries the list of API methods
[**ApirepoIdMethodsMethodIdDelete**](APIRepositoryApi.md#ApirepoIdMethodsMethodIdDelete) | **Delete** /apirepo/{id}/methods/{methodId} | Delete an API method
[**ApirepoIdMethodsMethodIdGet**](APIRepositoryApi.md#ApirepoIdMethodsMethodIdGet) | **Get** /apirepo/{id}/methods/{methodId} | Get API method by ID.
[**ApirepoIdMethodsMethodIdPut**](APIRepositoryApi.md#ApirepoIdMethodsMethodIdPut) | **Put** /apirepo/{id}/methods/{methodId} | Update an API method
[**ApirepoIdMethodsPost**](APIRepositoryApi.md#ApirepoIdMethodsPost) | **Post** /apirepo/{id}/methods | Create an API method
[**ApirepoIdPut**](APIRepositoryApi.md#ApirepoIdPut) | **Put** /apirepo/{id} | Updates an API
[**ApirepoImportFromExternalPost**](APIRepositoryApi.md#ApirepoImportFromExternalPost) | **Post** /apirepo/importFromExternal | Create one or more backend APIs for an external service
[**ApirepoImportFromGwPost**](APIRepositoryApi.md#ApirepoImportFromGwPost) | **Post** /apirepo/importFromGw | Create an API definition by importing a PolicyStudio-registered web service (REST or WSDL) hosted on the the API Gateway
[**ApirepoImportFromUrlPost**](APIRepositoryApi.md#ApirepoImportFromUrlPost) | **Post** /apirepo/importFromUrl | Create an API by loading a file from URL.
[**ApirepoImportPost**](APIRepositoryApi.md#ApirepoImportPost) | **Post** /apirepo/import | Create an API by uploading a file
[**ApirepoPost**](APIRepositoryApi.md#ApirepoPost) | **Post** /apirepo | Create an API definition



## ApirepoGet

> []ApiDefinition ApirepoGet(ctx, optional)

Get the list of API

Get the list of API from the API repository.  The list of API can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  id :    Matches the API by ID  name :    Matches the API by name  The __op__ is an operation and is one of:  eq :    Equal  ne :    Not equal  gt :     Greater than  lt :     Less than  ge :     Greater than or equal  le :     Less than or equal  like :    Like  gete :     Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApirepoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApirepoGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | **optional.String**|  | 
 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

### Return type

[**[]ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdDelete

> ApirepoIdDelete(ctx, id)

Deletes an API.

Deletes a backend API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifier. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdDownloadGet

> ApirepoIdDownloadGet(ctx, id, filename, original)

Downloads an API by ID.

Downloads an API by ID.  If __filename__ is not supplied, the API name will be used.  If the API was imported using [/import](#importApisFromFile) or [/import](#createApiFromUrl), then it is possible to download the original API definition by setting __original__ to __true__.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**filename** | **string**| Override the default filename for download | 
**original** | **bool**| If true, and the API was imported, this will download the original definition | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdGet

> ApiDefinition ApirepoIdGet(ctx, id)

Get an API by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdMethodsGet

> []Method ApirepoIdMethodsGet(ctx, id, optional)

Queries the list of API methods

Get the list of API methods from the API repository.  The list of methods can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  id :    Matches the API by ID  name :    Matches the API by name  The __op__ is an operation and is one of:  eq :    Equal  ne :    Not equal  gt :     Greater than  lt :     Less than  ge :     Greater than or equal  le :     Less than or equal  like :    Like  gete :     Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***ApirepoIdMethodsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApirepoIdMethodsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

### Return type

[**[]Method**](Method.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdMethodsMethodIdDelete

> ApirepoIdMethodsMethodIdDelete(ctx, id, methodId)

Delete an API method

Deletes a backend API method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifier. | 
**methodId** | **string**| The method identifier. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdMethodsMethodIdGet

> Method ApirepoIdMethodsMethodIdGet(ctx, id, methodId)

Get API method by ID.

Retrieves a method for a given API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifier. | 
**methodId** | **string**| The API method ID. | 

### Return type

[**Method**](Method.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdMethodsMethodIdPut

> Method ApirepoIdMethodsMethodIdPut(ctx, id, methodId, method)

Update an API method

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifierentifier to create a method. | 
**methodId** | **string**| The method identifier. | 
**method** | [**Method**](Method.md)| The method to update. | 

### Return type

[**Method**](Method.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdMethodsPost

> Method ApirepoIdMethodsPost(ctx, id, method)

Create an API method

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifierentifier to create a method. | 
**method** | [**Method**](Method.md)| The method to create. | 

### Return type

[**Method**](Method.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoIdPut

> ApiDefinition ApirepoIdPut(ctx, id, optional)

Updates an API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifier. | 
 **optional** | ***ApirepoIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApirepoIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiDefinition**](ApiDefinition.md)|  | 

### Return type

[**ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoImportFromExternalPost

> []ApiDefinition ApirepoImportFromExternalPost(ctx, organizationId, connectorId, name, description, api)

Create one or more backend APIs for an external service

Create one or more backend APIs for an external service. External APIs are imported via a connector. If the connector configuration specifies that all external APIs should be merged into a single new backend API, the name and description parameters are applied to this new API. Alternatively, if the connector specifies that a separate backend API should be created for each external API, the name and description parameters are ignored, and the names and descriptions of the new backend APIs are taken from the external service definitions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| The API development organization that owns the new APIs | 
**connectorId** | **string**| The API connector through which new APIs should be created | 
**name** | **string**| The name of the merged API (see description) | 
**description** | **string**| A description of the merged API (see description) | 
**api** | **string**| List of external APIs to be imported | 

### Return type

[**[]ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoImportFromGwPost

> ApiDefinition ApirepoImportFromGwPost(ctx, id, name, organizationId, optional)

Create an API definition by importing a PolicyStudio-registered web service (REST or WSDL) hosted on the the API Gateway

Imports an API definition from a Policy Studio REST or WSDL service hosted on the API Gateway. On import, a Swagger representation of the original API definition is retained, but the API is converted to an internal format for processing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the PolicyStudio-registered service to import. | 
**name** | **string**| The service name. | 
**organizationId** | **string**| The API development organization ID that owns the import. | 
 **optional** | ***ApirepoImportFromGwPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApirepoImportFromGwPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **instance** | **optional.String**|  | 
 **host** | **optional.String**|  | 
 **port** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 

### Return type

[**ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoImportFromUrlPost

> ApiDefinition ApirepoImportFromUrlPost(ctx, organizationId, type_, url, optional)

Create an API by loading a file from URL.

Imports an API definition from a valid standard Swagger or WADL definition from the specified __url__.  It is possible to supply an optional __username__ and __password__ if the __url__ requires HTTP Basic authentication.  On import, the original API definition is retained, but the API is converted to an internal format for processing. The API name currently defaults to the filename but this will be deprecated in a future release. The name parameter should be used to name the API and will be required in a future release.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| The API development organization ID that owns the import. | 
**type_** | **string**| The type of import, one of: swagger, wadl, raml. | 
**url** | **string**| The URL to import. | 
 **optional** | ***ApirepoImportFromUrlPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApirepoImportFromUrlPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **optional.String**| The name of the API. | 
 **fileName** | **optional.String**| The file name of the import. | 
 **username** | **optional.String**| HTTP Basic username to use for connection. | 
 **password** | **optional.String**| HTTP Basic password to use for connection. | 

### Return type

[**ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoImportPost

> ApiDefinition ApirepoImportPost(ctx, organizationId, name, type_, file)

Create an API by uploading a file

Imports an API definition from a valid standard Swagger or WADL definition.  On import, the original API definition is retained, but the API is converted to an internal format for processing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| The API development organization ID that owns the import. | 
**name** | **string**| The API name. | 
**type_** | **string**| The type of import, one of: swagger, wadl, raml | 
**file** | ***os.File*****os.File**| The API definition file to import | 

### Return type

[**ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApirepoPost

> ApiDefinition ApirepoPost(ctx, api)

Create an API definition

When creating an API, the __name__ and __basePath__ are required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | [**ApiDefinition**](ApiDefinition.md)| The API resource to create. | 

### Return type

[**ApiDefinition**](APIDefinition.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

