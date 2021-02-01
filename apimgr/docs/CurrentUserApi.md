# \CurrentUserApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrentuserChangepasswordPost**](CurrentUserApi.md#CurrentuserChangepasswordPost) | **Post** /currentuser/changepassword | Modify the current user&#39;s password
[**CurrentuserGet**](CurrentUserApi.md#CurrentuserGet) | **Get** /currentuser | Get the current user
[**CurrentuserPut**](CurrentUserApi.md#CurrentuserPut) | **Put** /currentuser | Modify the current user



## CurrentuserChangepasswordPost

> CurrentuserChangepasswordPost(ctx, oldPassword, newPassword)

Modify the current user's password

Modify the password of the authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oldPassword** | **string**| The user&#39;s old password | 
**newPassword** | **string**| The user&#39;s new password | 

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


## CurrentuserGet

> User CurrentuserGet(ctx, )

Get the current user

Get the account details of the authenticated user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrentuserPut

> User CurrentuserPut(ctx, optional)

Modify the current user

Modify the account details of the authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CurrentuserPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CurrentuserPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

