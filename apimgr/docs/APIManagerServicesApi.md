# \APIManagerServicesApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsGet**](APIManagerServicesApi.md#AlertsGet) | **Get** /alerts | Gets the alerts configured for the API Manager
[**AlertsPost**](APIManagerServicesApi.md#AlertsPost) | **Post** /alerts | Updates the API Manager alerts configuration
[**AppinfoGet**](APIManagerServicesApi.md#AppinfoGet) | **Get** /appinfo | Gets API Manager feature information.
[**CertinfoPost**](APIManagerServicesApi.md#CertinfoPost) | **Post** /certinfo | Extracts certificate information from the supplied data
[**ConfigGet**](APIManagerServicesApi.md#ConfigGet) | **Get** /config | Gets API Manager configuration
[**ConfigPut**](APIManagerServicesApi.md#ConfigPut) | **Put** /config | Updates the API Manager configuration
[**ConnectorsConnectorIdGet**](APIManagerServicesApi.md#ConnectorsConnectorIdGet) | **Get** /connectors/{connectorId} | Return a list of APIs for the specified connector
[**ConnectorsConnectorIdLoginPost**](APIManagerServicesApi.md#ConnectorsConnectorIdLoginPost) | **Post** /connectors/{connectorId}/login | Login to an external service from which APIs will be imported
[**ConnectorsGet**](APIManagerServicesApi.md#ConnectorsGet) | **Get** /connectors | Return a list of API connectors
[**FiledataPost**](APIManagerServicesApi.md#FiledataPost) | **Post** /filedata | Returns the DataURI representation of the uploaded file
[**LicenseGet**](APIManagerServicesApi.md#LicenseGet) | **Get** /license | Checks that the API Manager has a valid license
[**ListenersGet**](APIManagerServicesApi.md#ListenersGet) | **Get** /listeners | Gets the API Manager listeners
[**OauthclientprofilesGet**](APIManagerServicesApi.md#OauthclientprofilesGet) | **Get** /oauthclientprofiles | Get a list of OAuth profiles for use in backend API authorisation
[**PoliciesGet**](APIManagerServicesApi.md#PoliciesGet) | **Get** /policies | Gets a list of the specified policies
[**RemotehostsGet**](APIManagerServicesApi.md#RemotehostsGet) | **Get** /remotehosts | Returns a list of remote hosts
[**RemotehostsIdDelete**](APIManagerServicesApi.md#RemotehostsIdDelete) | **Delete** /remotehosts/{id} | Deletes a remote host.
[**RemotehostsIdPut**](APIManagerServicesApi.md#RemotehostsIdPut) | **Put** /remotehosts/{id} | Updates a remote host
[**RemotehostsPost**](APIManagerServicesApi.md#RemotehostsPost) | **Post** /remotehosts | Creates a new remote host
[**ServiceDiscoveryInstanceTypePost**](APIManagerServicesApi.md#ServiceDiscoveryInstanceTypePost) | **Post** /service-discovery/{instance}/{type} | Returns a list of services hosted on the specified Gateway instance
[**SysconfigGet**](APIManagerServicesApi.md#SysconfigGet) | **Get** /sysconfig | Gets API Manager system configuration
[**SysconfigPut**](APIManagerServicesApi.md#SysconfigPut) | **Put** /sysconfig | Update API Manager system configuration
[**TitleGet**](APIManagerServicesApi.md#TitleGet) | **Get** /title | Gets the API Manager&#39;s title
[**TokenstoresGet**](APIManagerServicesApi.md#TokenstoresGet) | **Get** /tokenstores | Gets a list of Token Stores
[**TopologyPost**](APIManagerServicesApi.md#TopologyPost) | **Post** /topology | Retrieves the Topology from the specified Admin Node Manager



## AlertsGet

> AlertConfig AlertsGet(ctx, )

Gets the alerts configured for the API Manager

Gets the alerts configured for the API Manager that shows which alerts are enabled/disabled for the application.  Only the API Administrator may call this method.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AlertConfig**](AlertConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPost

> AlertConfig AlertsPost(ctx, optional)

Updates the API Manager alerts configuration

Updates the API Manager alerts configuration.  Only the API Administrator may call this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AlertsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AlertConfig**](AlertConfig.md)|  | 

### Return type

[**AlertConfig**](AlertConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppinfoGet

> Config AppinfoGet(ctx, )

Gets API Manager feature information.

Returns an API Manager configuration object describing the application.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Config**](Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertinfoPost

> CaCert CertinfoPost(ctx, optional)

Extracts certificate information from the supplied data

Extracts certificate information from the supplied data.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertinfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CertinfoPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 
 **passphrase** | **optional.String**|  | 
 **inbound** | **optional.Bool**|  | 
 **outbound** | **optional.Bool**|  | 

### Return type

[**CaCert**](CACert.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGet

> Config ConfigGet(ctx, )

Gets API Manager configuration

Returns an API Manager configuration object.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Config**](Config.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPut

> Config ConfigPut(ctx, optional)

Updates the API Manager configuration

Updates the API Manager configuration.  Only the API Administrator may call this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Config**](Config.md)|  | 

### Return type

[**Config**](Config.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorsConnectorIdGet

> []map[string]map[string]interface{} ConnectorsConnectorIdGet(ctx, connectorId)

Return a list of APIs for the specified connector

Return a list of APIs for the specified connector.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string**| ID of the connector for which APIs should be returned. Connector IDs can be retrieved by calling /connectors. | 

### Return type

[**[]map[string]map[string]interface{}**](map.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorsConnectorIdLoginPost

> string ConnectorsConnectorIdLoginPost(ctx, connectorId, username, password)

Login to an external service from which APIs will be imported

Login to an external service from which APIs will be imported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string**| ID of the API connector. Connector IDs can be retrieved by calling /connectors. | 
**username** | **string**| External service username | 
**password** | **string**| External service password | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorsGet

> []map[string]map[string]interface{} ConnectorsGet(ctx, )

Return a list of API connectors

Return a list of API connectors.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]map[string]map[string]interface{}**](map.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FiledataPost

> FiledataPost(ctx, optional)

Returns the DataURI representation of the uploaded file

Returns the DataURI representation of the uploaded file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FiledataPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FiledataPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## LicenseGet

> string LicenseGet(ctx, )

Checks that the API Manager has a valid license

Returns an API Manager license configuration object.  Does not require authentication.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListenersGet

> []PortalTrafficListener ListenersGet(ctx, )

Gets the API Manager listeners

Returns a list of API Manager listeners.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PortalTrafficListener**](PortalTrafficListener.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthclientprofilesGet

> []ReferencedEntity OauthclientprofilesGet(ctx, )

Get a list of OAuth profiles for use in backend API authorisation

Return a list of OAuth Client Profiles for use in authorising API access to backend APIs.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ReferencedEntity**](ReferencedEntity.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGet

> []ReferencedEntity PoliciesGet(ctx, type_)

Gets a list of the specified policies

Returns the list of policies (of the specified type) that are available to Portal-registered APIs.  __type__ is one of: faulthandler, globalrequest, globalresponse, request, routing, response, promotion, authentication or token-info.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The type of policy to return. Possible values are: &#39;faulthandler&#39;, &#39;globalrequest&#39;, &#39;globalresponse&#39;, &#39;request&#39;, &#39;routing&#39;, &#39;response&#39;, &#39;authentication&#39;, &#39;oauthtokeninfo&#39;, &#39;promotion&#39; | 

### Return type

[**[]ReferencedEntity**](ReferencedEntity.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotehostsGet

> []RemoteHost RemotehostsGet(ctx, )

Returns a list of remote hosts

Returns a list of API Manager-registered remote hosts.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RemoteHost**](RemoteHost.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotehostsIdDelete

> RemotehostsIdDelete(ctx, id)

Deletes a remote host.

Deletes an API Manager-registered remote host. Dynamically removes the remote host from the API Gateway runtime.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The remote host identifier. | 

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


## RemotehostsIdPut

> RemoteHost RemotehostsIdPut(ctx, id, optional)

Updates a remote host

Updates an API Manager-registered remote host. Dynamically updates the API Gateway runtime so that the new remote host settings are available.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The remote host identifier. | 
 **optional** | ***RemotehostsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemotehostsIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RemoteHost**](RemoteHost.md)|  | 

### Return type

[**RemoteHost**](RemoteHost.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotehostsPost

> RemoteHost RemotehostsPost(ctx, optional)

Creates a new remote host

Creates a new API Manager-regsitered remote host. Dynamically updates the API Gateway runtime so that the remote host is available.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemotehostsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemotehostsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RemoteHost**](RemoteHost.md)|  | 

### Return type

[**RemoteHost**](RemoteHost.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceDiscoveryInstanceTypePost

> []Swagger ServiceDiscoveryInstanceTypePost(ctx, instance, type_, optional)

Returns a list of services hosted on the specified Gateway instance

Returns a list of services hosted on the specified Gateway instance.  __type__ is one of: rest, wsdl.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string**|  | 
**type_** | **string**| The type of service to return. Possible values are: &#39;rest&#39;, &#39;wsdl&#39; | 
 **optional** | ***ServiceDiscoveryInstanceTypePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ServiceDiscoveryInstanceTypePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **host** | **optional.String**|  | 
 **port** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 

### Return type

[**[]Swagger**](Swagger.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SysconfigGet

> SystemConfig SysconfigGet(ctx, )

Gets API Manager system configuration

Returns an API Manager system configuration object.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SystemConfig**](SystemConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SysconfigPut

> SystemConfig SysconfigPut(ctx, optional)

Update API Manager system configuration

Returns an API Manager system configuration object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SysconfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SysconfigPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SystemConfig**](SystemConfig.md)|  | 

### Return type

[**SystemConfig**](SystemConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TitleGet

> string TitleGet(ctx, )

Gets the API Manager's title

Returns the API Manager title.  Does not require authentication.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenstoresGet

> []ReferencedEntity TokenstoresGet(ctx, )

Gets a list of Token Stores

Returns a list of Token Stores to be used by OAuth Security Devices for inbound security on portal-registered APIs.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ReferencedEntity**](ReferencedEntity.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TopologyPost

> Topology TopologyPost(ctx, host, port, username, password)

Retrieves the Topology from the specified Admin Node Manager

Retrieves the Topology from the specified Admin Node Manager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**host** | **string**| The host on which the Admin Node Manager is running | 
**port** | **string**| The Admin Node Manager management port. | 
**username** | **string**| Username to use for Admin Node Manager authentication . | 
**password** | **string**| Password to use for Admin Node Manager authentication. | 

### Return type

[**Topology**](Topology.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

