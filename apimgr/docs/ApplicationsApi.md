# \ApplicationsApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsGet**](ApplicationsApi.md#ApplicationsGet) | **Get** /applications | Get the list of applications
[**ApplicationsIdApikeysApikeyidPut**](ApplicationsApi.md#ApplicationsIdApikeysApikeyidPut) | **Put** /applications/{id}/apikeys/{apikeyid} | Updates an API Key
[**ApplicationsIdApikeysGet**](ApplicationsApi.md#ApplicationsIdApikeysGet) | **Get** /applications/{id}/apikeys | Returns the API Keys associated with an application
[**ApplicationsIdApikeysKeyIdDelete**](ApplicationsApi.md#ApplicationsIdApikeysKeyIdDelete) | **Delete** /applications/{id}/apikeys/{keyId} | Delete an API Key
[**ApplicationsIdApikeysPost**](ApplicationsApi.md#ApplicationsIdApikeysPost) | **Post** /applications/{id}/apikeys | Creates a new API Key and secret for the application
[**ApplicationsIdApisApiAccessIdApprovePost**](ApplicationsApi.md#ApplicationsIdApisApiAccessIdApprovePost) | **Post** /applications/{id}/apis/{apiAccessId}/approve | Creates an API access request to an API for an application.
[**ApplicationsIdApisApiAccessIdDelete**](ApplicationsApi.md#ApplicationsIdApisApiAccessIdDelete) | **Delete** /applications/{id}/apis/{apiAccessId} | Deletes access to an API for an application
[**ApplicationsIdApisApiAccessIdPut**](ApplicationsApi.md#ApplicationsIdApisApiAccessIdPut) | **Put** /applications/{id}/apis/{apiAccessId} | Updates access to an API for an application
[**ApplicationsIdApisGet**](ApplicationsApi.md#ApplicationsIdApisGet) | **Get** /applications/{id}/apis | Get the list of APIs that the application can access
[**ApplicationsIdApisPost**](ApplicationsApi.md#ApplicationsIdApisPost) | **Post** /applications/{id}/apis | Create a request for an application to access an API.
[**ApplicationsIdApprovePost**](ApplicationsApi.md#ApplicationsIdApprovePost) | **Post** /applications/{id}/approve | Approves a pending application request
[**ApplicationsIdAvailablescopesGet**](ApplicationsApi.md#ApplicationsIdAvailablescopesGet) | **Get** /applications/{id}/availablescopes | Returns the scopes available to an application
[**ApplicationsIdDelete**](ApplicationsApi.md#ApplicationsIdDelete) | **Delete** /applications/{id} | Delete an application
[**ApplicationsIdExtclientsGet**](ApplicationsApi.md#ApplicationsIdExtclientsGet) | **Get** /applications/{id}/extclients | Returns the external clients associated with an application
[**ApplicationsIdExtclientsObjectIdDelete**](ApplicationsApi.md#ApplicationsIdExtclientsObjectIdDelete) | **Delete** /applications/{id}/extclients/{objectId} | Delete an external client
[**ApplicationsIdExtclientsObjectIdPut**](ApplicationsApi.md#ApplicationsIdExtclientsObjectIdPut) | **Put** /applications/{id}/extclients/{objectId} | Updates an external client for the application
[**ApplicationsIdExtclientsPost**](ApplicationsApi.md#ApplicationsIdExtclientsPost) | **Post** /applications/{id}/extclients | Maps a new external client to the application
[**ApplicationsIdGet**](ApplicationsApi.md#ApplicationsIdGet) | **Get** /applications/{id} | Get an application
[**ApplicationsIdImageGet**](ApplicationsApi.md#ApplicationsIdImageGet) | **Get** /applications/{id}/image | Get the image for an application
[**ApplicationsIdImagePost**](ApplicationsApi.md#ApplicationsIdImagePost) | **Post** /applications/{id}/image | Adds a JPEG image to an application
[**ApplicationsIdOauthClientIdPut**](ApplicationsApi.md#ApplicationsIdOauthClientIdPut) | **Put** /applications/{id}/oauth/{clientId} | Updates an OAuth Credential for the application
[**ApplicationsIdOauthClientidNewsecretPut**](ApplicationsApi.md#ApplicationsIdOauthClientidNewsecretPut) | **Put** /applications/{id}/oauth/{clientid}/newsecret | Updates an OAuth Credential for an application by generating a new secret
[**ApplicationsIdOauthGet**](ApplicationsApi.md#ApplicationsIdOauthGet) | **Get** /applications/{id}/oauth | Returns the OAuth Credentials associated with an application
[**ApplicationsIdOauthOauthIdDelete**](ApplicationsApi.md#ApplicationsIdOauthOauthIdDelete) | **Delete** /applications/{id}/oauth/{oauthId} | Delete an OAuth client ID and secret
[**ApplicationsIdOauthPost**](ApplicationsApi.md#ApplicationsIdOauthPost) | **Post** /applications/{id}/oauth | Creates a new OAuth client ID and secret for the application
[**ApplicationsIdOauthresourceGet**](ApplicationsApi.md#ApplicationsIdOauthresourceGet) | **Get** /applications/{id}/oauthresource | Returns the OAuth protected resources (scopes) associated with an application
[**ApplicationsIdOauthresourcePost**](ApplicationsApi.md#ApplicationsIdOauthresourcePost) | **Post** /applications/{id}/oauthresource | Adds an OAuth protected resource to an application
[**ApplicationsIdOauthresourceResourceIdDelete**](ApplicationsApi.md#ApplicationsIdOauthresourceResourceIdDelete) | **Delete** /applications/{id}/oauthresource/{resourceId} | Remove an OAuth protected resource from an application
[**ApplicationsIdOauthresourceResourceIdPut**](ApplicationsApi.md#ApplicationsIdOauthresourceResourceIdPut) | **Put** /applications/{id}/oauthresource/{resourceId} | Updates a protected resource associate with an application, sets enabled to true/false
[**ApplicationsIdPermissionsGet**](ApplicationsApi.md#ApplicationsIdPermissionsGet) | **Get** /applications/{id}/permissions | Get the list of permissions.
[**ApplicationsIdPermissionsPermIdDelete**](ApplicationsApi.md#ApplicationsIdPermissionsPermIdDelete) | **Delete** /applications/{id}/permissions/{permId} | Remove a permission
[**ApplicationsIdPermissionsPermIdPut**](ApplicationsApi.md#ApplicationsIdPermissionsPermIdPut) | **Put** /applications/{id}/permissions/{permId} | Modify a permission
[**ApplicationsIdPermissionsPost**](ApplicationsApi.md#ApplicationsIdPermissionsPost) | **Post** /applications/{id}/permissions | Create a new permission.
[**ApplicationsIdPut**](ApplicationsApi.md#ApplicationsIdPut) | **Put** /applications/{id} | Update an application
[**ApplicationsIdQuotaDelete**](ApplicationsApi.md#ApplicationsIdQuotaDelete) | **Delete** /applications/{id}/quota | Deletes a quota from an application
[**ApplicationsIdQuotaGet**](ApplicationsApi.md#ApplicationsIdQuotaGet) | **Get** /applications/{id}/quota | Returns the quota associated with an application.
[**ApplicationsIdQuotaPost**](ApplicationsApi.md#ApplicationsIdQuotaPost) | **Post** /applications/{id}/quota | Creates a new quota constraint for the application
[**ApplicationsIdQuotaPut**](ApplicationsApi.md#ApplicationsIdQuotaPut) | **Put** /applications/{id}/quota | Updates a quota contraint for an application
[**ApplicationsIdScopeGet**](ApplicationsApi.md#ApplicationsIdScopeGet) | **Get** /applications/{id}/scope | Returns the scopes associated with an application
[**ApplicationsIdScopePost**](ApplicationsApi.md#ApplicationsIdScopePost) | **Post** /applications/{id}/scope | Adds an OAuth protected resource to an application
[**ApplicationsIdScopeScopeIdDelete**](ApplicationsApi.md#ApplicationsIdScopeScopeIdDelete) | **Delete** /applications/{id}/scope/{scopeId} | Remove an OAuth protected resource from an application
[**ApplicationsIdScopeScopeIdPut**](ApplicationsApi.md#ApplicationsIdScopeScopeIdPut) | **Put** /applications/{id}/scope/{scopeId} | Updates a scope associated with an application, sets default to true/false
[**ApplicationsOauthclientClientIdGet**](ApplicationsApi.md#ApplicationsOauthclientClientIdGet) | **Get** /applications/oauthclient/{clientId} | Get an application associated with an OAuth Client ID
[**ApplicationsPost**](ApplicationsApi.md#ApplicationsPost) | **Post** /applications | Creates a new application.



## ApplicationsGet

> []Application ApplicationsGet(ctx, optional)

Get the list of applications

Get the list of applications that are visible to the authenticated user.  The list of applications can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  apiid : Matches the application if the application is using the API, specified by ID  userid : Matches the application if the user has explicit access to the application, specified by ID  description : The application's description  email : The application's contact email address  enabled :  The enabled state of the application, one of: enabled, disabled  createdOn : The date the application was created on, time in ms, e.g.: 1372755998542  name : The name of the application  orgid : Matches the application if the application is part of the organization, specified by ID  phone : The application's contact phone  state : The application's current state, one of: approved, pending  The __op__ is an operation and is one of:  eq : Equal  ne : Not equal  gt :  Greater than  lt :  Less than  ge :  Greater than or equal  le :  Less than or equal  like : Like  gete :  Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

### Return type

[**[]Application**](Application.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdApikeysApikeyidPut

> ApiKey ApplicationsIdApikeysApikeyidPut(ctx, id, apikeyid, optional)

Updates an API Key

Updates the secret, enabled and Cors origin field.  The fields __id__, __createdBy__, __createdOn__ are read only.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose API Key is to be updated. | 
**apikeyid** | **string**| The ID of the API Key to be updated. | 
 **optional** | ***ApplicationsIdApikeysApikeyidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdApikeysApikeyidPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiKey**](ApiKey.md)|  | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdApikeysGet

> []ApiKey ApplicationsIdApikeysGet(ctx, id)

Returns the API Keys associated with an application

Returns the API Keys associated with an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose API Keys are to be returned. | 

### Return type

[**[]ApiKey**](APIKey.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdApikeysKeyIdDelete

> ApplicationsIdApikeysKeyIdDelete(ctx, id, keyId)

Delete an API Key

Deletes an API Key. Deleting an API key means that it will no longer be accepted for application authentication.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API Key ID to be deleted. | 
**keyId** | **string**|  | 

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


## ApplicationsIdApikeysPost

> ApiKey ApplicationsIdApikeysPost(ctx, id, optional)

Creates a new API Key and secret for the application

Creates a new API Key and secret for the application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application requiring an API Key. | 
 **optional** | ***ApplicationsIdApikeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdApikeysPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | [**optional.Interface of ApiKey**](ApiKey.md)| The APIKey to create | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdApisApiAccessIdApprovePost

> ApplicationsIdApisApiAccessIdApprovePost(ctx, id, apiAccessId)

Creates an API access request to an API for an application.

Approving user must be API Manager Administrator or an Organization Administrator of the application's organization with the correct privileges to approve API access requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
**apiAccessId** | **string**| The API access ID. | 

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


## ApplicationsIdApisApiAccessIdDelete

> ApiAccess ApplicationsIdApisApiAccessIdDelete(ctx, id, apiAccessId)

Deletes access to an API for an application

Permanently deletes access to an API for an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
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


## ApplicationsIdApisApiAccessIdPut

> ApiAccess ApplicationsIdApisApiAccessIdPut(ctx, id, apiAccessId, optional)

Updates access to an API for an application

Updates access to an API for an application.  Only __enabled__ may be modified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
**apiAccessId** | **string**| The API access ID. | 
 **optional** | ***ApplicationsIdApisApiAccessIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdApisApiAccessIdPutOpts struct


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


## ApplicationsIdApisGet

> []ApiAccess ApplicationsIdApisGet(ctx, id)

Get the list of APIs that the application can access

Get the list of APIs that the application can access.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 

### Return type

[**[]ApiAccess**](APIAccess.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdApisPost

> ApiAccess ApplicationsIdApisPost(ctx, id, optional)

Create a request for an application to access an API.

Only API Manager Administrator, or an Organization Administrator of the application's organization with the correct privileges, or the application manager may create API access requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
 **optional** | ***ApplicationsIdApisPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdApisPostOpts struct


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


## ApplicationsIdApprovePost

> ApplicationsIdApprovePost(ctx, id)

Approves a pending application request

Approving user must be API Manager Administrator or an Organization Administrator of the application's organization with the correct privileges to approve new application requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application request ID. | 

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


## ApplicationsIdAvailablescopesGet

> [][]map[string]interface{} ApplicationsIdAvailablescopesGet(ctx, id, optional)

Returns the scopes available to an application

Returns the OAuth scopes available to  an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth scopes are to be returned. | 
 **optional** | ***ApplicationsIdAvailablescopesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdAvailablescopesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiScope** | **optional.Bool**|  | [default to false]

### Return type

[**[][]map[string]interface{}**](array.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdDelete

> ApplicationsIdDelete(ctx, id)

Delete an application

Only managers of the application, API Manager Administrators, or Organization Administrators with enabled delegated application management privileges, may delete applications.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application to delete. | 

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


## ApplicationsIdExtclientsGet

> []ExternalClient ApplicationsIdExtclientsGet(ctx, id)

Returns the external clients associated with an application

Returns the external clients associated with an application. External clients are used when authenticating the application through a 3rd party OAuth service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose external clients are to be returned. | 

### Return type

[**[]ExternalClient**](ExternalClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdExtclientsObjectIdDelete

> ApplicationsIdExtclientsObjectIdDelete(ctx, id, objectId)

Delete an external client

Deletes an external client. Deleting a mapping means that it will no longer be accepted for application authentication.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose external client is to be deleted. | 
**objectId** | **string**| The ID of the external client entry to be deleted. | 

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


## ApplicationsIdExtclientsObjectIdPut

> ExternalClient ApplicationsIdExtclientsObjectIdPut(ctx, id, objectId, optional)

Updates an external client for the application

Updates an external client for the application. External clients are used when authenticating the application through a 3rd party OAuth service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose external client is to be updated. | 
**objectId** | **string**| The external client entry to be updated. | 
 **optional** | ***ApplicationsIdExtclientsObjectIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdExtclientsObjectIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ExternalClient**](ExternalClient.md)|  | 

### Return type

[**ExternalClient**](ExternalClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdExtclientsPost

> ExternalClient ApplicationsIdExtclientsPost(ctx, id, optional)

Maps a new external client to the application

Maps a new external client to the application. External clients are used when authenticating the application through a 3rd party OAuth service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application to map to an external client. | 
 **optional** | ***ApplicationsIdExtclientsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdExtclientsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExternalClient**](ExternalClient.md)|  | 

### Return type

[**ExternalClient**](ExternalClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdGet

> Application ApplicationsIdGet(ctx, id)

Get an application

Retrieves the details of an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the application to be returned. | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdImageGet

> ApplicationsIdImageGet(ctx, id)

Get the image for an application

Get the JPEG image associated with an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the application whose image is to be return | 

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


## ApplicationsIdImagePost

> ApplicationsIdImagePost(ctx, id, file, optional)

Adds a JPEG image to an application

Adds a JPEG image to an application with a MultiPart POST

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the application whose image is being added | 
**file** | ***os.File*****os.File**| The file uploaded in the POST body as an element in a multipart post | 
 **optional** | ***ApplicationsIdImagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdImagePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**|  | 

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


## ApplicationsIdOauthClientIdPut

> OAuthClient ApplicationsIdOauthClientIdPut(ctx, id, clientId, optional)

Updates an OAuth Credential for the application

Updates an OAuth Credential for the application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth credential is to be updated. | 
**clientId** | **string**| The OAuth Credential ID to be updated. | 
 **optional** | ***ApplicationsIdOauthClientIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdOauthClientIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OAuthClient**](OAuthClient.md)|  | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdOauthClientidNewsecretPut

> OAuthClient ApplicationsIdOauthClientidNewsecretPut(ctx, id, clientid)

Updates an OAuth Credential for an application by generating a new secret

Updates an OAuth Credential for an application by generating a new client secret.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth credential is to be updated with a new secret | 
**clientid** | **string**| The OAuth Credential ID to be updated with a new secret | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdOauthGet

> []OAuthClient ApplicationsIdOauthGet(ctx, id)

Returns the OAuth Credentials associated with an application

Returns the OAuth Credentials associated with an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth Credentials are to be returned. | 

### Return type

[**[]OAuthClient**](OAuthClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdOauthOauthIdDelete

> ApplicationsIdOauthOauthIdDelete(ctx, id, oauthId)

Delete an OAuth client ID and secret

Deletes an OAuth client ID and secret. Deleting an OAuth client ID and secret means that it will no longer be accepted for OAuth application authentication.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth Credential is to be deleted. | 
**oauthId** | **string**| The OAuth Client ID to be deleted. | 

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


## ApplicationsIdOauthPost

> OAuthClient ApplicationsIdOauthPost(ctx, id, body)

Creates a new OAuth client ID and secret for the application

Creates a new OAuth client ID and secret for the application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application requiring an OAuth client ID and secret. | 
**body** | [**OAuthClient**](OAuthClient.md)| The OAuth credential to create | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdOauthresourceGet

> []OAuthResource ApplicationsIdOauthresourceGet(ctx, id)

Returns the OAuth protected resources (scopes) associated with an application

Returns the OAuth protected resources (scopes) associated with an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth protected resources (Scopes) are to be returned. | 

### Return type

[**[]OAuthResource**](OAuthResource.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdOauthresourcePost

> OAuthResource ApplicationsIdOauthresourcePost(ctx, id, body)

Adds an OAuth protected resource to an application

An application must define which OAuth Protected resources it wants to access. These resources will define the scope of the application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application requiring an OAuth protected resource. | 
**body** | [**OAuthResource**](OAuthResource.md)| The OAuth protected resource to add to the application | 

### Return type

[**OAuthResource**](OAuthResource.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdOauthresourceResourceIdDelete

> ApplicationsIdOauthresourceResourceIdDelete(ctx, id, resourceId)

Remove an OAuth protected resource from an application

Removes the association between an application and an OAuth protected resource on the API Server. The application will no longer have the scope associated with the resource. Tokens issued prior to the removal will still be scoped for the resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose protected resource is to be removed. | 
**resourceId** | **string**| The uri of the OAuth protected resource to be disassociated from the application. | 

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


## ApplicationsIdOauthresourceResourceIdPut

> OAuthResource ApplicationsIdOauthresourceResourceIdPut(ctx, id, resourceId, body)

Updates a protected resource associate with an application, sets enabled to true/false

An OAuth Protected resource associated with an application can be enabled or disabled with this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application requiring an updated OAuth protected resource. | 
**resourceId** | **string**| The ID of the OAuth protected resource to update | 
**body** | [**OAuthResource**](OAuthResource.md)| The updated OAuth protected resource | 

### Return type

[**OAuthResource**](OAuthResource.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdPermissionsGet

> []PermissionDto ApplicationsIdPermissionsGet(ctx, id)

Get the list of permissions.

Get the access-control list (ACL) for the application. Callers with view-only privilege can only access their own permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 

### Return type

[**[]PermissionDto**](PermissionDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdPermissionsPermIdDelete

> ApplicationsIdPermissionsPermIdDelete(ctx, id, permId)

Remove a permission

Remove an existing access-control entry from the application's ACL. Management privilege required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
**permId** | **string**| The permission ID. | 

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


## ApplicationsIdPermissionsPermIdPut

> PermissionDto ApplicationsIdPermissionsPermIdPut(ctx, id, permId, optional)

Modify a permission

Modify an existing access-control entry from the application's ACL. Management privilege required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
**permId** | **string**| The permission ID. | 
 **optional** | ***ApplicationsIdPermissionsPermIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdPermissionsPermIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PermissionDto**](PermissionDto.md)|  | 

### Return type

[**PermissionDto**](PermissionDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdPermissionsPost

> PermissionDto ApplicationsIdPermissionsPost(ctx, id, optional)

Create a new permission.

Add a new access-control entry to the application's ACL. Management privilege required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The application ID. | 
 **optional** | ***ApplicationsIdPermissionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdPermissionsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PermissionDto**](PermissionDto.md)|  | 

### Return type

[**PermissionDto**](PermissionDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdPut

> Application ApplicationsIdPut(ctx, id, optional)

Update an application

Only managers of the application, API Manager Administrators, or Organization Administrators with enabled delegated application management privileges, may update an application. Note, if a field is omitted from the payload, or its value is set to null, the existing value for this field will be retained.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the application to be updated | 
 **optional** | ***ApplicationsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Application**](Application.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdQuotaDelete

> ApplicationsIdQuotaDelete(ctx, id)

Deletes a quota from an application

Deletes a quota from an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application who&#39;s quota constraint is to be deleted. | 

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


## ApplicationsIdQuotaGet

> []QuotaDto ApplicationsIdQuotaGet(ctx, id)

Returns the quota associated with an application.

Returns the quota associated with an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application who&#39;s quota constraints are to be returned. | 

### Return type

[**[]QuotaDto**](QuotaDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdQuotaPost

> QuotaDto ApplicationsIdQuotaPost(ctx, id, optional)

Creates a new quota constraint for the application

Creates a new quota constraint for the application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application for the quota contraint. | 
 **optional** | ***ApplicationsIdQuotaPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdQuotaPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QuotaDto**](QuotaDto.md)|  | 

### Return type

[**QuotaDto**](QuotaDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdQuotaPut

> QuotaDto ApplicationsIdQuotaPut(ctx, id, optional)

Updates a quota contraint for an application

Updates a quota contraint for the given application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application who&#39;s quota is to be updated. | 
 **optional** | ***ApplicationsIdQuotaPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsIdQuotaPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QuotaDto**](QuotaDto.md)|  | 

### Return type

[**QuotaDto**](QuotaDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdScopeGet

> [][]map[string]interface{} ApplicationsIdScopeGet(ctx, id)

Returns the scopes associated with an application

Returns the OAuth scopes associated with an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose OAuth protected resources (Scopes) are to be returned. | 

### Return type

[**[][]map[string]interface{}**](array.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdScopePost

> OAuthAppScope ApplicationsIdScopePost(ctx, id, body)

Adds an OAuth protected resource to an application

An application must define which scopes it wants to access. These define the scope of the application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application requiring an OAuth scope. | 
**body** | [**OAuthAppScope**](OAuthAppScope.md)| The OAuth Scope to add to the application | 

### Return type

[**OAuthAppScope**](OAuthAppScope.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsIdScopeScopeIdDelete

> ApplicationsIdScopeScopeIdDelete(ctx, id, scopeId)

Remove an OAuth protected resource from an application

Removes the association between an application and an OAuth protected resource on the API Server. The application will no longer have the scope associated with the resource. Tokens issued prior to the removal will still be scoped for the resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application whose protected resource is to be removed. | 
**scopeId** | **string**| The id of the Scope to be disassociated from the application. | 

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


## ApplicationsIdScopeScopeIdPut

> OAuthAppScope ApplicationsIdScopeScopeIdPut(ctx, id, scopeId, body)

Updates a scope associated with an application, sets default to true/false

An OAuth Scope associated with an application can be set or unset as a default scope with this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of application requiring an updated OAuth protected resource. | 
**scopeId** | **string**| The ID of the Application Scope to update | 
**body** | [**OAuthAppScope**](OAuthAppScope.md)| The updated OAuth protected resource | 

### Return type

[**OAuthAppScope**](OAuthAppScope.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsOauthclientClientIdGet

> Application ApplicationsOauthclientClientIdGet(ctx, clientId)

Get an application associated with an OAuth Client ID

Retrieves the application associated with an OAuth Client ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string**| The OAuth Client ID associated with the Application. | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsPost

> Application ApplicationsPost(ctx, optional)

Creates a new application.

Creates a new application.  New applications may need to be approved using [/approve](#APIApplicationsapproveApp).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplicationRequest**](ApplicationRequest.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

