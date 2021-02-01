# \OrganizationsApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsGet**](OrganizationsApi.md#OrganizationsGet) | **Get** /organizations | List all organizations
[**OrganizationsIdApisApiAccessIdDelete**](OrganizationsApi.md#OrganizationsIdApisApiAccessIdDelete) | **Delete** /organizations/{id}/apis/{apiAccessId} | Deletes access to an API for an organization
[**OrganizationsIdApisApiAccessIdPut**](OrganizationsApi.md#OrganizationsIdApisApiAccessIdPut) | **Put** /organizations/{id}/apis/{apiAccessId} | Updates access to an API for an organization
[**OrganizationsIdApisGet**](OrganizationsApi.md#OrganizationsIdApisGet) | **Get** /organizations/{id}/apis | Get the list of approved APIs for the organization
[**OrganizationsIdApisPost**](OrganizationsApi.md#OrganizationsIdApisPost) | **Post** /organizations/{id}/apis | Grants access to an API for an organization.
[**OrganizationsIdDelete**](OrganizationsApi.md#OrganizationsIdDelete) | **Delete** /organizations/{id} | Delete an organization
[**OrganizationsIdGet**](OrganizationsApi.md#OrganizationsIdGet) | **Get** /organizations/{id} | Get an organization
[**OrganizationsIdImageGet**](OrganizationsApi.md#OrganizationsIdImageGet) | **Get** /organizations/{id}/image | Get the image for an organization
[**OrganizationsIdImagePost**](OrganizationsApi.md#OrganizationsIdImagePost) | **Post** /organizations/{id}/image | Set the image for an organization
[**OrganizationsIdPut**](OrganizationsApi.md#OrganizationsIdPut) | **Put** /organizations/{id} | Update the details of an organization
[**OrganizationsIdTokensGet**](OrganizationsApi.md#OrganizationsIdTokensGet) | **Get** /organizations/{id}/tokens | Get registration codes for an organization
[**OrganizationsIdTokensPost**](OrganizationsApi.md#OrganizationsIdTokensPost) | **Post** /organizations/{id}/tokens | Create a registration code
[**OrganizationsIdTokensTokenDelete**](OrganizationsApi.md#OrganizationsIdTokensTokenDelete) | **Delete** /organizations/{id}/tokens/{token} | Delete the registration code
[**OrganizationsIdTokensTokenGet**](OrganizationsApi.md#OrganizationsIdTokensTokenGet) | **Get** /organizations/{id}/tokens/{token} | Get registration code
[**OrganizationsIdTokensTokenPut**](OrganizationsApi.md#OrganizationsIdTokensTokenPut) | **Put** /organizations/{id}/tokens/{token} | Update a registration code
[**OrganizationsPost**](OrganizationsApi.md#OrganizationsPost) | **Post** /organizations | Creates a new organization



## OrganizationsGet

> []Organization OrganizationsGet(ctx, optional)

List all organizations

Get the list of organizations that are visible to the authenticated user.  Only API Administrators may list all organizations, all other users will see their organization.  The list of organizations can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  apiid : Matches the organization if the organization is using the API, specified by ID  description : The organization's description  email : The organization's contact email address  enabled :  The enabled state of the organization, one of: enabled, disabled  createdOn : The date the organization was created on, time in ms, e.g.: 1372755998542  name : The name of the organization  phone : The organization's contact phone  The __op__ is an operation and is one of:  eq : Equal  ne : Not equal  gt :  Greater than  lt :  Less than  ge :  Greater than or equal  le :  Less than or equal  like : Like  gete :  Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

### Return type

[**[]Organization**](Organization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdApisApiAccessIdDelete

> ApiAccess OrganizationsIdApisApiAccessIdDelete(ctx, id, apiAccessId)

Deletes access to an API for an organization

Permanently deletes access to an API for an organization.  Deleting API access will also delete API access to any application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID. | 
**apiAccessId** | **string**| The API access ID. | 

### Return type

[**ApiAccess**](APIAccess.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdApisApiAccessIdPut

> ApiAccess OrganizationsIdApisApiAccessIdPut(ctx, id, apiAccessId, optional)

Updates access to an API for an organization

Updates access to an API for an organization.  Only __enabled__ may be modified, and disabling access will also disable access to all applications that may be using it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID. | 
**apiAccessId** | **string**| The API access ID. | 
 **optional** | ***OrganizationsIdApisApiAccessIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsIdApisApiAccessIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiAccess**](ApiAccess.md)|  | 

### Return type

[**ApiAccess**](APIAccess.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdApisGet

> ApiAccess OrganizationsIdApisGet(ctx, id)

Get the list of approved APIs for the organization

Get the list of aproved APIs for the organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID who&#39;s approved APIs are required. | 

### Return type

[**ApiAccess**](APIAccess.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdApisPost

> ApiAccess OrganizationsIdApisPost(ctx, id, optional)

Grants access to an API for an organization.

Grants access to an API for an organization.  Only the API Admin may call this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID requesting access to an API. | 
 **optional** | ***OrganizationsIdApisPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsIdApisPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiAccess**](ApiAccess.md)|  | 

### Return type

[**ApiAccess**](APIAccess.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdDelete

> OrganizationsIdDelete(ctx, id)

Delete an organization

Deletes an organization. Deleting an organization will result in all users and associated applications being deleted

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID to delete. | 

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


## OrganizationsIdGet

> Organization OrganizationsIdGet(ctx, id)

Get an organization

Retrieves the details of an organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID to be returned. | 

### Return type

[**Organization**](Organization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdImageGet

> OrganizationsIdImageGet(ctx, id)

Get the image for an organization

Returns the jpeg image associated with an organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID whos image is to be returned | 

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


## OrganizationsIdImagePost

> OrganizationsIdImagePost(ctx, id, optional)

Set the image for an organization

Set the jpeg image to be associated with an organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID whos image is to be set | 
 **optional** | ***OrganizationsIdImagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsIdImagePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdPut

> Organization OrganizationsIdPut(ctx, id, optional)

Update the details of an organization

Updates an organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID to update. | 
 **optional** | ***OrganizationsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Organization**](Organization.md)|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdTokensGet

> []RegistrationToken OrganizationsIdTokensGet(ctx, id)

Get registration codes for an organization

Retrieves the registration codes for an organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID whos tokens are to be returned. | 

### Return type

[**[]RegistrationToken**](RegistrationToken.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdTokensPost

> RegistrationToken OrganizationsIdTokensPost(ctx, id, optional)

Create a registration code

Create a registration code for self service onboarding of users to the organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID to be associated with the registration code. | 
 **optional** | ***OrganizationsIdTokensPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsIdTokensPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RegistrationToken**](RegistrationToken.md)|  | 

### Return type

[**RegistrationToken**](RegistrationToken.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdTokensTokenDelete

> OrganizationsIdTokensTokenDelete(ctx, id, token)

Delete the registration code

Delete the registration code.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID associated with the registration code. | 
**token** | **string**| The registration code to be deleted. | 

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


## OrganizationsIdTokensTokenGet

> []RegistrationToken OrganizationsIdTokensTokenGet(ctx, id, token)

Get registration code

Retrieves the registration code.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID associated with the registration code. | 
**token** | **string**| The registration code to be returned. | 

### Return type

[**[]RegistrationToken**](RegistrationToken.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsIdTokensTokenPut

> RegistrationToken OrganizationsIdTokensTokenPut(ctx, id, token, optional)

Update a registration code

Update a registration code for self service onboarding of users to the organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The organization ID to be associated with the registration code. | 
**token** | **string**|  | 
 **optional** | ***OrganizationsIdTokensTokenPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsIdTokensTokenPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RegistrationToken**](RegistrationToken.md)|  | 

### Return type

[**RegistrationToken**](RegistrationToken.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsPost

> Organization OrganizationsPost(ctx, optional)

Creates a new organization

Creates a new organization.  Only API Administrators may create organizations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrganizationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Organization**](Organization.md)|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

