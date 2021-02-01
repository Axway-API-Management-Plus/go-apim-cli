# \OAuthAuthorizationsApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationsAuthzidDelete**](OAuthAuthorizationsApi.md#AuthorizationsAuthzidDelete) | **Delete** /authorizations/{authzid} | Delete the OAuth Authorization for the given authorization id.
[**AuthorizationsGet**](OAuthAuthorizationsApi.md#AuthorizationsGet) | **Get** /authorizations | Retrieve all stored OAuth Authorizations for the logged in user.
[**AuthorizationsOwnerSubjectidApplicationAppidDelete**](OAuthAuthorizationsApi.md#AuthorizationsOwnerSubjectidApplicationAppidDelete) | **Delete** /authorizations/owner/{subjectid}/application/{appid} | 



## AuthorizationsAuthzidDelete

> AuthorizationsAuthzidDelete(ctx, authzid)

Delete the OAuth Authorization for the given authorization id.

Admin or Resource Owner task to delete the given authorization id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authzid** | **string**|  | 

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


## AuthorizationsGet

> []Authorization AuthorizationsGet(ctx, )

Retrieve all stored OAuth Authorizations for the logged in user.

If user is a member of the admin group then all authorizations are returned. If not, then the logged in user's authorizations are returned.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Authorization**](Authorization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationsOwnerSubjectidApplicationAppidDelete

> AuthorizationsOwnerSubjectidApplicationAppidDelete(ctx, subjectid, appid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectid** | **string**|  | 
**appid** | **string**|  | 

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

