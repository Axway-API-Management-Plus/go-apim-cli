# \APIProxyRegistrationApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxiesExportIdGet**](APIProxyRegistrationApi.md#ProxiesExportIdGet) | **Get** /proxies/export/{id} | Downloads an API export.
[**ProxiesExportPost**](APIProxyRegistrationApi.md#ProxiesExportPost) | **Post** /proxies/export | Creates an API export.
[**ProxiesGet**](APIProxyRegistrationApi.md#ProxiesGet) | **Get** /proxies | Queries a list of frontend API.
[**ProxiesGrantaccessPost**](APIProxyRegistrationApi.md#ProxiesGrantaccessPost) | **Post** /proxies/grantaccess | Macro function to grant API access.
[**ProxiesIdDelete**](APIProxyRegistrationApi.md#ProxiesIdDelete) | **Delete** /proxies/{id} | Deletes an API proxy.
[**ProxiesIdDeprecatePost**](APIProxyRegistrationApi.md#ProxiesIdDeprecatePost) | **Post** /proxies/{id}/deprecate | Deprecates the API.
[**ProxiesIdGet**](APIProxyRegistrationApi.md#ProxiesIdGet) | **Get** /proxies/{id} | Gets a frontend API by ID.
[**ProxiesIdImageGet**](APIProxyRegistrationApi.md#ProxiesIdImageGet) | **Get** /proxies/{id}/image | Gets the image for the API.
[**ProxiesIdImagePost**](APIProxyRegistrationApi.md#ProxiesIdImagePost) | **Post** /proxies/{id}/image | Set the image for the frontend API.
[**ProxiesIdOperationsGet**](APIProxyRegistrationApi.md#ProxiesIdOperationsGet) | **Get** /proxies/{id}/operations | Gets a list of methods that are avilable to the API proxy.
[**ProxiesIdOperationsOperationIdDelete**](APIProxyRegistrationApi.md#ProxiesIdOperationsOperationIdDelete) | **Delete** /proxies/{id}/operations/{operationId} | Deletes an API method by ID.
[**ProxiesIdOperationsOperationIdGet**](APIProxyRegistrationApi.md#ProxiesIdOperationsOperationIdGet) | **Get** /proxies/{id}/operations/{operationId} | Gets an API method by ID.
[**ProxiesIdOperationsOperationIdPut**](APIProxyRegistrationApi.md#ProxiesIdOperationsOperationIdPut) | **Put** /proxies/{id}/operations/{operationId} | Updates an API proxy operation.
[**ProxiesIdPublishPost**](APIProxyRegistrationApi.md#ProxiesIdPublishPost) | **Post** /proxies/{id}/publish | Publish the API.
[**ProxiesIdPut**](APIProxyRegistrationApi.md#ProxiesIdPut) | **Put** /proxies/{id} | Updates an API proxy.
[**ProxiesIdUndeprecatePost**](APIProxyRegistrationApi.md#ProxiesIdUndeprecatePost) | **Post** /proxies/{id}/undeprecate | Undeprecates the API.
[**ProxiesIdUnpublishPost**](APIProxyRegistrationApi.md#ProxiesIdUnpublishPost) | **Post** /proxies/{id}/unpublish | Unpublish the API.
[**ProxiesImportFromUrlPost**](APIProxyRegistrationApi.md#ProxiesImportFromUrlPost) | **Post** /proxies/importFromUrl | Imports a previously exported API.
[**ProxiesImportPost**](APIProxyRegistrationApi.md#ProxiesImportPost) | **Post** /proxies/import | Imports a previously exported API.
[**ProxiesPost**](APIProxyRegistrationApi.md#ProxiesPost) | **Post** /proxies | Creates a new API proxy from a backend API.
[**ProxiesPromotePost**](APIProxyRegistrationApi.md#ProxiesPromotePost) | **Post** /proxies/promote | Invokes the internal API promotion policy for the specified API.
[**ProxiesUpgradeIdPost**](APIProxyRegistrationApi.md#ProxiesUpgradeIdPost) | **Post** /proxies/upgrade/{id} | Upgrades an existing frontend API to a newer frontend API.



## ProxiesExportIdGet

> ApiPromotion ProxiesExportIdGet(ctx, id, filename)

Downloads an API export.

The API export is produced from [/exportApis](APIProxyRegistration.html#APIProxyRegistrationexportApis).  If __filename__ is supplied, the download will use it as the `Content-Disposition` filename attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The export identifier. | 
**filename** | **string**| The export will be downloaded using a Content-Dispostion using the supplied filename | 

### Return type

[**ApiPromotion**](APIPromotion.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesExportPost

> ProxiesExportPost(ctx, id, password, filename)

Creates an API export.

Creates an export for use in promoting the API to a new environment.  The export contains the frontend [VirtualizedAPI](VirtualizedAPI.html), their settings, and all backend [APIDefinition](APIDefinition.html) that are required for the frontend API.  If **password** is supplied, the exported file will be encrypted with the password.  If successful, returns **201 Created**, and the HTTP `Location` header contain the of the URL of the export. The export is temporary, and may only be downloaded once.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**[]string**](string.md)| The frontend API identifier(s) to export. | 
**password** | **string**| Encrypts the list of API using the password. | 
**filename** | **string**| Optional filename to use in response. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesGet

> []VirtualizedApi ProxiesGet(ctx, optional)

Queries a list of frontend API.

Returns a list of API that are visible to the authenticated user.  The list of API can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  name :      The name of the API  apiid :      Matches the API if the API is virtualized from the specified backend API  createdOn :      The date the user was created on, time in ms, e.g.: 1372755998542  deprecated :      The deprecated state of the API, one of: true or false  retired :      The retired state of the API, one of: true or false  state :      The API's state, one of: unpublished, pending, or published  The __op__ is an operation and is one of:  eq :      Equal  ne :      Not equal  gt :      Greater than  lt :      Less than  ge :      Greater than or equal  le :      Less than or equal  like :      Like  gete :      Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProxiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProxiesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

### Return type

[**[]VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesGrantaccessPost

> ProxiesGrantaccessPost(ctx, action, apiId, grantOrgId, grantApiId)

Macro function to grant API access.

Function to macro-apply access to selected API.  The access can be granted to organizations or entities having access to specified API. If **action** is _all_orgs_, access will be granted to all organizations; if **action** is _orgs_, access will be granted to the organization(s) specified by **grantOrganizations**; if **action** is _orgs\\_with\\_apis_, access will be granted to the organizations with access to the apis specified by **grantApis**.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string**| Grant action to perform. Possible values are: all\\\\_orgs, orgs, and orgs\\\\_with\\\\_apis. | 
**apiId** | [**[]string**](string.md)| List of API ID to which access will be granted. | 
**grantOrgId** | [**[]string**](string.md)| List of target organization ID to which access to _apiId_ will be granted (action is _orgs_) | 
**grantApiId** | [**[]string**](string.md)| List of API ID to which access to to _apiId_ will be granted (action is _orgs\\_with\\_apis_). | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdDelete

> ProxiesIdDelete(ctx, id)

Deletes an API proxy.

Deletes an API proxy, removing all API access in the process.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 

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


## ProxiesIdDeprecatePost

> VirtualizedApi ProxiesIdDeprecatePost(ctx, id, optional)

Deprecates the API.

Only an API Administrator may deprecate an API, and only _published_ API may be deprecated.  When an API is _deprecated_, the API can still be used, but users will be informed that the API has been deprecated.  Optionally, a **retirementDate** may be set which will schedule the API to be automatically retired and removed from use. If specified, the **retirementDate** should be in the ISO-8601 format of yyyy-MM-ddTHH:mm:ssZ (e.g. 2015-01-01T12:00:00Z).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
 **optional** | ***ProxiesIdDeprecatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProxiesIdDeprecatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retirementDate** | **optional.String**| Optional API retirement date specified in supported ISO-8601 format.  Set to the past to retire immediately. | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdGet

> VirtualizedApi ProxiesIdGet(ctx, id)

Gets a frontend API by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdImageGet

> ProxiesIdImageGet(ctx, id)

Gets the image for the API.

Returns the jpeg image associated with the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 

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


## ProxiesIdImagePost

> ProxiesIdImagePost(ctx, id, optional)

Set the image for the frontend API.

Set the jpeg image to be associated with the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
 **optional** | ***ProxiesIdImagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProxiesIdImagePostOpts struct


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


## ProxiesIdOperationsGet

> []VirtualizedApiMethod ProxiesIdOperationsGet(ctx, id)

Gets a list of methods that are avilable to the API proxy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 

### Return type

[**[]VirtualizedApiMethod**](VirtualizedAPIMethod.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdOperationsOperationIdDelete

> ProxiesIdOperationsOperationIdDelete(ctx, id, operationId)

Deletes an API method by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
**operationId** | **string**| The frontend API method identifier. | 

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


## ProxiesIdOperationsOperationIdGet

> VirtualizedApiMethod ProxiesIdOperationsOperationIdGet(ctx, id, operationId)

Gets an API method by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
**operationId** | **string**| The frontend API method identifier. | 

### Return type

[**VirtualizedApiMethod**](VirtualizedAPIMethod.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdOperationsOperationIdPut

> VirtualizedApiMethod ProxiesIdOperationsOperationIdPut(ctx, id, operationId, body)

Updates an API proxy operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
**operationId** | **string**| The frontend API method identifier. | 
**body** | [**VirtualizedApiMethod**](VirtualizedApiMethod.md)| The method to update. | 

### Return type

[**VirtualizedApiMethod**](VirtualizedAPIMethod.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdPublishPost

> VirtualizedApi ProxiesIdPublishPost(ctx, id, name, vhost)

Publish the API.

If called by an API Administrator, then the API state will be _published_, otherwise the API state will be _pending_, and an email notification will be sent to the API Administrators, notifying them of the event. Optionally, on publishing, a new **name** for the API may be specified.  Similarly, an optional **vhost** may be specified.  The **vhost** is an externally resolvable virtual host from which the API will be accessed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
**name** | **string**| The name on which to publish this API. | 
**vhost** | **string**| The optional virtual host on which to publish this API. | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdPut

> VirtualizedApi ProxiesIdPut(ctx, id, body)

Updates an API proxy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 
**body** | [**VirtualizedApi**](VirtualizedApi.md)| The virtualized API to update | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdUndeprecatePost

> VirtualizedApi ProxiesIdUndeprecatePost(ctx, id)

Undeprecates the API.

Only an API Administrator may undeprecate an API, and only _published_ API, that are deprecated, may be undeprecated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesIdUnpublishPost

> VirtualizedApi ProxiesIdUnpublishPost(ctx, id)

Unpublish the API.

Only an API Administrator may unpublish an API.  When an API is _unpublished_, all access to the API is revoked from all applications, and all organizations, save the API development organization that owns the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The frontend API identifier. | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesImportFromUrlPost

> ProxiesImportFromUrlPost(ctx, organizationId, url, password)

Imports a previously exported API.

Imports API, previously exported using [/exportApis](APIProxyRegistration.html#APIProxyRegistrationexportApis).  If the API was exported using a password, then the file is encrypted, and a **password** argument must be provided to decrypt.  The import will create  [VirtualizedAPI](VirtualizedAPI.html), their settings, and all backend [APIDefinition](APIDefinition.html) necessary to support the frontend API.  The **url** should be a [data URI scheme](http://en.wikipedia.org/wiki/Data_URI_scheme).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| The organization identifier. | 
**url** | **string**| The data URI. | 
**password** | **string**| Optional password to decrypt the import file. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesImportPost

> ProxiesImportPost(ctx, organizationId, password, file)

Imports a previously exported API.

Imports API, previously exported using [/exportApis](APIProxyRegistration.html#APIProxyRegistrationexportApis).  If the API was exported using a password, then the file is encrypted, and a **password** argument must be provided to decrypt.  The import will create  [VirtualizedAPI](VirtualizedAPI.html), their settings, and all backend [APIDefinition](APIDefinition.html) necessary to support the frontend API.  This method is similar to [/importFromUrl](APIProxyRegistration.html#APIProxyRegistrationimportFromUrl), save that this method supports traditional form-based file upload, using `multipart/form-data`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string**| The organization identifier. | 
**password** | **string**| Optional password to decrypt the import file. | 
**file** | ***os.File*****os.File**| The data file to import. | 

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


## ProxiesPost

> VirtualizedApi ProxiesPost(ctx, body)

Creates a new API proxy from a backend API.

The [VirtualizedAPI apiId](VirtualizedAPI.html#apiId) is required.  If creating the APIas an API administrator, the [VirtualizedAPI organizationId](VirtualizedAPI.html#organizationId) must also be specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**VirtualizedApi**](VirtualizedApi.md)| The frontend API to create. | 

### Return type

[**VirtualizedApi**](VirtualizedAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesPromotePost

> ProxiesPromotePost(ctx, apiId)

Invokes the internal API promotion policy for the specified API.

In API Manager, API promotion must first be enabled in Settings. Also, in Policy Studio (Server Settings -> API Manager -> API Promotion) a promotion policy must be selected. By default a sample promotion policy is installed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | [**[]string**](string.md)| The frontend API identifier(s) to promote | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxiesUpgradeIdPost

> ProxiesUpgradeIdPost(ctx, id, upgradeApiId, deprecate, retire, retirementDate)

Upgrades an existing frontend API to a newer frontend API.

During an API lifecycle, it is necessary to upgrade users to use a newer frontend API.  The idea being that the old frontend API should be phased-out, and developers should move their applications to use the newer frontend API. This method assigns all organizations and applications the same access to the new frontend API (identified by **upgradeApiId**) that they have to the old API (identified by **id**). Optionally, the old frontend API may be deprecated or retired using **deprecate**, **retire**, or scheduled to be retired using **retirementDate**. If specified, the **retirementDate** should be in the ISO-8601 format of yyyy-MM-ddTHH:mm:ssZ (e.g. 2015-01-01T12:00:00Z).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**upgradeApiId** | **string**| The id of the frontend API which will be used to upgrade this virtualized API | 
**deprecate** | **bool**| Specifies whether or not the API being upgraded should be depreated | 
**retire** | **bool**| Specifies whether or not the API being upgraded should be retired | 
**retirementDate** | **string**| Specifies the retirement date of the the API being upgraded if its being retired | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

