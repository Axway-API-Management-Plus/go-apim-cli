# \QuotasApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotasGet**](QuotasApi.md#QuotasGet) | **Get** /quotas | Returns all quotas
[**QuotasIdDelete**](QuotasApi.md#QuotasIdDelete) | **Delete** /quotas/{id} | Deletes a quota
[**QuotasIdGet**](QuotasApi.md#QuotasIdGet) | **Get** /quotas/{id} | Returns the quota with the given ID
[**QuotasIdPut**](QuotasApi.md#QuotasIdPut) | **Put** /quotas/{id} | Updates a quota
[**QuotasPost**](QuotasApi.md#QuotasPost) | **Post** /quotas | Creates a new quota



## QuotasGet

> QuotaDto QuotasGet(ctx, )

Returns all quotas

This method may be called by any member of the Portal, however only the API Administrator may retrieve the system quota.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**QuotaDto**](QuotaDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotasIdDelete

> QuotasIdDelete(ctx, id)

Deletes a quota

Deletes a quota.  Only API Administrators may update quotas.  Default system and application quotas may not be deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The quota ID to delete. | 

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


## QuotasIdGet

> QuotaDto QuotasIdGet(ctx, id)

Returns the quota with the given ID

Returns a quota.  This method may be called by any member of the Portal, however, only API Administrators may retrieve the system quota.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The quota ID to retrieve. | 

### Return type

[**QuotaDto**](QuotaDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotasIdPut

> QuotaDto QuotasIdPut(ctx, id, optional)

Updates a quota

Updates an existing quota. Only API Administrators may update quotas.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***QuotasIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotasIdPutOpts struct


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


## QuotasPost

> QuotaDto QuotasPost(ctx, optional)

Creates a new quota

Creates a new quota.  Only API Administrators may create quotas.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotasPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QuotasPostOpts struct


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

