# \APIDiscoveryApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveryApisGet**](APIDiscoveryApi.md#DiscoveryApisGet) | **Get** /discovery/apis | Lists all APIs/services virtualised in the API Server.
[**DiscoveryOauthresourcesGet**](APIDiscoveryApi.md#DiscoveryOauthresourcesGet) | **Get** /discovery/oauthresources | Gets a list OAuth protected resources and their associated scopes.
[**DiscoveryScopesGet**](APIDiscoveryApi.md#DiscoveryScopesGet) | **Get** /discovery/scopes | Retrieves every resource on the API Server that is protected by OAuth, and the scopes that cover those resources
[**DiscoverySdkIdPlatformGet**](APIDiscoveryApi.md#DiscoverySdkIdPlatformGet) | **Get** /discovery/sdk/{id}/{platform} | Generates an SDK package for the specified API based on the type of the client requested
[**DiscoverySwaggerApiIdIdGet**](APIDiscoveryApi.md#DiscoverySwaggerApiIdIdGet) | **Get** /discovery/swagger/api/id/{id} | Retrieves an extended Swagger feed for the specified API.
[**DiscoverySwaggerApiNameGet**](APIDiscoveryApi.md#DiscoverySwaggerApiNameGet) | **Get** /discovery/swagger/api/{name} | Retrieves an extended Swagger feed for the specified API.
[**DiscoverySwaggerApisGet**](APIDiscoveryApi.md#DiscoverySwaggerApisGet) | **Get** /discovery/swagger/apis | Convenience method for retrieving all Swagger feeds for all virtualised services.
[**DiscoverySwaggerApisIdImageGet**](APIDiscoveryApi.md#DiscoverySwaggerApisIdImageGet) | **Get** /discovery/swagger/apis/{id}/image | Retrieves the API image
[**DiscoverySwaggerApisIdServiceDefinitionGet**](APIDiscoveryApi.md#DiscoverySwaggerApisIdServiceDefinitionGet) | **Get** /discovery/swagger/apis/{id}/service-definition | Retrieves the service definition of the API. 



## DiscoveryApisGet

> []DiscoveryApi DiscoveryApisGet(ctx, )

Lists all APIs/services virtualised in the API Server.

Lists all APIs/services virtualised in the API Server. API Administrators see all APIs/services. Users see APIs/services for their organization.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DiscoveryApi**](DiscoveryAPI.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryOauthresourcesGet

> DiscoveryOauthresourcesGet(ctx, )

Gets a list OAuth protected resources and their associated scopes.

Gets a list OAuth protected resources and their associated scopes.

### Required Parameters

This endpoint does not need any parameter.

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


## DiscoveryScopesGet

> []OAuthProtectedResource DiscoveryScopesGet(ctx, )

Retrieves every resource on the API Server that is protected by OAuth, and the scopes that cover those resources

Retrieves every resource on the API Server that is protected by OAuth, and the scopes that cover those resources. Only API Administrators will be able to retrieve information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]OAuthProtectedResource**](OAuthProtectedResource.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverySdkIdPlatformGet

> DiscoverySdkIdPlatformGet(ctx, id, platform, optional)

Generates an SDK package for the specified API based on the type of the client requested

Generates a client SDK package for the specified API based on the platform. Supported platform are Android, iOS-swift, NodeJS, Titanium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The name of the API to generate the client SDK package for | 
**platform** | **string**| Generated client type, one of: android, iOS-swift, nodejS, titanium | 
 **optional** | ***DiscoverySdkIdPlatformGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoverySdkIdPlatformGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **packagename** | **optional.String**| Java package name generated only for Android platform. It must be a valid java package name. | [default to com.axway]

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


## DiscoverySwaggerApiIdIdGet

> DiscoverySwaggerApiIdIdGet(ctx, id, optional)

Retrieves an extended Swagger feed for the specified API.

Retrieves an extended Swagger feed for the specified API. API Administrators will always see the API. Users will only see the API if it is available for their organization.If __filename__ is supplied, the download will use it as the `Content-Disposition` filename attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The unique ID of the API to return | 
 **optional** | ***DiscoverySwaggerApiIdIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoverySwaggerApiIdIdGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **optional.String**| Override the default filename for download | 
 **swaggerVersion** | **optional.String**| The Swagger version of the feed, either 1.1 (default) or 2.0. | [default to 1.1]
 **extensions** | **optional.Bool**| If true, extensions such as the x-axway object are returned in the Swagger definitions (default&#x3D;true) | [default to true]

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


## DiscoverySwaggerApiNameGet

> DiscoverySwaggerApiNameGet(ctx, name, optional)

Retrieves an extended Swagger feed for the specified API.

Retrieves an extended Swagger feed for the specified API. API Administrators will always see the API. Users will only see the API if it is available for their organization.If __filename__ is supplied, the download will use it as the `Content-Disposition` filename attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the API to return | 
 **optional** | ***DiscoverySwaggerApiNameGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoverySwaggerApiNameGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **optional.String**| Override the default filename for download | 
 **apiVersion** | **optional.String**| The version of the API. Should always be provided if there is more than one version | 
 **swaggerVersion** | **optional.String**| The Swagger version of the feed, either 1.1 (default) or 2.0. | [default to 1.1]
 **extensions** | **optional.Bool**| If true, extensions such as the x-axway object are returned in the Swagger definitions (default&#x3D;true) | [default to true]

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


## DiscoverySwaggerApisGet

> DiscoverySwaggerApisGet(ctx, optional)

Convenience method for retrieving all Swagger feeds for all virtualised services.

Convenience method for retrieving all Swagger feeds for all virtualised services that are visible to the authenticated user.  The list of APIs can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  id :        Matches the API by the specified ID  name :        Matches the API by the specified name  description :        Matches the API by the specified description  summary :        Matches the API by the specified summary  version :        Matches the API by the specified version  type :        Matches the API by the specified type. Type can be 'rest' or 'wsdl'  resourcepath :        Matches the API by the specified inbound path  taggroup :        Matches the API by the specified tag group  tag :        Matches the API by the specified tag value  methodtaggroup :        Matches the API by the specified method tag group, i.e. if the API contains a method that contains a tag group matching that specified  methodtag :        Matches the API by the specified method tag value, i.e. if the API contains a method that contains a tag value matching that specified  The __op__ is an operation and is one of:  eq :    Equal  ne :    Not equal  gt :     Greater than  lt :     Less than  ge :     Greater than or equal  le :     Less than or equal  like :    Like  gete :     Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiscoverySwaggerApisGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoverySwaggerApisGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The Swagger API version | 
 **swaggerVersion** | **optional.String**| The Swagger version | [default to 1.2]
 **extensions** | **optional.Bool**| If true, extensions such as the x-axway object are returned in the Swagger definitions (default&#x3D;true) | [default to true]
 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

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


## DiscoverySwaggerApisIdImageGet

> DiscoverySwaggerApisIdImageGet(ctx, id)

Retrieves the API image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifier | 

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


## DiscoverySwaggerApisIdServiceDefinitionGet

> DiscoverySwaggerApisIdServiceDefinitionGet(ctx, id)

Retrieves the service definition of the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The API identifier | 

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

